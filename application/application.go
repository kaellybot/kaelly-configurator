package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	guildRepo "github.com/kaellybot/kaelly-configurator/repositories/guilds"
	serverRepo "github.com/kaellybot/kaelly-configurator/repositories/servers"
	"github.com/kaellybot/kaelly-configurator/services/configurators"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
	"github.com/kaellybot/kaelly-configurator/services/servers"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type ApplicationInterface interface {
	Run() error
	Shutdown()
}

type Application struct {
	serverService       servers.ServerService
	guildService        guilds.GuildService
	configuratorService configurators.ConfiguratorService
	broker              amqp.MessageBrokerInterface
}

func New() (*Application, error) {
	// misc
	db, err := databases.New()
	if err != nil {
		return nil, err
	}

	broker, err := amqp.New(constants.RabbitMQClientId, viper.GetString(constants.RabbitMqAddress),
		[]amqp.Binding{configurators.GetBinding()})
	if err != nil {
		return nil, err
	}

	// repositories
	serverRepo := serverRepo.New(db)
	guildRepo := guildRepo.New(db)

	// services
	serverService, err := servers.New(serverRepo)
	if err != nil {
		return nil, err
	}

	guildService, err := guilds.New(guildRepo)
	if err != nil {
		return nil, err
	}

	configService, err := configurators.New(broker, serverService, guildService)
	if err != nil {
		return nil, err
	}

	return &Application{
		serverService:       serverService,
		guildService:        guildService,
		configuratorService: configService,
		broker:              broker,
	}, nil
}

func (app *Application) Run() error {
	return app.configuratorService.Consume()
}

func (app *Application) Shutdown() {
	app.broker.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
