package guilds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type GuildRepository interface {
	Get(guildId string) (entities.Guild, error)
	Save(guild entities.Guild) error
}

type GuildRepositoryImpl struct {
	db databases.MySQLConnection
}
