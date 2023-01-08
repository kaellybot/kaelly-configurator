package guilds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
	"gorm.io/gorm/clause"
)

type GuildRepository interface {
	Save(guild entities.Guild) error
}

type GuildRepositoryImpl struct {
	db databases.MySQLConnection
}

func New(db databases.MySQLConnection) *GuildRepositoryImpl {
	return &GuildRepositoryImpl{db: db}
}

func (repo *GuildRepositoryImpl) Save(guild entities.Guild) error {
	return repo.db.GetDB().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&guild).Error
}
