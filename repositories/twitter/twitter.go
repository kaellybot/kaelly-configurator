package twitter

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(guildID, channelID string, locale amqp.Language) (*entities.WebhookTwitter, error) {
	var webhook entities.WebhookTwitter
	err := repo.db.GetDB().
		Where("guild_id = ? AND channel_id = ? AND locale = ?", guildID, channelID, locale).
		Limit(1).
		Find(&webhook).Error
	if err != nil {
		return nil, err
	}

	if webhook == (entities.WebhookTwitter{}) {
		return nil, nil
	}

	return &webhook, nil
}

func (repo *Impl) Save(webhook entities.WebhookTwitter) error {
	return repo.db.GetDB().Save(&webhook).Error
}

func (repo *Impl) Delete(webhook entities.WebhookTwitter) error {
	return repo.db.GetDB().Delete(&webhook).Error
}
