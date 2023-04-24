package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/services/channels"
	"github.com/kaellybot/kaelly-configurator/services/configurators"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
)

type ApplicationInterface interface {
	Run() error
	Shutdown()
}

type Application struct {
	guildService        guilds.GuildService
	channelService      channels.ChannelService
	configuratorService configurators.ConfiguratorService
	broker              amqp.MessageBrokerInterface
}
