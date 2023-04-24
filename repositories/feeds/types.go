package feeds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type FeedWebhookRepository interface {
	Save(channelWebhook entities.WebhookFeed) error
}

type FeedWebhookRepositoryImpl struct {
	db databases.MySQLConnection
}
