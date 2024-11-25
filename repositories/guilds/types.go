package guilds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type Repository interface {
	Get(guildID string, game amqp.Game) (entities.Guild, error)
	Create(guildID string, game amqp.Game) error
	Update(guild entities.Guild) error
	Delete(guildID string, game amqp.Game) error
}

type Impl struct {
	db databases.MySQLConnection
}
