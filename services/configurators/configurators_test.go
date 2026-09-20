package configurators

import (
	"testing"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/services/channels"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
)

func newTestService(replies *[]*amqp.RabbitMQMessage) *Impl {
	broker := amqp.Mock{
		ReplyFunc: func(msg *amqp.RabbitMQMessage, _, _ string) error {
			*replies = append(*replies, msg)
			return nil
		},
	}
	return &Impl{broker: &broker}
}

type requestCase struct {
	requestType amqp.RabbitMQMessage_Type
	answerType  amqp.RabbitMQMessage_Type
}

func requestTypes() []requestCase {
	return []requestCase{
		{
			requestType: amqp.RabbitMQMessage_CONFIGURATION_GET_REQUEST,
			answerType:  amqp.RabbitMQMessage_CONFIGURATION_GET_ANSWER,
		},
		{
			requestType: amqp.RabbitMQMessage_CONFIGURATION_SET_SERVER_REQUEST,
			answerType:  amqp.RabbitMQMessage_CONFIGURATION_SET_SERVER_ANSWER,
		},
		{
			requestType: amqp.RabbitMQMessage_CONFIGURATION_SET_NOTIFICATION_REQUEST,
			answerType:  amqp.RabbitMQMessage_CONFIGURATION_SET_NOTIFICATION_ANSWER,
		},
	}
}

func TestHandleRunsHandlerForSupportedGames(t *testing.T) {
	t.Parallel()

	for _, game := range []amqp.Game{amqp.Game_DOFUS_GAME, amqp.Game_DOFUS_TOUCH} {
		t.Run(game.String(), func(t *testing.T) {
			t.Parallel()

			called := 0
			published := make([]*amqp.RabbitMQMessage, 0)
			service := newTestService(&published)

			service.handle(amqp.Context{CorrelationID: "correlationID"},
				&amqp.RabbitMQMessage{Game: game},
				amqp.RabbitMQMessage_CONFIGURATION_GET_ANSWER,
				func(_ amqp.Context, _ *amqp.RabbitMQMessage) { called++ })

			if called != 1 {
				t.Errorf("handler called %d time(s), want 1", called)
			}
			if len(published) != 0 {
				t.Errorf("handle published %d reply(ies), want 0", len(published))
			}
		})
	}
}

func TestConsumeRequestsRefusesUnsupportedGame(t *testing.T) {
	t.Parallel()

	for _, test := range requestTypes() {
		t.Run(test.requestType.String(), func(t *testing.T) {
			t.Parallel()

			published := make([]*amqp.RabbitMQMessage, 0)
			service := newTestService(&published)

			service.consumeRequests(amqp.Context{CorrelationID: "correlationID"},
				&amqp.RabbitMQMessage{Type: test.requestType, Game: amqp.Game_DOFUS_RETRO})

			if len(published) != 1 {
				t.Fatalf("published %d reply(ies), want 1", len(published))
			}
			if published[0].GetStatus() != amqp.RabbitMQMessage_FAILED {
				t.Errorf("Status = %v, want %v", published[0].GetStatus(),
					amqp.RabbitMQMessage_FAILED)
			}
			if published[0].GetGame() != amqp.Game_DOFUS_RETRO {
				t.Errorf("Game = %v, want %v", published[0].GetGame(), amqp.Game_DOFUS_RETRO)
			}
			if published[0].GetType() != test.answerType {
				t.Errorf("Type = %v, want %v", published[0].GetType(), test.answerType)
			}
		})
	}
}

func TestConsumeRequestsDropsRequestWithoutGame(t *testing.T) {
	t.Parallel()

	for _, test := range requestTypes() {
		t.Run(test.requestType.String(), func(t *testing.T) {
			t.Parallel()

			published := make([]*amqp.RabbitMQMessage, 0)
			service := newTestService(&published)

			service.consumeRequests(amqp.Context{CorrelationID: "correlationID"},
				&amqp.RabbitMQMessage{Type: test.requestType})

			if len(published) != 0 {
				t.Errorf("published %d reply(ies), want 0", len(published))
			}
		})
	}
}

// Guild news gets no answer, so an unusable game must be dropped rather than acted
// on: creating or deleting a guild row under the wrong game is unrecoverable.
func TestConsumeNewsDropsUnusableGame(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		game amqp.Game
	}{
		{name: "no game", game: amqp.Game_ANY_GAME},
		{name: "unsupported game", game: amqp.Game_DOFUS_RETRO},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			published := make([]*amqp.RabbitMQMessage, 0)
			service := newTestService(&published)
			// A nil guildService would panic if the news were acted on.
			service.guildService = guilds.Service(nil)
			service.channelService = channels.Service(nil)

			service.consumeNews(amqp.Context{CorrelationID: "correlationID"},
				&amqp.RabbitMQMessage{
					Type: amqp.RabbitMQMessage_NEWS_GUILD,
					Game: test.game,
					NewsGuildMessage: &amqp.NewsGuildMessage{
						Id:    "guildID",
						Event: amqp.NewsGuildMessage_CREATE,
					},
				})

			if len(published) != 0 {
				t.Errorf("published %d message(s), want 0", len(published))
			}
		})
	}
}

func TestIsGameSupported(t *testing.T) {
	t.Parallel()

	tests := map[amqp.Game]bool{
		amqp.Game_ANY_GAME:    false,
		amqp.Game_DOFUS_GAME:  true,
		amqp.Game_DOFUS_TOUCH: true,
		amqp.Game_DOFUS_RETRO: false,
	}

	for game, want := range tests {
		if got := isGameSupported(game); got != want {
			t.Errorf("isGameSupported(%v) = %v, want %v", game, got, want)
		}
	}
}
