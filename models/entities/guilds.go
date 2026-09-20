package entities

import amqp "github.com/kaellybot/kaelly-amqp"

// The webhook associations carry no constraint: production has no foreign key from
// the webhook tables to guilds, so declaring one here would misdescribe the schema.
//
//nolint:lll,nolintlint // Much clear like that.
type Guild struct {
	ID              string           `gorm:"primaryKey;type:varchar(100)"`
	Game            amqp.Game        `gorm:"primaryKey;type:int"`
	ServerID        *string          `gorm:"type:varchar(100)"`
	Server          *Server          `gorm:"foreignKey:ServerID,Game;references:ID,Game;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ChannelServers  []ChannelServer  `gorm:"foreignKey:GuildID,Game;references:ID,Game;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AlmanaxWebhooks []WebhookAlmanax `gorm:"foreignKey:GuildID,Game;references:ID,Game"`
	FeedWebhooks    []WebhookFeed    `gorm:"foreignKey:GuildID,Game;references:ID,Game"`
	TwitterWebhooks []WebhookTwitter `gorm:"foreignKey:GuildID,Game;references:ID,Game"`
}
