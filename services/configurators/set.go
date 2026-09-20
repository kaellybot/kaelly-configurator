package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/utils/replies"
)

func (service *Impl) publishSucceededSetNotificationAnswer(ctx amqp.Context,
	request *amqp.RabbitMQMessage, webhookID string) {
	service.publishSetNotificationAnswer(ctx, request, webhookID, amqp.RabbitMQMessage_SUCCESS)
}

func (service *Impl) publishFailedSetNotificationAnswer(ctx amqp.Context,
	request *amqp.RabbitMQMessage, webhookID string) {
	service.publishSetNotificationAnswer(ctx, request, webhookID, amqp.RabbitMQMessage_FAILED)
}

func (service *Impl) publishSetNotificationAnswer(ctx amqp.Context,
	request *amqp.RabbitMQMessage, webhookID string, status amqp.RabbitMQMessage_Status) {
	message := amqp.NewReply(request,
		amqp.RabbitMQMessage_CONFIGURATION_SET_NOTIFICATION_ANSWER, status)
	message.ConfigurationSetNotificationAnswer = &amqp.ConfigurationSetNotificationAnswer{
		RemoveWebhook: webhookID != "",
		WebhookId:     webhookID,
	}

	replies.Reply(ctx, service.broker, message)
}

func (service *Impl) publishSucceededSetServerAnswer(ctx amqp.Context,
	request *amqp.RabbitMQMessage) {
	replies.Reply(ctx, service.broker, amqp.NewReply(request,
		amqp.RabbitMQMessage_CONFIGURATION_SET_SERVER_ANSWER, amqp.RabbitMQMessage_SUCCESS))
}

func (service *Impl) publishFailedSetServerAnswer(ctx amqp.Context,
	request *amqp.RabbitMQMessage) {
	replies.FailedAnswer(ctx, service.broker, request,
		amqp.RabbitMQMessage_CONFIGURATION_SET_SERVER_ANSWER)
}
