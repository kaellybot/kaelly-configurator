package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *Impl) rssRequest(message *amqp.RabbitMQMessage, correlationId string) {
	request := message.ConfigurationSetRssWebhookRequest
	if !isValidFeedRequest(request) {
		service.publishFailedSetAnswer(correlationId, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationId, correlationId).
		Str(constants.LogGuildId, request.GuildId).
		Str(constants.LogChannelId, request.ChannelId).
		Msgf("Set feed webhook configuration request received")

	oldWebhook, err := service.channelService.GetFeedWebhook(request.GuildId, request.ChannelId, request.FeedId, request.Language)
	if err != nil {
		log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).
			Str(constants.LogGuildId, request.GuildId).
			Str(constants.LogChannelId, request.ChannelId).
			Str(constants.LogFeedTypeId, request.FeedId).
			Msgf("Feed webhook retrieval has failed, answering with failed response")
		service.publishFailedSetWebhookAnswer(correlationId, request.WebhookId, message.Language)
		return
	}

	if request.Enabled {
		err := service.channelService.SaveFeedWebhook(entities.WebhookFeed{
			WebhookId:    request.WebhookId,
			WebhookToken: request.WebhookToken,
			GuildId:      request.GuildId,
			ChannelId:    request.ChannelId,
			Locale:       request.Language,
			FeedTypeId:   request.FeedId,
			RetryNumber:  0,
		})
		if err != nil {
			log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).
				Str(constants.LogGuildId, request.GuildId).
				Str(constants.LogChannelId, request.ChannelId).
				Str(constants.LogFeedTypeId, request.FeedId).
				Msgf("Feed webhook save has failed, answering with failed response")
			service.publishFailedSetWebhookAnswer(correlationId, request.WebhookId, message.Language)
			return
		}
	} else {
		err := service.channelService.DeleteFeedWebhook(oldWebhook)
		if err != nil {
			log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).
				Str(constants.LogGuildId, request.GuildId).
				Str(constants.LogChannelId, request.ChannelId).
				Str(constants.LogFeedTypeId, request.FeedId).
				Msgf("Feed webhook removal has failed, answering with failed response")
			service.publishFailedSetAnswer(correlationId, message.Language)
			return
		}
	}

	if oldWebhook != nil {
		service.publishSucceededSetWebhookAnswer(correlationId, oldWebhook.WebhookId, message.Language)
	} else {
		service.publishSucceededSetAnswer(correlationId, message.Language)
	}
}

func isValidFeedRequest(request *amqp.ConfigurationSetRssWebhookRequest) bool {
	return request != nil
}
