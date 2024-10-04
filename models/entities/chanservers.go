package entities

type ChannelServer struct {
	GuildID   string `gorm:"primaryKey;type:varchar(100);"`
	ChannelID string `gorm:"primaryKey;type:varchar(100);"`
	ServerID  string `gorm:"type:varchar(100);"`
	Guild     Guild  `gorm:"foreignKey:GuildID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Server    Server `gorm:"foreignKey:ServerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
