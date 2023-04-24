package webhookalmanax

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type AlmanaxWebhookRepository interface {
	Save(channelWebhook entities.AlmanaxWebhook) error
}

type AlmanaxWebhookRepositoryImpl struct {
	db databases.MySQLConnection
}
