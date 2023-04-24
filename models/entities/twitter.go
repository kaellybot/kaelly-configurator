package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

type TwitterAccount struct {
	Name   string
	Locale amqp.Language `gorm:"primaryKey"`
}
