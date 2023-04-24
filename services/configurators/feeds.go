package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
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

	// TODO else delete, checks if it exists, etc.
	if request.Enabled {
		err := service.channelService.SaveFeedWebhook(entities.WebhookFeed{
			WebhookId:    request.WebhookId,
			WebhookToken: request.WebhookToken,
			GuildId:      request.GuildId,
			ChannelId:    request.ChannelId,
			Locale:       request.Language,
			FeedTypeId:   request.FeedId,
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

func isValidRssRequest(request *amqp.ConfigurationSetRssWebhookRequest) bool {
	return request != nil
}
