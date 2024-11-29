package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *Impl) rssRequest(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	request := message.ConfigurationSetRssWebhookRequest
	if !isValidFeedRequest(request) {
		service.publishFailedSetAnswer(ctx, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationID, ctx.CorrelationID).
		Str(constants.LogGuildID, request.GuildId).
		Str(constants.LogChannelID, request.ChannelId).
		Str(constants.LogGame, message.Game.String()).
		Msgf("Set feed webhook configuration request received")

	oldWebhook, errGet := service.channelService.GetFeedWebhook(request.GuildId,
		request.ChannelId, request.FeedId, message.GetGame())
	if errGet != nil {
		log.Error().Err(errGet).Str(constants.LogCorrelationID, ctx.CorrelationID).
			Str(constants.LogGuildID, request.GuildId).
			Str(constants.LogChannelID, request.ChannelId).
			Str(constants.LogFeedTypeID, request.FeedId).
			Str(constants.LogGame, message.Game.String()).
			Msgf("Feed webhook retrieval has failed, answering with failed response")
		service.publishFailedSetWebhookAnswer(ctx, request.WebhookId, message.Language)
		return
	}

	if request.Enabled {
		errSave := service.channelService.SaveFeedWebhook(entities.WebhookFeed{
			WebhookID:    request.WebhookId,
			WebhookToken: request.WebhookToken,
			GuildID:      request.GuildId,
			ChannelID:    request.ChannelId,
			FeedTypeID:   request.FeedId,
			Game:         message.Game,
			Locale:       message.Language,
		})
		if errSave != nil {
			log.Error().Err(errSave).Str(constants.LogCorrelationID, ctx.CorrelationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogFeedTypeID, request.FeedId).
				Str(constants.LogGame, message.Game.String()).
				Msgf("Feed webhook save has failed, answering with failed response")
			service.publishFailedSetWebhookAnswer(ctx, request.WebhookId, message.Language)
			return
		}
	} else {
		errDel := service.channelService.DeleteFeedWebhook(oldWebhook)
		if errDel != nil {
			log.Error().Err(errDel).Str(constants.LogCorrelationID, ctx.CorrelationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogFeedTypeID, request.FeedId).
				Str(constants.LogGame, message.Game.String()).
				Msgf("Feed webhook removal has failed, answering with failed response")
			service.publishFailedSetAnswer(ctx, message.Language)
			return
		}
	}

	if oldWebhook != nil {
		service.publishSucceededSetWebhookAnswer(ctx, oldWebhook.WebhookID, message.Language)
	} else {
		service.publishSucceededSetAnswer(ctx, message.Language)
	}
}

func isValidFeedRequest(request *amqp.ConfigurationSetRssWebhookRequest) bool {
	return request != nil
}
