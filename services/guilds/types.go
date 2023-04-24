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
