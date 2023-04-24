package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/rs/zerolog/log"
)

func (service *ConfiguratorServiceImpl) rssRequest(message *amqp.RabbitMQMessage, correlationId string) {
	request := message.ConfigurationSetRssWebhookRequest
	if !isValidRssRequest(request) {
		service.publishFailedSetAnswer(correlationId, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationId, correlationId).
		Str(constants.LogGuildId, request.GuildId).
		Str(constants.LogChannelId, request.ChannelId).
		Msgf("Set rss webhook configuration request received")

	// TODO

	service.publishSucceededSetAnswer(correlationId, message.Language)
}

func isValidRssRequest(request *amqp.ConfigurationSetRssWebhookRequest) bool {
	return request != nil
}
