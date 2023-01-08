package configurators

import (
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/rs/zerolog/log"
)

func (service *ConfiguratorServiceImpl) webhookRequest(correlationId, guildId, channelId string) {
	log.Info().Str(constants.LogCorrelationId, correlationId).
		Str(constants.LogGuildId, guildId).
		Str(constants.LogChannelId, channelId).
		Msgf("Webhook configuration request received")

	// TODO
}
