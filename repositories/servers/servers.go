package servers

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Save(channelServer entities.ChannelServer) error {
	return repo.db.GetDB().Save(&channelServer).Error
}
