//nolint:dupl,nolintlint // OK for DRY concept but refactor at any cost is not relevant here.
package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *Impl) almanaxRequest(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	request := message.ConfigurationSetNotificationRequest
	log.Info().Str(constants.LogCorrelationID, ctx.CorrelationID).
		Str(constants.LogGuildID, request.GuildId).
		Str(constants.LogChannelID, request.ChannelId).
		Str(constants.LogGame, message.Game.String()).
		Msgf("Set almanax webhook configuration request received")

	oldWebhook, errGet := service.channelService.GetAlmanaxWebhook(request.GuildId,
		request.ChannelId, message.GetGame())
	if errGet != nil {
		log.Error().Err(errGet).Str(constants.LogCorrelationID, ctx.CorrelationID).
			Str(constants.LogGuildID, request.GuildId).
			Str(constants.LogChannelID, request.ChannelId).
			Str(constants.LogGame, message.Game.String()).
			Msgf("Almanax webhook retrieval has failed, answering with failed response")
		service.publishFailedSetNotificationAnswer(ctx, request.WebhookId, message.Language)
		return
	}

	if request.Enabled {
		errSave := service.channelService.SaveAlmanaxWebhook(entities.WebhookAlmanax{
			WebhookID: request.WebhookId,
			GuildID:   request.GuildId,
			ChannelID: request.ChannelId,
			Game:      message.Game,
			Locale:    message.Language,
		})
		if errSave != nil {
			log.Error().Err(errSave).Str(constants.LogCorrelationID, ctx.CorrelationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogGame, message.Game.String()).
				Msgf("Almanax webhook save has failed, answering with failed response")
			service.publishFailedSetNotificationAnswer(ctx, request.WebhookId, message.Language)
			return
		}
	} else {
		errDel := service.channelService.DeleteAlmanaxWebhook(oldWebhook)
		if errDel != nil {
			log.Error().Err(errDel).Str(constants.LogCorrelationID, ctx.CorrelationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogGame, message.Game.String()).
				Msgf("Almanax webhook removal has failed, answering with failed response")
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
