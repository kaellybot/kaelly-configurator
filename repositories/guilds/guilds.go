package guilds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *GuildRepositoryImpl {
	return &GuildRepositoryImpl{db: db}
}

func (repo *GuildRepositoryImpl) Get(guildId string) (entities.Guild, error) {
	var guild entities.Guild
	return guild, repo.db.GetDB().
		Preload("ChannelServers").
		Preload("AlmanaxWebhooks").
		Preload("RssWebhooks").
		Preload("TwitterWebhooks.Twitter").
		First(&guild).Error
}

func (repo *GuildRepositoryImpl) Save(guild entities.Guild) error {
	return repo.db.GetDB().Save(&guild).Error
}
