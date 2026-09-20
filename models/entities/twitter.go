package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

// TwitterAccount is identified by (id, game): the same account may be followed in
// several games, with its own locale per game.
type TwitterAccount struct {
	ID     string    `gorm:"primaryKey"`
	Game   amqp.Game `gorm:"primaryKey"`
	Name   string    `gorm:"type:varchar(250)"`
	Locale amqp.Language
}
