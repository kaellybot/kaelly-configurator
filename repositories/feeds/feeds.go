package feeds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *FeedWebhookRepositoryImpl {
	return &FeedWebhookRepositoryImpl{db: db}
}

func (repo *FeedWebhookRepositoryImpl) Get(guildId, channelId, feedTypeId string,
	locale amqp.Language) (*entities.WebhookFeed, error) {

	var webhook entities.WebhookFeed
	err := repo.db.GetDB().
		Where("guild_id = ? AND channel_id = ? AND feed_type_id = ? AND locale = ?",
			guildId, channelId, feedTypeId, locale).
		Limit(1).
		Find(&webhook).Error
	if err != nil {
		return nil, err
	}

	if webhook == (entities.WebhookFeed{}) {
		return nil, nil
	}

	return &webhook, nil
}

func (repo *FeedWebhookRepositoryImpl) Save(webhook entities.WebhookFeed) error {
	return repo.db.GetDB().Save(&webhook).Error
}

func (repo *FeedWebhookRepositoryImpl) Delete(webhook entities.WebhookFeed) error {
	return repo.db.GetDB().Delete(&webhook).Error
}
