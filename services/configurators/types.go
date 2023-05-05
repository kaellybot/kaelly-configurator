package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/services/channels"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
)

const (
	requestQueueName   = "configurator-requests"
	requestsRoutingkey = "requests.configs"
	answersRoutingkey  = "answers.configs"
)

type Service interface {
	Consume() error
}

type Impl struct {
	broker         amqp.MessageBroker
	guildService   guilds.Service
	channelService channels.Service
}
