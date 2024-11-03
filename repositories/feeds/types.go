package feeds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type Repository interface {
	Get(guildID, channelID, feedTypeID string, game amqp.Game,
	) (*entities.WebhookFeed, error)
	Save(channelWebhook entities.WebhookFeed) error
	Delete(channelWebhook entities.WebhookFeed) error
}

type Impl struct {
	db databases.MySQLConnection
}
