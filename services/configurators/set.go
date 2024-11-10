package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/utils/replies"
	"github.com/rs/zerolog/log"
)

func (service *Impl) publishSucceededSetWebhookAnswer(ctx amqp.Context, webhookID string,
	lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_SET_ANSWER,
		Status:   amqp.RabbitMQMessage_SUCCESS,
		Language: lg,
		ConfigurationSetAnswer: &amqp.ConfigurationSetAnswer{
			RemoveWebhook: true,
			WebhookId:     webhookID,
		},
	}

	replies.SucceededAnswer(ctx, service.broker, &message)
}

func (service *Impl) publishSucceededSetAnswer(ctx amqp.Context, lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_SET_ANSWER,
		Status:   amqp.RabbitMQMessage_SUCCESS,
		Language: lg,
		ConfigurationSetAnswer: &amqp.ConfigurationSetAnswer{
			RemoveWebhook: false,
		},
	}

	replies.SucceededAnswer(ctx, service.broker, &message)
}

func (service *Impl) publishFailedSetWebhookAnswer(ctx amqp.Context, webhookID string,
	lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_SET_ANSWER,
		Status:   amqp.RabbitMQMessage_FAILED,
		Language: lg,
		ConfigurationSetAnswer: &amqp.ConfigurationSetAnswer{
			RemoveWebhook: true,
			WebhookId:     webhookID,
		},
	}

	err := service.broker.Reply(&message, ctx.CorrelationID, ctx.ReplyTo)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msgf("Cannot publish via broker, request ignored")
	}
}

func (service *Impl) publishFailedSetAnswer(ctx amqp.Context, lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_SET_ANSWER,
		Status:   amqp.RabbitMQMessage_FAILED,
		Language: lg,
		ConfigurationSetAnswer: &amqp.ConfigurationSetAnswer{
			RemoveWebhook: false,
		},
	}

	err := service.broker.Reply(&message, ctx.CorrelationID, ctx.ReplyTo)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msgf("Cannot publish via broker, request ignored")
	}
}
