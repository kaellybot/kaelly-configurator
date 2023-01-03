package chanservers

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
	"gorm.io/gorm/clause"
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
	response := repo.db.GetDB().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Omit("Guild", "Server").Create(&channelServer)

	return response.Error
}
