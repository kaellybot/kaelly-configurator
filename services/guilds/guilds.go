package guilds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/guilds"
)

func New(guildRepo guilds.Repository) (*Impl, error) {
	return &Impl{guildRepo: guildRepo}, nil
}

func (service *Impl) Get(guildID string) (entities.Guild, error) {
	return service.guildRepo.Get(guildID)
}

func (service *Impl) Save(guild entities.Guild) error {
	return service.guildRepo.Save(guild)
}
