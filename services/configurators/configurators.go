package configurators

import (
	"context"

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

func GetBinding() amqp.Binding {
	return amqp.Binding{
		Exchange:   amqp.ExchangeRequest,
		RoutingKey: requestsRoutingkey,
		Queue:      requestQueueName,
	}
}

func (service *Impl) Consume() error {
	log.Info().Msgf("Consuming configurator requests...")
	return service.broker.Consume(requestQueueName, service.consume)
}

func (service *Impl) consume(_ context.Context,
	message *amqp.RabbitMQMessage, correlationId string) {
	//exhaustive:ignore Don't need to be exhaustive here since they will be handled by default case
	switch message.Type {
	case amqp.RabbitMQMessage_CONFIGURATION_GET_REQUEST:
		service.getRequest(message, correlationId)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_SERVER_REQUEST:
		service.serverRequest(message, correlationId)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_ALMANAX_WEBHOOK_REQUEST:
		service.almanaxRequest(message, correlationId)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_RSS_WEBHOOK_REQUEST:
		service.rssRequest(message, correlationId)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_TWITCH_WEBHOOK_REQUEST:
		service.twitchRequest(message, correlationId)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_TWITTER_WEBHOOK_REQUEST:
		service.twitterRequest(message, correlationId)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_YOUTUBE_WEBHOOK_REQUEST:
		service.youtubeRequest(message, correlationId)
	default:
		log.Warn().
			Str(constants.LogCorrelationId, correlationId).
			Msgf("Type not recognized, request ignored")
	}
}
