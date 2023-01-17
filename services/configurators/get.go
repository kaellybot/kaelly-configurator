package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *ConfiguratorServiceImpl) getRequest(message *amqp.RabbitMQMessage, correlationId string) {
	request := message.ConfigurationGetRequest
	if !isValidConfigurationGetRequest(request) {
		service.publishFailedGetAnswer(correlationId, message.Language)
		return
	}

	log.Info().Str(constants.LogCorrelationId, correlationId).
		Str(constants.LogGuildId, request.GuildId).
		Msgf("Get configuration request received")

	guild, err := service.guildService.Get(request.GuildId)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationId, correlationId).
			Str(constants.LogGuildId, request.GuildId).
			Msgf("Returning failed answer")
		service.publishFailedGetAnswer(correlationId, message.Language)
		return
	}

	service.publishSucceededGetAnswer(correlationId, guild, message.Language)
}

func (service *ConfiguratorServiceImpl) publishSucceededGetAnswer(correlationId string,
	guild entities.Guild, lg amqp.Language) {

	message := amqp.RabbitMQMessage{
		Type:                   amqp.RabbitMQMessage_CONFIGURATION_GET_ANSWER,
		Status:                 amqp.RabbitMQMessage_SUCCESS,
		Language:               lg,
		ConfigurationGetAnswer: mappers.MapGuild(guild),
	}

	err := service.broker.Publish(&message, amqp.ExchangeAnswer, answersRoutingkey, correlationId)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationId, correlationId).
			Msgf("Cannot publish via broker, request ignored")
	}
}

func (service *ConfiguratorServiceImpl) publishFailedGetAnswer(correlationId string, lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_GET_ANSWER,
		Status:   amqp.RabbitMQMessage_FAILED,
		Language: lg,
	}

	err := service.broker.Publish(&message, amqp.ExchangeAnswer, answersRoutingkey, correlationId)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationId, correlationId).
			Msgf("Cannot publish via broker, request ignored")
	}
}

func isValidConfigurationGetRequest(request *amqp.ConfigurationGetRequest) bool {
	return request != nil
}
