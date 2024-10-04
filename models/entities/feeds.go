package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

type FeedSource struct {
	FeedTypeID string        `gorm:"primaryKey"`
	Locale     amqp.Language `gorm:"primaryKey"`
}
