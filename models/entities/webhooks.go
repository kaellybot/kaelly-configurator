package entities

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
)

type WebhookAlmanax struct {
	WebhookID    string
	WebhookToken string
	GuildID      string        `gorm:"primaryKey"`
	ChannelID    string        `gorm:"primaryKey"`
	Locale       amqp.Language `gorm:"primaryKey"`
	Game         amqp.Game     `gorm:"primaryKey"`
	RetryNumber  int64         `gorm:"default:0"`
	UpdatedAt    time.Time
}

type WebhookFeed struct {
	WebhookID    string
	WebhookToken string
	GuildID      string        `gorm:"primaryKey"`
	ChannelID    string        `gorm:"primaryKey"`
	FeedTypeID   string        `gorm:"primaryKey"`
	Locale       amqp.Language `gorm:"primaryKey"`
	Game         amqp.Game     `gorm:"primaryKey"`
	FeedSource   FeedSource    `gorm:"foreignKey:FeedTypeID,Locale,Game"`
	RetryNumber  int64         `gorm:"default:0"`
	UpdatedAt    time.Time
}

type WebhookTwitch struct {
	WebhookID    string
	WebhookToken string
	GuildID      string   `gorm:"primaryKey"`
	ChannelID    string   `gorm:"primaryKey"`
	StreamerID   string   `gorm:"primaryKey"`
	Streamer     Streamer `gorm:"foreignKey:StreamerID"`
	RetryNumber  int64    `gorm:"default:0"`
	UpdatedAt    time.Time
}

type WebhookTwitter struct {
	WebhookID      string
	WebhookToken   string
	GuildID        string         `gorm:"primaryKey"`
	ChannelID      string         `gorm:"primaryKey"`
	TwitterID      string         `gorm:"primaryKey"`
	TwitterAccount TwitterAccount `gorm:"foreignKey:TwitterID"`
	RetryNumber    int64          `gorm:"default:0"`
	UpdatedAt      time.Time
}

type WebhookYoutube struct {
	WebhookID    string
	WebhookToken string
	GuildID      string  `gorm:"primaryKey"`
	ChannelID    string  `gorm:"primaryKey"`
	VideastID    string  `gorm:"primaryKey"`
	Videast      Videast `gorm:"foreignKey:VideastID"`
	RetryNumber  int64   `gorm:"default:0"`
	UpdatedAt    time.Time
}
