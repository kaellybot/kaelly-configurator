package feeds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *FeedWebhookRepositoryImpl {
	return &FeedWebhookRepositoryImpl{db: db}
}

func (repo *FeedWebhookRepositoryImpl) Save(webhook entities.WebhookFeed) error {
	return repo.db.GetDB().Save(&webhook).Error
}
