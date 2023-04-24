package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/rs/zerolog/log"
)

func (service *ConfiguratorServiceImpl) almanaxRequest(message *amqp.RabbitMQMessage, correlationId string) {
	request := message.ConfigurationSetAlmanaxWebhookRequest
	if !isValidAlmanaxRequest(request) {
		service.publishFailedSetAnswer(correlationId, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationId, correlationId).
		Str(constants.LogGuildId, request.GuildId).
		Str(constants.LogChannelId, request.ChannelId).
		Msgf("Set almanax webhook configuration request received")

	// TODO

	service.publishSucceededSetAnswer(correlationId, message.Language)
}

func isValidAlmanaxRequest(request *amqp.ConfigurationSetAlmanaxWebhookRequest) bool {
	return request != nil
}
