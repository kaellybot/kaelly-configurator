package entities

import amqp "github.com/kaellybot/kaelly-amqp"

type AlmanaxWebhook struct {
	GuildId      string `gorm:"primaryKey"`
	ChannelId    string `gorm:"primaryKey"`
	WebhookId    string
	WebhookToken string
	Language     amqp.Language `gorm:"primaryKey"`
}

type RssWebhook struct {
	GuildId      string `gorm:"primaryKey"`
	ChannelId    string `gorm:"primaryKey"`
	FeedId       string `gorm:"primaryKey"` // TODO feed Type
	WebhookId    string
	WebhookToken string
	Language     amqp.Language `gorm:"primaryKey"`
}

type TwitterWebhook struct {
	GuildId      string `gorm:"primaryKey"`
	ChannelId    string `gorm:"primaryKey"`
	WebhookId    string
	WebhookToken string
	Language     amqp.Language `gorm:"primaryKey"`
}
