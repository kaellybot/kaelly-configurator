package mappers

import (
	"testing"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
)

func TestMapGuildCarriesRequestGame(t *testing.T) {
	t.Parallel()

	for _, game := range []amqp.Game{amqp.Game_DOFUS_GAME, amqp.Game_DOFUS_TOUCH} {
		t.Run(game.String(), func(t *testing.T) {
			t.Parallel()

			request := &amqp.RabbitMQMessage{
				Type:     amqp.RabbitMQMessage_CONFIGURATION_GET_REQUEST,
				Game:     game,
				Language: amqp.Language_DE,
			}

			answer := MapGuild(request, entities.Guild{ID: "guildID", Game: game})

			if answer.GetGame() != game {
				t.Errorf("Game = %v, want %v", answer.GetGame(), game)
			}
			if !answer.IsGameSet() {
				t.Error("answer game must be set, otherwise the broker refuses to publish it")
			}
			if answer.GetLanguage() != amqp.Language_DE {
				t.Errorf("Language = %v, want %v", answer.GetLanguage(), amqp.Language_DE)
			}
			if answer.GetStatus() != amqp.RabbitMQMessage_SUCCESS {
				t.Errorf("Status = %v, want %v", answer.GetStatus(), amqp.RabbitMQMessage_SUCCESS)
			}
			if answer.GetConfigurationGetAnswer().GetGuildId() != "guildID" {
				t.Errorf("GuildId = %q, want %q",
					answer.GetConfigurationGetAnswer().GetGuildId(), "guildID")
			}
		})
	}
}
