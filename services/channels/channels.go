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

func New(channelServerRepo chanservers.ChannelServerRepository) (*ChannelServiceImpl, error) {
	return &ChannelServiceImpl{channelServerRepo: channelServerRepo}, nil
}

func (service *ChannelServiceImpl) SaveChannelServer(channelServer entities.ChannelServer) error {
	return service.channelServerRepo.Save(channelServer)
}
