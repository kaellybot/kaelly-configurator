package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/repositories/almanax"
	"github.com/kaellybot/kaelly-configurator/repositories/feeds"
	guildRepo "github.com/kaellybot/kaelly-configurator/repositories/guilds"
	"github.com/kaellybot/kaelly-configurator/repositories/servers"
	"github.com/kaellybot/kaelly-configurator/repositories/twitter"
	"github.com/kaellybot/kaelly-configurator/services/channels"
	"github.com/kaellybot/kaelly-configurator/services/configurators"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
	"github.com/kaellybot/kaelly-configurator/utils/insights"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Impl, error) {
	// misc
	broker := amqp.New(constants.RabbitMQClientID, viper.GetString(constants.RabbitMQAddress),
		amqp.WithBindings(configurators.GetBindings()...))
	db := databases.New()
	probes := insights.NewProbes(broker.IsConnected, db.IsConnected)
	prom := insights.NewPrometheusMetrics()

	// repositories
	guildRepo := guildRepo.New(db)
	chanServerRepo := servers.New(db)
	almanaxRepo := almanax.New(db)
	feedsRepo := feeds.New(db)
	twitterRepo := twitter.New(db)

	// services
	guildService, err := guilds.New(guildRepo)
	if err != nil {
		return nil, err
	}

	channelService, err := channels.New(chanServerRepo, almanaxRepo, feedsRepo, twitterRepo)
	if err != nil {
		return nil, err
	}

	configService, err := configurators.New(broker, guildService, channelService)
	if err != nil {
		return nil, err
	}

	return &Impl{
		guildService:        guildService,
		channelService:      channelService,
		configuratorService: configService,
		broker:              broker,
		db:                  db,
		probes:              probes,
		prom:                prom,
	}, nil
}

func (app *Impl) Run() error {
	app.probes.ListenAndServe()
	app.prom.ListenAndServe()

	if err := app.db.Run(); err != nil {
		return err
	}

	if err := app.broker.Run(); err != nil {
		return err
	}

	app.configuratorService.Consume()
	return nil
}

func (app *Impl) Shutdown() {
	app.broker.Shutdown()
	app.db.Shutdown()
	app.prom.Shutdown()
	app.probes.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
