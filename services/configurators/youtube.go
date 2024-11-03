//nolint:dupl,nolintlint // OK for DRY concept but refactor at any cost is not relevant here.
package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *Impl) youtubeRequest(message *amqp.RabbitMQMessage, correlationID string) {
	request := message.ConfigurationSetYoutubeWebhookRequest
	if !isValidYoutubeRequest(request) {
		service.publishFailedSetAnswer(correlationID, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationID, correlationID).
		Str(constants.LogGuildID, request.GuildId).
		Str(constants.LogChannelID, request.ChannelId).
		Msgf("Set youtube webhook configuration request received")

	oldWebhook, errGet := service.channelService.GetYoutubeWebhook(request.GuildId, request.ChannelId, request.VideastId)
	if errGet != nil {
		log.Error().Err(errGet).Str(constants.LogCorrelationID, correlationID).
			Str(constants.LogGuildID, request.GuildId).
			Str(constants.LogChannelID, request.ChannelId).
			Str(constants.LogVideastID, request.VideastId).
			Msgf("Youtube webhook retrieval has failed, answering with failed response")
		service.publishFailedSetWebhookAnswer(correlationID, request.WebhookId, message.Language)
		return
	}

	if request.Enabled {
		errSave := service.channelService.SaveYoutubeWebhook(entities.WebhookYoutube{
			WebhookID:    request.WebhookId,
			WebhookToken: request.WebhookToken,
			GuildID:      request.GuildId,
			ChannelID:    request.ChannelId,
			VideastID:    request.VideastId,
			Locale:       message.Language,
			RetryNumber:  0,
		})
		if errSave != nil {
			log.Error().Err(errSave).Str(constants.LogCorrelationID, correlationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogVideastID, request.VideastId).
				Msgf("Youtube webhook save has failed, answering with failed response")
			service.publishFailedSetWebhookAnswer(correlationID, request.WebhookId, message.Language)
			return
		}
	} else {
		errDel := service.channelService.DeleteYoutubeWebhook(oldWebhook)
		if errDel != nil {
			log.Error().Err(errDel).Str(constants.LogCorrelationID, correlationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogVideastID, request.VideastId).
				Msgf("Youtube webhook removal has failed, answering with failed response")
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

func isValidYoutubeRequest(request *amqp.ConfigurationSetYoutubeWebhookRequest) bool {
	return request != nil
}
