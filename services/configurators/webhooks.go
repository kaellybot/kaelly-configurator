package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/rs/zerolog/log"
)

func (service *ConfiguratorServiceImpl) webhookRequest(correlationId string,
	request *amqp.ConfigurationSetRequest, lg amqp.RabbitMQMessage_Language) {

	if !isValidConfigurationWebhookRequest(request) {
		service.publishFailedSetAnswer(correlationId, lg)
		return
	}

	log.Info().Str(constants.LogCorrelationId, correlationId).
		Str(constants.LogGuildId, request.GuildId).
		Str(constants.LogChannelId, request.ChannelId).
		Msgf("Set webhook configuration request received")

	// TODO

	service.publishSucceededSetAnswer(correlationId, lg)
}

func isValidConfigurationWebhookRequest(request *amqp.ConfigurationSetRequest) bool {
	return request.GetWebhookField() != nil
}
