package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *Impl) rssRequest(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	request := message.ConfigurationSetNotificationRequest
	log.Info().Str(constants.LogCorrelationID, ctx.CorrelationID).
		Str(constants.LogGuildID, request.GuildId).
		Str(constants.LogChannelID, request.ChannelId).
		Str(constants.LogGame, message.Game.String()).
		Msgf("Set feed webhook configuration request received")

	oldWebhook, errGet := service.channelService.GetFeedWebhook(request.GuildId,
		request.ChannelId, request.Label, message.GetGame())
	if errGet != nil {
		log.Error().Err(errGet).Str(constants.LogCorrelationID, ctx.CorrelationID).
			Str(constants.LogGuildID, request.GuildId).
			Str(constants.LogChannelID, request.ChannelId).
			Str(constants.LogFeedTypeID, request.Label).
			Str(constants.LogGame, message.Game.String()).
			Msgf("Feed webhook retrieval has failed, answering with failed response")
		service.publishFailedSetNotificationAnswer(ctx, request.WebhookId, message.Language)
		return
	}

	if request.Enabled {
		errSave := service.channelService.SaveFeedWebhook(entities.WebhookFeed{
			WebhookID:  request.WebhookId,
			GuildID:    request.GuildId,
			ChannelID:  request.ChannelId,
			FeedTypeID: request.Label,
			Game:       message.Game,
			Locale:     message.Language,
		})
		if errSave != nil {
			log.Error().Err(errSave).Str(constants.LogCorrelationID, ctx.CorrelationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogFeedTypeID, request.Label).
				Str(constants.LogGame, message.Game.String()).
				Msgf("Feed webhook save has failed, answering with failed response")
			service.publishFailedSetNotificationAnswer(ctx, request.WebhookId, message.Language)
			return
		}
	} else {
		errDel := service.channelService.DeleteFeedWebhook(oldWebhook)
		if errDel != nil {
			log.Error().Err(errDel).Str(constants.LogCorrelationID, ctx.CorrelationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogFeedTypeID, request.Label).
				Str(constants.LogGame, message.Game.String()).
				Msgf("Feed webhook removal has failed, answering with failed response")
			service.publishFailedSetNotificationAnswer(ctx, "", message.Language)
			return
		}
	}

	if oldWebhook != nil {
		service.publishSucceededSetNotificationAnswer(ctx, oldWebhook.WebhookID, message.Language)
	} else {
		service.publishSucceededSetNotificationAnswer(ctx, "", message.Language)
	}
}
