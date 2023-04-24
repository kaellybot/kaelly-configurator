package entities

type Guild struct {
	Id              string           `gorm:"primaryKey;type:varchar(100)"`
	ServerId        *string          `gorm:"type:varchar(100)"`
	Server          *Server          `gorm:"foreignKey:ServerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChannelServers  []ChannelServer  `gorm:"foreignKey:GuildId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AlmanaxWebhooks []AlmanaxWebhook `gorm:"foreignKey:GuildId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RssWebhooks     []RssWebhook     `gorm:"foreignKey:GuildId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TwitterWebhooks []TwitterWebhook `gorm:"foreignKey:GuildId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
