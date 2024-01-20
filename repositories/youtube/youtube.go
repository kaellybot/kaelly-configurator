package youtube

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(guildID, channelID, videastID string) (*entities.WebhookYoutube, error) {
	var webhook entities.WebhookYoutube
	err := repo.db.GetDB().
		Where("guild_id = ? AND channel_id = ? AND videast_id = ?",
			guildID, channelID, videastID).
		Limit(1).
		Find(&webhook).Error
	if err != nil {
		return nil, err
	}

	if webhook == (entities.WebhookYoutube{}) {
		return nil, nil
	}

	return &webhook, nil
}

func (repo *Impl) Save(webhook entities.WebhookYoutube) error {
	return repo.db.GetDB().Save(&webhook).Error
}

func (repo *Impl) Delete(webhook entities.WebhookYoutube) error {
	return repo.db.GetDB().Delete(&webhook).Error
}
