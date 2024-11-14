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

func GetBinding() amqp.Binding {
	return amqp.Binding{
		Exchange:   amqp.ExchangeRequest,
		RoutingKey: requestsRoutingkey,
		Queue:      requestQueueName,
	}
}

func (service *Impl) Consume() {
	log.Info().Msgf("Consuming configurator requests...")
	service.broker.Consume(requestQueueName, service.consume)
}

func (service *Impl) consume(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	//exhaustive:ignore Don't need to be exhaustive here since they will be handled by default case
	switch message.Type {
	case amqp.RabbitMQMessage_CONFIGURATION_GET_REQUEST:
		service.getRequest(ctx, message)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_SERVER_REQUEST:
		service.serverRequest(ctx, message)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_ALMANAX_WEBHOOK_REQUEST:
		service.almanaxRequest(ctx, message)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_RSS_WEBHOOK_REQUEST:
		service.rssRequest(ctx, message)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_TWITCH_WEBHOOK_REQUEST:
		service.twitchRequest(ctx, message)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_TWITTER_WEBHOOK_REQUEST:
		service.twitterRequest(ctx, message)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_YOUTUBE_WEBHOOK_REQUEST:
		service.youtubeRequest(ctx, message)
	default:
		log.Warn().
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msgf("Type not recognized, request ignored")
	}
}
