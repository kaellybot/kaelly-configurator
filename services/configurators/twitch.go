//nolint:dupl,nolintlint // OK for DRY concept but refactor at any cost is not relevant here.
package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *Impl) twitchRequest(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	request := message.ConfigurationSetTwitchWebhookRequest
	if !isValidTwitchRequest(request) {
		service.publishFailedSetAnswer(ctx, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationID, ctx.CorrelationID).
		Str(constants.LogGuildID, request.GuildId).
		Str(constants.LogChannelID, request.ChannelId).
		Str(constants.LogStreamerID, request.StreamerId).
		Msgf("Set twitch webhook configuration request received")

	oldWebhook, errGet := service.channelService.GetTwitchWebhook(request.GuildId, request.ChannelId, request.StreamerId)
	if errGet != nil {
		log.Error().Err(errGet).Str(constants.LogCorrelationID, ctx.CorrelationID).
			Str(constants.LogGuildID, request.GuildId).
			Str(constants.LogChannelID, request.ChannelId).
			Str(constants.LogStreamerID, request.StreamerId).
			Msgf("Twitch webhook retrieval has failed, answering with failed response")
		service.publishFailedSetWebhookAnswer(ctx, request.WebhookId, message.Language)
		return
	}

	if request.Enabled {
		errSave := service.channelService.SaveTwitchWebhook(entities.WebhookTwitch{
			WebhookID:    request.WebhookId,
			WebhookToken: request.WebhookToken,
			GuildID:      request.GuildId,
			ChannelID:    request.ChannelId,
			StreamerID:   request.StreamerId,
			Game:         message.Game,
			Locale:       message.Language,
		})
		if errSave != nil {
			log.Error().Err(errSave).Str(constants.LogCorrelationID, ctx.CorrelationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogStreamerID, request.StreamerId).
				Msgf("Twitch webhook save has failed, answering with failed response")
			service.publishFailedSetWebhookAnswer(ctx, request.WebhookId, message.Language)
			return
		}
	} else {
		errDel := service.channelService.DeleteTwitchWebhook(oldWebhook)
		if errDel != nil {
			log.Error().Err(errDel).Str(constants.LogCorrelationID, ctx.CorrelationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogStreamerID, request.StreamerId).
				Msgf("Twitch webhook removal has failed, answering with failed response")
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

func isValidTwitchRequest(request *amqp.ConfigurationSetTwitchWebhookRequest) bool {
	return request != nil
}
