package entities

import amqp "github.com/kaellybot/kaelly-amqp"

//nolint:lll,nolintlint // Much clear like that.
type Guild struct {
	ID              string           `gorm:"primaryKey;type:varchar(100)"`
	Game            amqp.Game        `gorm:"primaryKey;type:int"`
	ServerID        *string          `gorm:"type:varchar(100)"`
	Server          *Server          `gorm:"foreignKey:ServerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChannelServers  []ChannelServer  `gorm:"foreignKey:GuildID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AlmanaxWebhooks []WebhookAlmanax `gorm:"foreignKey:GuildID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FeedWebhooks    []WebhookFeed    `gorm:"foreignKey:GuildID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TwitterWebhooks []WebhookTwitter `gorm:"foreignKey:GuildID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
