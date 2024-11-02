package guilds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type Repository interface {
	Get(guildID string, game amqp.Game) (entities.Guild, error)
	Save(guild entities.Guild) error
}

type Impl struct {
	db databases.MySQLConnection
}
