package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
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

	// TODO else delete, checks if it exists, etc.
	if request.Enabled {
		err := service.channelService.SaveTwitterWebhook(entities.WebhookTwitter{
			WebhookId:    request.WebhookId,
			WebhookToken: request.WebhookToken,
			GuildId:      request.GuildId,
			ChannelId:    request.ChannelId,
			Locale:       request.Language,
			
		})
		if err != nil {
			log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).
				Str(constants.LogGuildId, request.GuildId).
				Str(constants.LogChannelId, request.ChannelId).
				Msgf("Set almanax webhook configuration request received")
			service.publishFailedSetAnswer(correlationId, message.Language)
			return
		}
	}

	service.publishSucceededSetAnswer(correlationId, message.Language)
}

func isValidTwitterRequest(request *amqp.ConfigurationSetTwitterWebhookRequest) bool {
	return request != nil
}
