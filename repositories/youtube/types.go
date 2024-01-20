package youtube

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type Repository interface {
	Get(guildID, channelID, videastID string) (*entities.WebhookYoutube, error)
	Save(channelWebhook entities.WebhookYoutube) error
	Delete(channelWebhook entities.WebhookYoutube) error
}

type Impl struct {
	db databases.MySQLConnection
}
