package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

type TwitterAccount struct {
	Id     string `gorm:"primaryKey"`
	Name   string
	Locale amqp.Language `gorm:"primaryKey"`
}
