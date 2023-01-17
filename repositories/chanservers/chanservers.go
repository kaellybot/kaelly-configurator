package chanservers

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type ChannelServerRepository interface {
	Save(channelServer entities.ChannelServer) error
}

type ChannelServerRepositoryImpl struct {
	db databases.MySQLConnection
}

func New(db databases.MySQLConnection) *ChannelServerRepositoryImpl {
	return &ChannelServerRepositoryImpl{db: db}
}

func (repo *ChannelServerRepositoryImpl) Save(channelServer entities.ChannelServer) error {
	return repo.db.GetDB().Save(&channelServer).Error
}
