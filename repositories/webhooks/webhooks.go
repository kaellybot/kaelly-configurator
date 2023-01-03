package webhooks

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type WebhookRepository interface {
	GetGuild(id string) (entities.Webhook, error)
}

type WebhookRepositoryImpl struct {
	db databases.MySQLConnection
}

func New(db databases.MySQLConnection) *WebhookRepositoryImpl {
	return &WebhookRepositoryImpl{db: db}
}

// TODO

func (repo *WebhookRepositoryImpl) GetGuild(id string) (entities.Webhook, error) {
	var guild entities.Webhook
	response := repo.db.GetDB().First(&guild)
	return guild, response.Error
}
