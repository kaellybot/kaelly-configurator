package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/rs/zerolog/log"
)

func (service *ConfiguratorServiceImpl) setRequest(message *amqp.RabbitMQMessage, correlationId string) {
	request := message.GetConfigurationSetRequest()
	if !isValidConfigurationSetRequest(request) {
		service.publishFailedSetAnswer(correlationId, message.Language)
		return
	}

	switch request.Field {
	case amqp.ConfigurationSetRequest_WEBHOOK:
		service.webhookRequest(correlationId, request, message.Language)
	case amqp.ConfigurationSetRequest_SERVER:
		service.serverRequest(correlationId, request, message.Language)
	default:
		log.Error().
			Str(constants.LogCorrelationId, correlationId).
			Msgf("Config field not recognized, request failed")
		service.publishFailedSetAnswer(correlationId, message.Language)
	}
}

func (service *ConfiguratorServiceImpl) publishSucceededSetAnswer(correlationId string, lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_SET_ANSWER,
		Status:   amqp.RabbitMQMessage_SUCCESS,
		Language: lg,
	}

	err := service.broker.Publish(&message, amqp.ExchangeAnswer, answersRoutingkey, correlationId)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationId, correlationId).
			Msgf("Cannot publish via broker, request ignored")
	}
}

func (service *ConfiguratorServiceImpl) publishFailedSetAnswer(correlationId string, lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_SET_ANSWER,
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

func isValidConfigurationSetRequest(request *amqp.ConfigurationSetRequest) bool {
	return request != nil
}
