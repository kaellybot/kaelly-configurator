package webhooktwitter

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type TwitterWebhookRepository interface {
	Save(channelWebhook entities.TwitterWebhook) error
}

type TwitterWebhookRepositoryImpl struct {
	db databases.MySQLConnection
}
