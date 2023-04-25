package twitter

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *TwitterWebhookRepositoryImpl {
	return &TwitterWebhookRepositoryImpl{db: db}
}

func (repo *TwitterWebhookRepositoryImpl) Get(guildId, channelId string, locale amqp.Language) (*entities.WebhookTwitter, error) {
	var webhook entities.WebhookTwitter
	err := repo.db.GetDB().
		Where("guild_id = ? AND channel_id = ? AND locale = ?", guildId, channelId, locale).
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

func (repo *TwitterWebhookRepositoryImpl) Save(webhook entities.WebhookTwitter) error {
	return repo.db.GetDB().Save(&webhook).Error
}

func (repo *TwitterWebhookRepositoryImpl) Delete(webhook entities.WebhookTwitter) error {
	return repo.db.GetDB().Delete(&webhook).Error
}
