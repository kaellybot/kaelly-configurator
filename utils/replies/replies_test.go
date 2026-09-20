package replies

import (
	"testing"

	amqp "github.com/kaellybot/kaelly-amqp"
)

func TestFailedAnswerCarriesRequestGame(t *testing.T) {
	t.Parallel()

	for _, game := range []amqp.Game{amqp.Game_DOFUS_GAME, amqp.Game_DOFUS_TOUCH} {
		t.Run(game.String(), func(t *testing.T) {
			t.Parallel()

			var published *amqp.RabbitMQMessage
			broker := amqp.Mock{
				ReplyFunc: func(msg *amqp.RabbitMQMessage, _, _ string) error {
					published = msg
					return nil
				},
			}

			request := &amqp.RabbitMQMessage{
				Type:     amqp.RabbitMQMessage_CONFIGURATION_GET_REQUEST,
				Game:     game,
				Language: amqp.Language_PT,
			}

			FailedAnswer(amqp.Context{CorrelationID: "correlationID"}, &broker, request,
				amqp.RabbitMQMessage_CONFIGURATION_GET_ANSWER)

			if published == nil {
				t.Fatal("no reply published")
			}
			if published.GetGame() != game {
				t.Errorf("Game = %v, want %v", published.GetGame(), game)
			}
			if published.GetLanguage() != amqp.Language_PT {
				t.Errorf("Language = %v, want %v", published.GetLanguage(), amqp.Language_PT)
			}
			if published.GetStatus() != amqp.RabbitMQMessage_FAILED {
				t.Errorf("Status = %v, want %v", published.GetStatus(), amqp.RabbitMQMessage_FAILED)
			}
		})
	}
}
