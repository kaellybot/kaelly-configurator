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

func New(db databases.MySQLConnection) *GuildRepositoryImpl {
	return &GuildRepositoryImpl{db: db}
}

func (repo *GuildRepositoryImpl) Get(guildId string) (entities.Guild, error) {
	var guild entities.Guild
	return guild, repo.db.GetDB().
		Preload("ChannelServers").
		Preload("AlmanaxWebhooks").
		Preload("RssWebhooks").
		Preload("TwitterWebhooks").
		First(&guild).Error
}

func (repo *GuildRepositoryImpl) Save(guild entities.Guild) error {
	return repo.db.GetDB().Save(&guild).Error
}
