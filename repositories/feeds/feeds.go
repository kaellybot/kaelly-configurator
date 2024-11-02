package feeds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(guildID, channelID, feedTypeID string,
	locale amqp.Language, game amqp.Game) (*entities.WebhookFeed, error) {
	var webhook entities.WebhookFeed
	err := repo.db.GetDB().
		Where("guild_id = ? AND channel_id = ? AND feed_type_id = ? AND locale = ? AND game = ?",
			guildID, channelID, feedTypeID, locale, game).
		Limit(1).
		Find(&webhook).Error
	if err != nil {
		return nil, err
	}

	if webhook == (entities.WebhookFeed{}) {
		return &webhook, nil
	}

	return &webhook, nil
}

func (repo *Impl) Save(webhook entities.WebhookFeed) error {
	return repo.db.GetDB().Save(&webhook).Error
}

func (repo *Impl) Delete(webhook entities.WebhookFeed) error {
	return repo.db.GetDB().Delete(&webhook).Error
}
