package entities

type Guild struct {
	Id       string  `gorm:"primaryKey;type:varchar(100)"`
	ServerId *string `gorm:"type:varchar(100)"`
	Server   *Server `gorm:"foreignKey:ServerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
