package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/mappers"
	"github.com/kaellybot/kaelly-configurator/utils/replies"
	"github.com/rs/zerolog/log"
)

func (service *Impl) getRequest(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	request := message.ConfigurationGetRequest
	if !isValidConfigurationGetRequest(request) {
		replies.FailedAnswer(ctx, service.broker, amqp.RabbitMQMessage_CONFIGURATION_GET_ANSWER,
			message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationID, ctx.CorrelationID).
		Str(constants.LogGuildID, request.GuildId).
		Str(constants.LogGame, message.Game.String()).
		Msgf("Get configuration request received")

	guild, err := service.guildService.Get(request.GuildId, message.Game)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Str(constants.LogGuildID, request.GuildId).
			Str(constants.LogGame, message.Game.String()).
			Msgf("Returning failed answer")
		replies.FailedAnswer(ctx, service.broker, amqp.RabbitMQMessage_CONFIGURATION_GET_ANSWER,
			message.Language)
		return
	}

	response := mappers.MapGuild(guild, message.Language)
	replies.SucceededAnswer(ctx, service.broker, response)
}

func isValidConfigurationGetRequest(request *amqp.ConfigurationGetRequest) bool {
	return request != nil
}
