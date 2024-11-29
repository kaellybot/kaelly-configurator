//nolint:dupl,nolintlint // OK for DRY concept but refactor at any cost is not relevant here.
package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/rs/zerolog/log"
)

func (service *Impl) twitterRequest(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	request := message.ConfigurationSetTwitterWebhookRequest
	if !isValidTwitterRequest(request) {
		service.publishFailedSetAnswer(ctx, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationID, ctx.CorrelationID).
		Str(constants.LogGuildID, request.GuildId).
		Str(constants.LogChannelID, request.ChannelId).
		Str(constants.LogTwitterID, request.TwitterId).
		Msgf("Set twitter webhook configuration request received")

	oldWebhook, errGet := service.channelService.GetTwitterWebhook(request.GuildId, request.ChannelId, request.TwitterId)
	if errGet != nil {
		log.Error().Err(errGet).Str(constants.LogCorrelationID, ctx.CorrelationID).
			Str(constants.LogGuildID, request.GuildId).
			Str(constants.LogChannelID, request.ChannelId).
			Str(constants.LogTwitterID, request.TwitterId).
			Msgf("Twitter webhook retrieval has failed, answering with failed response")
		service.publishFailedSetWebhookAnswer(ctx, request.WebhookId, message.Language)
		return
	}

	if request.Enabled {
		errSave := service.channelService.SaveTwitterWebhook(entities.WebhookTwitter{
			WebhookID:    request.WebhookId,
			WebhookToken: request.WebhookToken,
			GuildID:      request.GuildId,
			ChannelID:    request.ChannelId,
			TwitterID:    request.TwitterId,
			Game:         message.Game,
			Locale:       message.Language,
		})
		if errSave != nil {
			log.Error().Err(errSave).Str(constants.LogCorrelationID, ctx.CorrelationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogTwitterID, request.TwitterId).
				Msgf("Twitter webhook save has failed, answering with failed response")
			service.publishFailedSetWebhookAnswer(ctx, request.WebhookId, message.Language)
			return
		}
	} else {
		errDel := service.channelService.DeleteTwitterWebhook(oldWebhook)
		if errDel != nil {
			log.Error().Err(errDel).Str(constants.LogCorrelationID, ctx.CorrelationID).
				Str(constants.LogGuildID, request.GuildId).
				Str(constants.LogChannelID, request.ChannelId).
				Str(constants.LogTwitterID, request.TwitterId).
				Msgf("Twitter webhook removal has failed, answering with failed response")
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

func isValidTwitterRequest(request *amqp.ConfigurationSetTwitterWebhookRequest) bool {
	return request != nil
}
