package chanwebhooks

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type ChannelWebhookRepository interface {
	Save(channelWebhook entities.ChannelWebhook) error
}

type ChannelWebhookRepositoryImpl struct {
	db databases.MySQLConnection
}

func New(db databases.MySQLConnection) *ChannelWebhookRepositoryImpl {
	return &ChannelWebhookRepositoryImpl{db: db}
}

func (repo *ChannelWebhookRepositoryImpl) Save(channelWebhook entities.ChannelWebhook) error {
	return repo.db.GetDB().Save(&channelWebhook).Error
}
