package twitter

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type TwitterWebhookRepository interface {
	Get(guildId, channelId string, locale amqp.Language) (*entities.WebhookTwitter, error)
	Save(channelWebhook entities.WebhookTwitter) error
	Delete(channelWebhook entities.WebhookTwitter) error
}

type TwitterWebhookRepositoryImpl struct {
	db databases.MySQLConnection
}
