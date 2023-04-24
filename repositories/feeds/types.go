package feeds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type RssWebhookRepository interface {
	Save(channelWebhook entities.WebhookFeed) error
}

type RssWebhookRepositoryImpl struct {
	db databases.MySQLConnection
}
