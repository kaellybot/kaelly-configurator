package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *Impl) youtubeRequest(message *amqp.RabbitMQMessage, correlationId string) {
	request := message.ConfigurationSetYoutubeWebhookRequest
	if !isValidYoutubeRequest(request) {
		service.publishFailedSetAnswer(correlationId, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationId, correlationId).
		Str(constants.LogGuildId, request.GuildId).
		Str(constants.LogChannelId, request.ChannelId).
		Msgf("Set youtube webhook configuration request received")

	oldWebhook, err := service.channelService.GetYoutubeWebhook(request.GuildId, request.ChannelId, request.VideastId)
	if err != nil {
		log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).
			Str(constants.LogGuildId, request.GuildId).
			Str(constants.LogChannelId, request.ChannelId).
			Str(constants.LogVideastId, request.VideastId).
			Msgf("Youtube webhook retrieval has failed, answering with failed response")
		service.publishFailedSetWebhookAnswer(correlationId, request.WebhookId, message.Language)
		return
	}

	if request.Enabled {
		err := service.channelService.SaveYoutubeWebhook(entities.WebhookYoutube{
			WebhookId:    request.WebhookId,
			WebhookToken: request.WebhookToken,
			GuildId:      request.GuildId,
			ChannelId:    request.ChannelId,
			VideastId:    request.VideastId,
			RetryNumber:  0,
		})
		if err != nil {
			log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).
				Str(constants.LogGuildId, request.GuildId).
				Str(constants.LogChannelId, request.ChannelId).
				Str(constants.LogVideastId, request.VideastId).
				Msgf("Youtube webhook save has failed, answering with failed response")
			service.publishFailedSetWebhookAnswer(correlationId, request.WebhookId, message.Language)
			return
		}
	} else {
		err := service.channelService.DeleteYoutubeWebhook(oldWebhook)
		if err != nil {
			log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).
				Str(constants.LogGuildId, request.GuildId).
				Str(constants.LogChannelId, request.ChannelId).
				Str(constants.LogVideastId, request.VideastId).
				Msgf("Youtube webhook removal has failed, answering with failed response")
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

func isValidYoutubeRequest(request *amqp.ConfigurationSetYoutubeWebhookRequest) bool {
	return request != nil
}
