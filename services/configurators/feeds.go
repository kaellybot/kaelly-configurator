package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *Impl) rssRequest(message *amqp.RabbitMQMessage, correlationID string) {
	request := message.ConfigurationSetRssWebhookRequest
	if !isValidFeedRequest(request) {
		service.publishFailedSetAnswer(correlationID, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationID, correlationID).
		Str(constants.LogGuildID, request.GuildId).
		Str(constants.LogChannelID, request.ChannelId).
		Str(constants.LogGame, message.Game.String()).
		Msgf("Set feed webhook configuration request received")

	oldWebhook, errGet := service.channelService.GetFeedWebhook(request.GuildId,
		request.ChannelId, request.FeedId, request.Language, message.GetGame())
	if errGet != nil {
		log.Error().Err(errGet).Str(constants.LogCorrelationID, correlationID).
			Str(constants.LogGuildID, request.GuildId).
			Str(constants.LogChannelID, request.ChannelId).
			Str(constants.LogFeedTypeID, request.FeedId).
			Str(constants.LogGame, message.Game.String()).
			Msgf("Feed webhook retrieval has failed, answering with failed response")
		service.publishFailedSetWebhookAnswer(correlationID, request.WebhookId, message.Language)
		return
	}

	if request.Enabled {
		errSave := service.channelService.SaveFeedWebhook(entities.WebhookFeed{
			WebhookID:    request.WebhookId,
			WebhookToken: request.WebhookToken,
			GuildID:      request.GuildId,
			ChannelID:    request.ChannelId,
			Locale:       request.Language,
			FeedTypeID:   request.FeedId,
			Game:         message.Game,
			RetryNumber:  0,
		})
		if errSave != nil {
			log.Error().Err(errSave).Str(constants.LogCorrelationID, correlationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogFeedTypeID, request.FeedId).
				Str(constants.LogGame, message.Game.String()).
				Msgf("Feed webhook save has failed, answering with failed response")
			service.publishFailedSetWebhookAnswer(correlationID, request.WebhookId, message.Language)
			return
		}
	} else {
		errDel := service.channelService.DeleteFeedWebhook(oldWebhook)
		if errDel != nil {
			log.Error().Err(errDel).Str(constants.LogCorrelationID, correlationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogFeedTypeID, request.FeedId).
				Str(constants.LogGame, message.Game.String()).
				Msgf("Feed webhook removal has failed, answering with failed response")
			service.publishFailedSetAnswer(correlationID, message.Language)
			return
		}
	}

	if oldWebhook != nil {
		service.publishSucceededSetWebhookAnswer(correlationID, oldWebhook.WebhookID, message.Language)
	} else {
		service.publishSucceededSetAnswer(correlationID, message.Language)
	}
}

func isValidFeedRequest(request *amqp.ConfigurationSetRssWebhookRequest) bool {
	return request != nil
}
