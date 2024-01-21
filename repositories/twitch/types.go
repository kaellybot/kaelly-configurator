package twitch

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type Repository interface {
	Get(guildID, channelID, videastID string) (*entities.WebhookTwitch, error)
	Save(channelWebhook entities.WebhookTwitch) error
	Delete(channelWebhook entities.WebhookTwitch) error
}

type Impl struct {
	db databases.MySQLConnection
}
