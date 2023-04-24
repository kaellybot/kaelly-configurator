package entities

import amqp "github.com/kaellybot/kaelly-amqp"

type WebhookAlmanax struct {
	WebhookId    string
	WebhookToken string
	GuildId      string        `gorm:"primaryKey"`
	ChannelId    string        `gorm:"primaryKey"`
	Locale       amqp.Language `gorm:"primaryKey"`
}

type WebhookFeed struct {
	WebhookId    string
	WebhookToken string
	GuildId      string        `gorm:"primaryKey"`
	ChannelId    string        `gorm:"primaryKey"`
	FeedTypeId   string        `gorm:"primaryKey"`
	Locale       amqp.Language `gorm:"primaryKey"`
	FeedSource   FeedSource    `gorm:"foreignKey:FeedTypeId,Locale"`
}

type WebhookTwitter struct {
	WebhookId      string
	WebhookToken   string
	GuildId        string         `gorm:"primaryKey"`
	ChannelId      string         `gorm:"primaryKey"`
	Locale         amqp.Language  `gorm:"primaryKey"`
	TwitterAccount TwitterAccount `gorm:"foreignKey:Locale"`
}
