package guilds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(guildID string) (entities.Guild, error) {
	var guild entities.Guild
	return guild, repo.db.GetDB().
		Preload("ChannelServers").
		Preload("AlmanaxWebhooks").
		Preload("FeedWebhooks").
		Preload("TwitchWebhooks").
		Preload("TwitterWebhooks.TwitterAccount").
		Preload("YoutubeWebhooks").
		First(&guild).Error
}

func (repo *Impl) Save(guild entities.Guild) error {
	return repo.db.GetDB().Save(&guild).Error
}
