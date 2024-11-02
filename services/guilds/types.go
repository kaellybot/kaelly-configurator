package guilds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/guilds"
)

type Service interface {
	Get(guildID string, game amqp.Game) (entities.Guild, error)
	Save(guild entities.Guild) error
}

type Impl struct {
	guildRepo guilds.Repository
}
