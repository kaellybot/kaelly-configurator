package guilds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/guilds"
)

type Service interface {
	Get(guildID string) (entities.Guild, error)
	Save(guild entities.Guild) error
}

type Impl struct {
	guildRepo guilds.Repository
}
