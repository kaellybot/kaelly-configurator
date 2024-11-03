//nolint:dupl,nolintlint // OK for DRY concept but refactor at any cost is not relevant here.
package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *Impl) twitchRequest(message *amqp.RabbitMQMessage, correlationID string) {
	request := message.ConfigurationSetTwitchWebhookRequest
	if !isValidTwitchRequest(request) {
		service.publishFailedSetAnswer(correlationID, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationID, correlationID).
		Str(constants.LogGuildID, request.GuildId).
		Str(constants.LogChannelID, request.ChannelId).
		Msgf("Set twitch webhook configuration request received")

	oldWebhook, errGet := service.channelService.GetTwitchWebhook(request.GuildId, request.ChannelId, request.StreamerId)
	if errGet != nil {
		log.Error().Err(errGet).Str(constants.LogCorrelationID, correlationID).
			Str(constants.LogGuildID, request.GuildId).
			Str(constants.LogChannelID, request.ChannelId).
			Str(constants.LogStreamerID, request.StreamerId).
			Msgf("Twitch webhook retrieval has failed, answering with failed response")
		service.publishFailedSetWebhookAnswer(correlationID, request.WebhookId, message.Language)
		return
	}

	if request.Enabled {
		errSave := service.channelService.SaveTwitchWebhook(entities.WebhookTwitch{
			WebhookID:    request.WebhookId,
			WebhookToken: request.WebhookToken,
			GuildID:      request.GuildId,
			ChannelID:    request.ChannelId,
			StreamerID:   request.StreamerId,
			Locale:       message.Language,
			RetryNumber:  0,
		})
		if errSave != nil {
			log.Error().Err(errSave).Str(constants.LogCorrelationID, correlationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogStreamerID, request.StreamerId).
				Msgf("Twitch webhook save has failed, answering with failed response")
			service.publishFailedSetWebhookAnswer(correlationID, request.WebhookId, message.Language)
			return
		}
	} else {
		errDel := service.channelService.DeleteTwitchWebhook(oldWebhook)
		if errDel != nil {
			log.Error().Err(errDel).Str(constants.LogCorrelationID, correlationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogStreamerID, request.StreamerId).
				Msgf("Twitch webhook removal has failed, answering with failed response")
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

func isValidTwitchRequest(request *amqp.ConfigurationSetTwitchWebhookRequest) bool {
	return request != nil
}
