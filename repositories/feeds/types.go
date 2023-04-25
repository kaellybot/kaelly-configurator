package feeds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type FeedWebhookRepository interface {
	Get(guildId, channelId, feedTypeId string, locale amqp.Language) (*entities.WebhookFeed, error)
	Save(channelWebhook entities.WebhookFeed) error
	Delete(channelWebhook entities.WebhookFeed) error
}

type FeedWebhookRepositoryImpl struct {
	db databases.MySQLConnection
}
