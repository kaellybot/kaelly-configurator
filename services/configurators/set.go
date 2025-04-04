package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/utils/replies"
	"github.com/rs/zerolog/log"
)

func (service *Impl) publishSucceededSetNotificationAnswer(ctx amqp.Context, webhookID string,
	lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_SET_NOTIFICATION_ANSWER,
		Status:   amqp.RabbitMQMessage_SUCCESS,
		Language: lg,
		ConfigurationSetNotificationAnswer: &amqp.ConfigurationSetNotificationAnswer{
			RemoveWebhook: webhookID != "",
			WebhookId:     webhookID,
		},
	}

	replies.SucceededAnswer(ctx, service.broker, &message)
}

func (service *Impl) publishFailedSetNotificationAnswer(ctx amqp.Context, webhookID string,
	lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_SET_NOTIFICATION_ANSWER,
		Status:   amqp.RabbitMQMessage_FAILED,
		Language: lg,
		ConfigurationSetNotificationAnswer: &amqp.ConfigurationSetNotificationAnswer{
			RemoveWebhook: webhookID != "",
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

func (service *Impl) publishSucceededSetServerAnswer(ctx amqp.Context, lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_SET_SERVER_ANSWER,
		Status:   amqp.RabbitMQMessage_SUCCESS,
		Language: lg,
	}

	replies.SucceededAnswer(ctx, service.broker, &message)
}

func (service *Impl) publishFailedSetServerAnswer(ctx amqp.Context, lg amqp.Language) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_SET_SERVER_ANSWER,
		Status:   amqp.RabbitMQMessage_FAILED,
		Language: lg,
	}

	err := service.broker.Reply(&message, ctx.CorrelationID, ctx.ReplyTo)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msgf("Cannot publish via broker, request ignored")
	}
}
