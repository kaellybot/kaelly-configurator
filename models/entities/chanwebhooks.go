package entities

import "github.com/kaellybot/kaelly-configurator/models/constants"

type ChannelWebhook struct {
	GuildId      string                `gorm:"primaryKey"`
	ChannelId    string                `gorm:"primaryKey"`
	WebhookId    string                `gorm:"primaryKey"`
	WebhookType  constants.WebhookType `gorm:"primaryKey"`
	WebhookToken string                `gorm:"unique"`
	Language     string                // TODO
}
