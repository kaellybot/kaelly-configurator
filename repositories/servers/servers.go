package servers

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *ChannelServerRepositoryImpl {
	return &ChannelServerRepositoryImpl{db: db}
}

func (repo *ChannelServerRepositoryImpl) Save(channelServer entities.ChannelServer) error {
	return repo.db.GetDB().Save(&channelServer).Error
}
