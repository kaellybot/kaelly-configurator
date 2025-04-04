package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

type WebhookAlmanax struct {
	WebhookID string    `gorm:"unique;not null"`
	GuildID   string    `gorm:"primaryKey"`
	ChannelID string    `gorm:"primaryKey"`
	Game      amqp.Game `gorm:"primaryKey"`
	Locale    amqp.Language
}

type WebhookFeed struct {
	WebhookID  string    `gorm:"unique;not null"`
	GuildID    string    `gorm:"primaryKey"`
	ChannelID  string    `gorm:"primaryKey"`
	FeedTypeID string    `gorm:"primaryKey"`
	FeedType   FeedType  `gorm:"foreignKey:FeedTypeID"`
	Game       amqp.Game `gorm:"primaryKey"`
	Locale     amqp.Language
}

type WebhookTwitter struct {
	WebhookID      string         `gorm:"unique;not null"`
	GuildID        string         `gorm:"primaryKey"`
	ChannelID      string         `gorm:"primaryKey"`
	TwitterID      string         `gorm:"primaryKey"`
	TwitterAccount TwitterAccount `gorm:"foreignKey:TwitterID"`
	Game           amqp.Game      `gorm:"primaryKey"`
	Locale         amqp.Language
}
