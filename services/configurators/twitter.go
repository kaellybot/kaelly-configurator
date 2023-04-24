package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/rs/zerolog/log"
)

func (service *ConfiguratorServiceImpl) twitterRequest(message *amqp.RabbitMQMessage, correlationId string) {
	request := message.ConfigurationSetTwitterWebhookRequest
	if !isValidTwitterRequest(request) {
		service.publishFailedSetAnswer(correlationId, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationId, correlationId).
		Str(constants.LogGuildId, request.GuildId).
		Str(constants.LogChannelId, request.ChannelId).
		Msgf("Set twitter webhook configuration request received")

	// TODO

	service.publishSucceededSetAnswer(correlationId, message.Language)
}

func isValidTwitterRequest(request *amqp.ConfigurationSetTwitterWebhookRequest) bool {
	return request != nil
}
