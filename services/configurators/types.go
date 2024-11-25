package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/services/channels"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
)

const (
	requestQueueName = "configurator-requests"
	newsQueueName    = "configurator-news"

	requestsRoutingkey = "requests.configs"
	newsRoutingkey     = "news.guilds"
)

type Service interface {
	Consume()
}

type Impl struct {
	broker         amqp.MessageBroker
	guildService   guilds.Service
	channelService channels.Service
}
