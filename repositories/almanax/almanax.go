package almanax

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(guildID, channelID string, locale amqp.Language, game amqp.Game,
) (*entities.WebhookAlmanax, error) {
	var webhook entities.WebhookAlmanax
	err := repo.db.GetDB().
		Where("guild_id = ? AND channel_id = ? AND locale = ? AND game = ?",
			guildID, channelID, locale, game).
		Limit(1).
		Find(&webhook).Error
	if err != nil {
		return nil, err
	}

	if webhook == (entities.WebhookAlmanax{}) {
		return &webhook, nil
	}

	return &webhook, nil
}

func (repo *Impl) Save(webhook entities.WebhookAlmanax) error {
	return repo.db.GetDB().Save(&webhook).Error
}

func (repo *Impl) Delete(webhook entities.WebhookAlmanax) error {
	return repo.db.GetDB().Delete(&webhook).Error
}
