package webhookrss

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type RssWebhookRepository interface {
	Save(channelWebhook entities.RssWebhook) error
}

type RssWebhookRepositoryImpl struct {
	db databases.MySQLConnection
}
