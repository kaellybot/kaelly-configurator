package entities

import amqp "github.com/kaellybot/kaelly-amqp"

type ChannelServer struct {
	GuildID   string    `gorm:"primaryKey;type:varchar(100);"`
	ChannelID string    `gorm:"primaryKey;type:varchar(100);"`
	Game      amqp.Game `gorm:"primaryKey;type:int"`
	ServerID  string    `gorm:"type:varchar(100);"`
	Guild     Guild     `gorm:"foreignKey:GuildID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Server    Server    `gorm:"foreignKey:ServerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
