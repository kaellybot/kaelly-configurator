package configurators

import (
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *ConfiguratorServiceImpl) serverRequest(correlationId, guildId, channelId, serverId string) {

	var err error
	if len(channelId) == 0 {
		err = service.updateGuildServer(guildId, serverId)
	} else {
		err = service.updateChannelServer(guildId, channelId, serverId)
	}

	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationId, correlationId).
			Str(constants.LogGuildId, guildId).
			Str(constants.LogChannelId, channelId).
			Str(constants.LogServerId, serverId).
			Msgf("Returning failed message")
		service.publishFailedAnswer(correlationId)
		return
	}

	service.publishSucceededAnswer(correlationId)
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
