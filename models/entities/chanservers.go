package entities

type ChannelServer struct {
	GuildId   string `gorm:"primaryKey;type:varchar(100);"`
	ChannelId string `gorm:"primaryKey;type:varchar(100);"`
	ServerId  string `gorm:"type:varchar(100);"`
	Guild     Guild  `gorm:"foreignKey:GuildId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Server    Server `gorm:"foreignKey:ServerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
