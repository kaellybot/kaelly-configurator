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
	return guild, repo.db.GetDB().
		Preload("ChannelServers").
		Preload("AlmanaxWebhooks").
		Preload("FeedWebhooks").
		Preload("TwitterWebhooks.TwitterAccount").
		Where(entities.Guild{ID: guildID, Game: game}).
		Find(&guild).Limit(1).Error
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
