package guilds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/guilds"
)

type GuildService interface {
	Get(guildId string) (entities.Guild, error)
	Save(guild entities.Guild) error
}

type GuildServiceImpl struct {
	guildRepo guilds.GuildRepository
}

func New(guildRepo guilds.GuildRepository) (*GuildServiceImpl, error) {
	return &GuildServiceImpl{guildRepo: guildRepo}, nil
}

func (service *GuildServiceImpl) Get(guildId string) (entities.Guild, error) {
	return service.guildRepo.Get(guildId)
}

func (service *GuildServiceImpl) Save(guild entities.Guild) error {
	return service.guildRepo.Save(guild)
}
