package almanax

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type AlmanaxWebhookRepository interface {
	Get(guildId, channelId string, locale amqp.Language) (*entities.WebhookAlmanax, error)
	Save(channelWebhook entities.WebhookAlmanax) error
	Delete(channelWebhook entities.WebhookAlmanax) error
}

type AlmanaxWebhookRepositoryImpl struct {
	db databases.MySQLConnection
}
