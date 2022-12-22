package entities

type Guild struct {
	Id     string `gorm:"primaryKey"`
	Server string `gorm:"unique"`
}
