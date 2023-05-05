package servers

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

type Repository interface {
	Save(channelServer entities.ChannelServer) error
}

type Impl struct {
	db databases.MySQLConnection
}
