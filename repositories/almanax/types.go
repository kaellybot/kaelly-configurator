package almanax

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type Repository interface {
	Get(guildID, channelID string, locale amqp.Language,
		game amqp.Game) (*entities.WebhookAlmanax, error)
	Save(channelWebhook entities.WebhookAlmanax) error
	Delete(channelWebhook entities.WebhookAlmanax) error
}

type Impl struct {
	db databases.MySQLConnection
}
