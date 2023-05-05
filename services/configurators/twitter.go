package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *Impl) twitterRequest(message *amqp.RabbitMQMessage, correlationId string) {
	request := message.ConfigurationSetTwitterWebhookRequest
	if !isValidTwitterRequest(request) {
		service.publishFailedSetAnswer(correlationId, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationId, correlationId).
		Str(constants.LogGuildId, request.GuildId).
		Str(constants.LogChannelId, request.ChannelId).
		Msgf("Set twitter webhook configuration request received")

	oldWebhook, err := service.channelService.GetTwitterWebhook(request.GuildId, request.ChannelId, request.Language)
	if err != nil {
		log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).
			Str(constants.LogGuildId, request.GuildId).
			Str(constants.LogChannelId, request.ChannelId).
			Msgf("Twitter webhook retrieval has failed, answering with failed response")
		service.publishFailedSetWebhookAnswer(correlationId, request.WebhookId, message.Language)
		return
	}

	if request.Enabled {
		err := service.channelService.SaveTwitterWebhook(entities.WebhookTwitter{
			WebhookId:    request.WebhookId,
			WebhookToken: request.WebhookToken,
			GuildId:      request.GuildId,
			ChannelId:    request.ChannelId,
			Locale:       request.Language,
			RetryNumber:  0,
		})
		if err != nil {
			log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).
				Str(constants.LogGuildId, request.GuildId).
				Str(constants.LogChannelId, request.ChannelId).
				Msgf("Twitter webhook save has failed, answering with failed response")
			service.publishFailedSetWebhookAnswer(correlationId, request.WebhookId, message.Language)
			return
		}
	} else {
		err := service.channelService.DeleteTwitterWebhook(oldWebhook)
		if err != nil {
			log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).
				Str(constants.LogGuildId, request.GuildId).
				Str(constants.LogChannelId, request.ChannelId).
				Msgf("Twitter webhook removal has failed, answering with failed response")
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

func isValidTwitterRequest(request *amqp.ConfigurationSetTwitterWebhookRequest) bool {
	return request != nil
}
