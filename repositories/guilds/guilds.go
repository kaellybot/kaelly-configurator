package guilds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
	"gorm.io/gorm"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(guildID string, game amqp.Game) (entities.Guild, error) {
	guild := entities.Guild{
		ID: guildID,
	}
	// Every preload filters on the game: without it, a guild configured for both
	// games would return the other bot's channels and webhooks.
	return guild, repo.db.GetDB().
		Preload("ChannelServers", "game = ?", game).
		Preload("AlmanaxWebhooks", "game = ?", game).
		Preload("FeedWebhooks", "game = ?", game).
		Preload("TwitterWebhooks", "game = ?", game).
		Preload("TwitterWebhooks.TwitterAccount").
		Where(entities.Guild{ID: guildID, Game: game}).
		Find(&guild).Error
}

func (repo *Impl) Create(id string, game amqp.Game) error {
	return repo.db.GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.FirstOrCreate(&entities.Guild{
			ID:   id,
			Game: game,
		}).Error
	})
}

func (repo *Impl) Update(guild entities.Guild) error {
	return repo.db.GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.Save(&guild).Error
	})
}

func (repo *Impl) Delete(id string, game amqp.Game) error {
	return repo.db.GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&entities.Guild{
			ID:   id,
			Game: game,
		}).Error
	})
}
