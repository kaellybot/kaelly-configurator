package guilds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/guilds"
)

func New(guildRepo guilds.Repository) (*Impl, error) {
	return &Impl{guildRepo: guildRepo}, nil
}

func (service *Impl) Get(guildID string, game amqp.Game) (entities.Guild, error) {
	return service.guildRepo.Get(guildID, game)
}

func (service *Impl) Save(guild entities.Guild) error {
	return service.guildRepo.Save(guild)
}
