package configurators

import (
	"errors"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/services/channels"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
)

const (
	requestQueueName   = "configurator-requests"
	requestsRoutingkey = "requests.configs"
	answersRoutingkey  = "answers.configs"
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
