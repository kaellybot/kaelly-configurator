package twitter

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type TwitterWebhookRepository interface {
	Save(channelWebhook entities.WebhookTwitter) error
}

type TwitterWebhookRepositoryImpl struct {
	db databases.MySQLConnection
}
