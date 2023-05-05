package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/services/channels"
	"github.com/kaellybot/kaelly-configurator/services/configurators"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
)

type Application interface {
	Run() error
	Shutdown()
}

type Impl struct {
	guildService        guilds.Service
	channelService      channels.Service
	configuratorService configurators.Service
	broker              amqp.MessageBroker
}
