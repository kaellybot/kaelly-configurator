package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

type TwitterAccount struct {
	ID     string `gorm:"primaryKey"`
	Name   string
	Locale amqp.Language
	Game   amqp.Game
}
