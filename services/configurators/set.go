package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/rs/zerolog/log"
)

func (service *ConfiguratorServiceImpl) publishSucceededSetWebhookAnswer(correlationId, webhookId string,
	lg amqp.Language) {

	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_SET_ANSWER,
		Status:   amqp.RabbitMQMessage_SUCCESS,
		Language: lg,
		ConfigurationSetAnswer: &amqp.ConfigurationSetAnswer{
			RemoveWebhook: true,
			WebhookId:     webhookId,
		},
	}

	err := service.broker.Publish(&message, amqp.ExchangeAnswer, answersRoutingkey, correlationId)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationId, correlationId).
			Msgf("Cannot publish via broker, request ignored")
	}
}

func (service *ConfiguratorServiceImpl) publishSucceededSetAnswer(correlationId string, lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_SET_ANSWER,
		Status:   amqp.RabbitMQMessage_SUCCESS,
		Language: lg,
		ConfigurationSetAnswer: &amqp.ConfigurationSetAnswer{
			RemoveWebhook: false,
		},
	}

	err := service.broker.Publish(&message, amqp.ExchangeAnswer, answersRoutingkey, correlationId)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationId, correlationId).
			Msgf("Cannot publish via broker, request ignored")
	}
}

func (service *ConfiguratorServiceImpl) publishFailedSetWebhookAnswer(correlationId, webhookId string,
	lg amqp.Language) {

	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_SET_ANSWER,
		Status:   amqp.RabbitMQMessage_FAILED,
		Language: lg,
		ConfigurationSetAnswer: &amqp.ConfigurationSetAnswer{
			RemoveWebhook: true,
			WebhookId:     webhookId,
		},
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
		ConfigurationSetAnswer: &amqp.ConfigurationSetAnswer{
			RemoveWebhook: false,
		},
	}

	err := service.broker.Publish(&message, amqp.ExchangeAnswer, answersRoutingkey, correlationId)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationId, correlationId).
			Msgf("Cannot publish via broker, request ignored")
	}
}
