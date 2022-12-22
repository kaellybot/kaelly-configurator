package guilds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type GuildRepository interface {
	GetGuild(id string) (entities.Guild, error)
}

type GuildRepositoryImpl struct {
	db databases.MySQLConnection
}

func New(db databases.MySQLConnection) *GuildRepositoryImpl {
	return &GuildRepositoryImpl{db: db}
}

func (repo *GuildRepositoryImpl) GetGuild(id string) (entities.Guild, error) {
	var guild entities.Guild
	response := repo.db.GetDB().First(&guild)
	return guild, response.Error
}
