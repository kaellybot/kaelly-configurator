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
