package twitter

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type Repository interface {
	Get(guildID, channelID, twitterID string) (*entities.WebhookTwitter, error)
	Save(channelWebhook entities.WebhookTwitter) error
	Delete(channelWebhook entities.WebhookTwitter) error
}

type Impl struct {
	db databases.MySQLConnection
}
