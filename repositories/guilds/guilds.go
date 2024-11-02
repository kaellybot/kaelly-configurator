package guilds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(guildID string, game amqp.Game) (entities.Guild, error) {
	guild := entities.Guild{
		ID: guildID,
	}
	return guild, repo.db.GetDB().
		Preload("ChannelServers").
		Preload("AlmanaxWebhooks").
		Preload("FeedWebhooks").
		Preload("TwitchWebhooks").
		Preload("TwitterWebhooks.TwitterAccount").
		Preload("YoutubeWebhooks").
		Where(entities.Guild{ID: guildID, Game: game}).
		Find(&guild).Limit(1).Error
}

func (repo *Impl) Save(guild entities.Guild) error {
	return repo.db.GetDB().Save(&guild).Error
}
