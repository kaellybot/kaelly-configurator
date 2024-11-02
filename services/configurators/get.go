package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) getRequest(message *amqp.RabbitMQMessage, correlationID string) {
	request := message.ConfigurationGetRequest
	if !isValidConfigurationGetRequest(request) {
		service.publishFailedGetAnswer(correlationID, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationID, correlationID).
		Str(constants.LogGuildID, request.GuildId).
		Str(constants.LogGame, message.Game.String()).
		Msgf("Get configuration request received")

	guild, err := service.guildService.Get(request.GuildId, message.Game)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationID, correlationID).
			Str(constants.LogGuildID, request.GuildId).
			Str(constants.LogGame, message.Game.String()).
			Msgf("Returning failed answer")
		service.publishFailedGetAnswer(correlationID, message.Language)
		return
	}

	service.publishSucceededGetAnswer(correlationID, guild, message.Language)
}

func (service *Impl) publishSucceededGetAnswer(correlationID string,
	guild entities.Guild, lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:                   amqp.RabbitMQMessage_CONFIGURATION_GET_ANSWER,
		Status:                 amqp.RabbitMQMessage_SUCCESS,
		Language:               lg,
		ConfigurationGetAnswer: mappers.MapGuild(guild),
	}

	err := service.broker.Publish(&message, amqp.ExchangeAnswer, answersRoutingkey, correlationID)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationID, correlationID).
			Msgf("Cannot publish via broker, request ignored")
	}
}

func (service *Impl) publishFailedGetAnswer(correlationID string, lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_GET_ANSWER,
		Status:   amqp.RabbitMQMessage_FAILED,
		Language: lg,
	}

	err := service.broker.Publish(&message, amqp.ExchangeAnswer, answersRoutingkey, correlationID)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationID, correlationID).
			Msgf("Cannot publish via broker, request ignored")
	}
}

func isValidConfigurationGetRequest(request *amqp.ConfigurationGetRequest) bool {
	return request != nil
}
