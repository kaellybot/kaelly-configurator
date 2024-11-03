package twitter

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(guildID, channelID, twitterID string) (*entities.WebhookTwitter, error) {
	var webhook entities.WebhookTwitter
	err := repo.db.GetDB().
		Where("guild_id = ? AND channel_id = ? AND twitter_id = ?",
			guildID, channelID, twitterID).
		Limit(1).
		Find(&webhook).Error
	if err != nil {
		return nil, err
	}

	if webhook == (entities.WebhookTwitter{}) {
		return &webhook, nil
	}

	return &webhook, nil
}

func (repo *Impl) Save(webhook entities.WebhookTwitter) error {
	return repo.db.GetDB().Save(&webhook).Error
}

func (repo *Impl) Delete(webhook entities.WebhookTwitter) error {
	if webhook != (entities.WebhookTwitter{}) {
		return repo.db.GetDB().Delete(&webhook).Error
	}

	return nil
}
