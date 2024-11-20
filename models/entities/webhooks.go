package entities

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
)

type WebhookAlmanax struct {
	WebhookID    string `gorm:"unique;not null"`
	WebhookToken string
	GuildID      string    `gorm:"primaryKey"`
	ChannelID    string    `gorm:"primaryKey"`
	Game         amqp.Game `gorm:"primaryKey"`
	Locale       amqp.Language
	RetryNumber  int64 `gorm:"default:0"`
	PublishedAt  *time.Time
	FailedAt     *time.Time
}

type WebhookFeed struct {
	WebhookID    string `gorm:"unique;not null"`
	WebhookToken string
	GuildID      string    `gorm:"primaryKey"`
	ChannelID    string    `gorm:"primaryKey"`
	FeedTypeID   string    `gorm:"primaryKey"`
	FeedType     FeedType  `gorm:"foreignKey:FeedTypeID"`
	Game         amqp.Game `gorm:"primaryKey"`
	Locale       amqp.Language
	RetryNumber  int64 `gorm:"default:0"`
	PublishedAt  *time.Time
	FailedAt     *time.Time
}

type WebhookTwitch struct {
	WebhookID    string `gorm:"unique;not null"`
	WebhookToken string
	GuildID      string   `gorm:"primaryKey"`
	ChannelID    string   `gorm:"primaryKey"`
	StreamerID   string   `gorm:"primaryKey"`
	Streamer     Streamer `gorm:"foreignKey:StreamerID"`
	Locale       amqp.Language
	RetryNumber  int64 `gorm:"default:0"`
	PublishedAt  *time.Time
	FailedAt     *time.Time
}

type WebhookTwitter struct {
	WebhookID      string `gorm:"unique;not null"`
	WebhookToken   string
	GuildID        string         `gorm:"primaryKey"`
	ChannelID      string         `gorm:"primaryKey"`
	TwitterID      string         `gorm:"primaryKey"`
	TwitterAccount TwitterAccount `gorm:"foreignKey:TwitterID"`
	Locale         amqp.Language
	RetryNumber    int64 `gorm:"default:0"`
	PublishedAt    *time.Time
	FailedAt       *time.Time
}

type WebhookYoutube struct {
	WebhookID    string `gorm:"unique;not null"`
	WebhookToken string
	GuildID      string  `gorm:"primaryKey"`
	ChannelID    string  `gorm:"primaryKey"`
	VideastID    string  `gorm:"primaryKey"`
	Videast      Videast `gorm:"foreignKey:VideastID"`
	Locale       amqp.Language
	RetryNumber  int64 `gorm:"default:0"`
	PublishedAt  *time.Time
	FailedAt     *time.Time
}
