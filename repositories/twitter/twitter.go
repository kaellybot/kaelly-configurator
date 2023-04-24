package twitter

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *TwitterWebhookRepositoryImpl {
	return &TwitterWebhookRepositoryImpl{db: db}
}

func (repo *TwitterWebhookRepositoryImpl) Save(webhook entities.WebhookTwitter) error {
	return repo.db.GetDB().Save(&webhook).Error
}
