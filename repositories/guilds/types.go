package guilds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type Repository interface {
	Get(guildID string) (entities.Guild, error)
	Save(guild entities.Guild) error
}

type Impl struct {
	db databases.MySQLConnection
}
