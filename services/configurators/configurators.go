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
	message *amqp.RabbitMQMessage, correlationID string) {
	//exhaustive:ignore Don't need to be exhaustive here since they will be handled by default case
	switch message.Type {
	case amqp.RabbitMQMessage_CONFIGURATION_GET_REQUEST:
		service.getRequest(message, correlationID)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_SERVER_REQUEST:
		service.serverRequest(message, correlationID)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_ALMANAX_WEBHOOK_REQUEST:
		service.almanaxRequest(message, correlationID)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_RSS_WEBHOOK_REQUEST:
		service.rssRequest(message, correlationID)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_TWITCH_WEBHOOK_REQUEST:
		service.twitchRequest(message, correlationID)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_TWITTER_WEBHOOK_REQUEST:
		service.twitterRequest(message, correlationID)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_YOUTUBE_WEBHOOK_REQUEST:
		service.youtubeRequest(message, correlationID)
	default:
		log.Warn().
			Str(constants.LogCorrelationID, correlationID).
			Msgf("Type not recognized, request ignored")
	}
}
