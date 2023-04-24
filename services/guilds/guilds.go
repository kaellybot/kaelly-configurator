package guilds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/guilds"
)

func New(guildRepo guilds.GuildRepository) (*GuildServiceImpl, error) {
	return &GuildServiceImpl{guildRepo: guildRepo}, nil
}

func (service *GuildServiceImpl) Get(guildId string) (entities.Guild, error) {
	return service.guildRepo.Get(guildId)
}

func (service *GuildServiceImpl) Save(guild entities.Guild) error {
	return service.guildRepo.Save(guild)
}
