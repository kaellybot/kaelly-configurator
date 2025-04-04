package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

type TwitterAccount struct {
	ID     string `gorm:"primaryKey"`
	Name   string `gorm:"type:varchar(250)"`
	Locale amqp.Language
	Game   amqp.Game
}
