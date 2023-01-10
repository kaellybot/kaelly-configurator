package configurators

import (
	"context"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/services/channels"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
	"github.com/rs/zerolog/log"
)

func New(broker amqp.MessageBrokerInterface, guildService guilds.GuildService,
	channelService channels.ChannelService) (*ConfiguratorServiceImpl, error) {

	return &ConfiguratorServiceImpl{
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

func (service *ConfiguratorServiceImpl) Consume() error {
	log.Info().Msgf("Consuming configurator requests...")
	return service.broker.Consume(requestQueueName, requestsRoutingkey, service.consume)
}

func (service *ConfiguratorServiceImpl) consume(ctx context.Context,
	message *amqp.RabbitMQMessage, correlationId string) {

	switch message.Type {
	case amqp.RabbitMQMessage_CONFIGURATION_GET_REQUEST:
		service.getRequest(message, correlationId)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_REQUEST:
		service.setRequest(message, correlationId)
	default:
		log.Warn().
			Str(constants.LogCorrelationId, correlationId).
			Msgf("Type not recognized, request ignored")
	}
}
