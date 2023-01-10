package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *ConfiguratorServiceImpl) serverRequest(correlationId string,
	request *amqp.ConfigurationSetRequest, lg amqp.RabbitMQMessage_Language) {

	if !isValidConfigurationServerRequest(request) {
		service.publishFailedSetAnswer(correlationId, lg)
		return
	}

	log.Info().Str(constants.LogCorrelationId, correlationId).
		Str(constants.LogGuildId, request.GuildId).
		Str(constants.LogChannelId, request.ChannelId).
		Str(constants.LogServerId, request.ServerField.ServerId).
		Msgf("Set server configuration request received")

	var err error
	if len(request.ChannelId) == 0 {
		err = service.updateGuildServer(request.GuildId, request.ServerField.ServerId)
	} else {
		err = service.updateChannelServer(request.GuildId, request.ChannelId, request.ServerField.ServerId)
	}

	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationId, correlationId).
			Str(constants.LogGuildId, request.GuildId).
			Str(constants.LogChannelId, request.ChannelId).
			Str(constants.LogServerId, request.ServerField.ServerId).
			Msgf("Returning failed message")
		service.publishFailedSetAnswer(correlationId, lg)
		return
	}

	service.publishSucceededSetAnswer(correlationId, lg)
}

func (service *ConfiguratorServiceImpl) updateGuildServer(guildId, serverId string) error {
	return service.guildService.Save(entities.Guild{
		Id:       guildId,
		ServerId: &serverId,
	})
}

func (service *ConfiguratorServiceImpl) updateChannelServer(guildId, channelId, serverId string) error {
	return service.channelService.SaveChannelServer(entities.ChannelServer{
		GuildId:   guildId,
		ChannelId: channelId,
		ServerId:  serverId,
	})
}

func isValidConfigurationServerRequest(request *amqp.ConfigurationSetRequest) bool {
	return request.GetServerField() != nil
}
