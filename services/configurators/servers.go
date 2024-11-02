package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *Impl) serverRequest(message *amqp.RabbitMQMessage, correlationID string) {
	request := message.ConfigurationSetServerRequest
	if !isValidConfigurationServerRequest(request) {
		service.publishFailedSetAnswer(correlationID, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationID, correlationID).
		Str(constants.LogGuildID, request.GuildId).
		Str(constants.LogChannelID, request.ChannelId).
		Str(constants.LogServerID, request.ServerId).
		Str(constants.LogGame, message.Game.String()).
		Msgf("Set server configuration request received")

	var err error
	if len(request.ChannelId) == 0 {
		err = service.updateGuildServer(request.GuildId, request.ServerId, message.Game)
	} else {
		err = service.updateChannelServer(request.GuildId, request.ChannelId, request.ServerId, message.Game)
	}

	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationID, correlationID).
			Str(constants.LogGuildID, request.GuildId).
			Str(constants.LogChannelID, request.ChannelId).
			Str(constants.LogServerID, request.ServerId).
			Str(constants.LogGame, message.Game.String()).
			Msgf("Returning failed message")
		service.publishFailedSetAnswer(correlationID, message.Language)
		return
	}

	service.publishSucceededSetAnswer(correlationID, message.Language)
}

func (service *Impl) updateGuildServer(guildID, serverID string, game amqp.Game) error {
	return service.guildService.Save(entities.Guild{
		ID:       guildID,
		ServerID: &serverID,
		Game:     game,
	})
}

func (service *Impl) updateChannelServer(guildID, channelID, serverID string, game amqp.Game) error {
	return service.channelService.SaveChannelServer(entities.ChannelServer{
		GuildID:   guildID,
		ChannelID: channelID,
		ServerID:  serverID,
		Game:      game,
	})
}

func isValidConfigurationServerRequest(request *amqp.ConfigurationSetServerRequest) bool {
	return request != nil
}
