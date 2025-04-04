package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/services/channels"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
	"github.com/rs/zerolog/log"
)

func New(broker amqp.MessageBroker, guildService guilds.Service, channelService channels.Service) (*Impl, error) {
	return &Impl{
		guildService:   guildService,
		channelService: channelService,
		broker:         broker,
	}, nil
}

func GetBindings() []amqp.Binding {
	return []amqp.Binding{
		{
			Exchange:   amqp.ExchangeRequest,
			RoutingKey: requestsRoutingkey,
			Queue:      requestQueueName,
		},
		{
			Exchange:   amqp.ExchangeNews,
			RoutingKey: newsRoutingkey,
			Queue:      newsQueueName,
		},
	}
}

func (service *Impl) Consume() {
	log.Info().Msgf("Consuming configurator news and requests...")
	service.broker.Consume(requestQueueName, service.consumeRequests)
	service.broker.Consume(newsQueueName, service.consumeNews)
}

func (service *Impl) consumeRequests(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	//exhaustive:ignore Don't need to be exhaustive here since they will be handled by default case
	switch message.Type {
	case amqp.RabbitMQMessage_CONFIGURATION_GET_REQUEST:
		service.getRequest(ctx, message)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_SERVER_REQUEST:
		service.serverRequest(ctx, message)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_NOTIFICATION_REQUEST:
		service.notificationRequest(ctx, message)
	default:
		log.Warn().
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msgf("Type not recognized, request ignored")
	}
}

func (service *Impl) notificationRequest(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	request := message.ConfigurationSetNotificationRequest
	if !isValidNotificationRequest(request) {
		service.publishFailedSetNotificationAnswer(ctx, "", message.Language)
		return
	}

	switch request.NotificationType {
	case amqp.NotificationType_ALMANAX:
		service.almanaxRequest(ctx, message)
	case amqp.NotificationType_RSS:
		service.rssRequest(ctx, message)
	case amqp.NotificationType_TWITTER:
		service.twitterRequest(ctx, message)
	case amqp.NotificationType_UNKNOWN:
		fallthrough
	default:
		log.Warn().
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msgf("Notification type not recognized, request ignored")
	}
}

func (service *Impl) consumeNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	//exhaustive:ignore Don't need to be exhaustive here since they will be handled by default case
	switch message.Type {
	case amqp.RabbitMQMessage_NEWS_GUILD:
		service.guildNews(message)
	default:
		log.Warn().
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msgf("Type not recognized, news ignored")
	}
}

func isValidNotificationRequest(request *amqp.ConfigurationSetNotificationRequest) bool {
	return request != nil
}
