package configurators

import (
	"context"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
	"github.com/kaellybot/kaelly-configurator/services/servers"
	"github.com/rs/zerolog/log"
)

const (
	requestQueueName   = "configurator-requests"
	requestsRoutingkey = "requests.configurator"
	answersRoutingkey  = "answers.configurator"
)

type ConfiguratorService interface {
	Consume() error
}

type ConfiguratorServiceImpl struct {
	broker        amqp.MessageBrokerInterface
	serverService servers.ServerService
	guildService  guilds.GuildService
}

func New(broker amqp.MessageBrokerInterface, serverService servers.ServerService,
	guildService guilds.GuildService) (*ConfiguratorServiceImpl, error) {

	return &ConfiguratorServiceImpl{
		serverService: serverService,
		guildService:  guildService,
		broker:        broker,
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

	// TODO
}
