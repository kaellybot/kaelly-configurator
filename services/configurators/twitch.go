package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *Impl) twitchRequest(message *amqp.RabbitMQMessage, correlationId string) {
	request := message.ConfigurationSetTwitchWebhookRequest
	if !isValidTwitchRequest(request) {
		service.publishFailedSetAnswer(correlationId, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationId, correlationId).
		Str(constants.LogGuildId, request.GuildId).
		Str(constants.LogChannelId, request.ChannelId).
		Msgf("Set twitch webhook configuration request received")

	oldWebhook, err := service.channelService.GetTwitchWebhook(request.GuildId, request.ChannelId, request.StreamerId)
	if err != nil {
		log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).
			Str(constants.LogGuildId, request.GuildId).
			Str(constants.LogChannelId, request.ChannelId).
			Str(constants.LogStreamerId, request.StreamerId).
			Msgf("Twitch webhook retrieval has failed, answering with failed response")
		service.publishFailedSetWebhookAnswer(correlationId, request.WebhookId, message.Language)
		return
	}

	if request.Enabled {
		err := service.channelService.SaveTwitchWebhook(entities.WebhookTwitch{
			WebhookId:    request.WebhookId,
			WebhookToken: request.WebhookToken,
			GuildId:      request.GuildId,
			ChannelId:    request.ChannelId,
			StreamerId:   request.StreamerId,
			RetryNumber:  0,
		})
		if err != nil {
			log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).
				Str(constants.LogGuildId, request.GuildId).
				Str(constants.LogChannelId, request.ChannelId).
				Str(constants.LogStreamerId, request.StreamerId).
				Msgf("Twitch webhook save has failed, answering with failed response")
			service.publishFailedSetWebhookAnswer(correlationId, request.WebhookId, message.Language)
			return
		}
	} else {
		err := service.channelService.DeleteTwitchWebhook(oldWebhook)
		if err != nil {
			log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).
				Str(constants.LogGuildId, request.GuildId).
				Str(constants.LogChannelId, request.ChannelId).
				Str(constants.LogStreamerId, request.StreamerId).
				Msgf("Twitch webhook removal has failed, answering with failed response")
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

func isValidTwitchRequest(request *amqp.ConfigurationSetTwitchWebhookRequest) bool {
	return request != nil
}
