package twitch

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(guildID, channelID, streamerID string) (*entities.WebhookTwitch, error) {
	var webhook entities.WebhookTwitch
	err := repo.db.GetDB().
		Where("guild_id = ? AND channel_id = ? AND streamer_id = ?",
			guildID, channelID, streamerID).
		Limit(1).
		Find(&webhook).Error
	if err != nil {
		return nil, err
	}

	if webhook == (entities.WebhookTwitch{}) {
		return nil, nil
	}

	return &webhook, nil
}

func (repo *Impl) Save(webhook entities.WebhookTwitch) error {
	return repo.db.GetDB().Save(&webhook).Error
}

func (repo *Impl) Delete(webhook entities.WebhookTwitch) error {
	return repo.db.GetDB().Delete(&webhook).Error
}
