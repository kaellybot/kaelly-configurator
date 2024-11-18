package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/services/channels"
	"github.com/kaellybot/kaelly-configurator/services/configurators"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
	"github.com/kaellybot/kaelly-configurator/utils/insights"
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
	db                  databases.MySQLConnection
	probes              insights.Probes
	prom                insights.PrometheusMetrics
}
