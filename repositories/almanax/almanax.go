package almanax

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *AlmanaxWebhookRepositoryImpl {
	return &AlmanaxWebhookRepositoryImpl{db: db}
}

func (repo *AlmanaxWebhookRepositoryImpl) Save(webhook entities.WebhookAlmanax) error {
	return repo.db.GetDB().Save(&webhook).Error
}
