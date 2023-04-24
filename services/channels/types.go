package channels

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/chanservers"
)

type ChannelService interface {
	SaveChannelServer(channelServer entities.ChannelServer) error
}

type ChannelServiceImpl struct {
	channelServerRepo chanservers.ChannelServerRepository
}
