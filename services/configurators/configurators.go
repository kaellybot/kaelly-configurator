package configurators

import (
	"context"
	"errors"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/services/channels"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
	"github.com/rs/zerolog/log"
)

const (
	requestQueueName   = "configurator-requests"
	requestsRoutingkey = "requests.configurator"
	answersRoutingkey  = "answers.configurator"
)

var (
	errInvalidMessage = errors.New("Invalid configurator request message, probably badly built")
)

type ConfiguratorService interface {
	Consume() error
}

type ConfiguratorServiceImpl struct {
	broker         amqp.MessageBrokerInterface
	guildService   guilds.GuildService
	channelService channels.ChannelService
}

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

	if !isValidConfiguratorRequest(message) {
		log.Error().Err(errInvalidMessage).
			Str(constants.LogCorrelationId, correlationId).
			Msgf("Cannot treat request, returning failed message")
		service.publishFailedAnswer(correlationId)
		return
	}

	request := message.GetConfigurationRequest()

	switch request.Field {
	case amqp.ConfigurationRequest_WEBHOOK:
		// TODO
	case amqp.ConfigurationRequest_SERVER:
		service.serverRequest(correlationId, request.GuildId, request.ChannelId, request.ServerField.ServerId)
	default:
		log.Warn().Str(constants.LogCorrelationId, correlationId).Msgf("Config field not recognized, request ignored")
	}
}

func isValidConfiguratorRequest(message *amqp.RabbitMQMessage) bool {
	return message.Type == amqp.RabbitMQMessage_CONFIGURATION_REQUEST && message.GetConfigurationRequest() != nil
}

func (service *ConfiguratorServiceImpl) publishSucceededAnswer(correlationId string) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_ANSWER,
		Status:   amqp.RabbitMQMessage_SUCCESS,
		Language: amqp.RabbitMQMessage_ANY,
	}

	err := service.broker.Publish(&message, amqp.ExchangeAnswer, answersRoutingkey, correlationId)
	if err != nil {
		log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).Msgf("Cannot publish via broker, request ignored")
	}
}

func (service *ConfiguratorServiceImpl) publishFailedAnswer(correlationId string) {
	message := amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_ANSWER,
		Status:   amqp.RabbitMQMessage_FAILED,
		Language: amqp.RabbitMQMessage_ANY,
	}

	err := service.broker.Publish(&message, amqp.ExchangeAnswer, answersRoutingkey, correlationId)
	if err != nil {
		log.Error().Err(err).Str(constants.LogCorrelationId, correlationId).Msgf("Cannot publish via broker, request ignored")
	}
}
