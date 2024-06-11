package entity

type Media struct {
	ID        uint `gorm:"primary_key;auto_increment"`
	MomentID  uint
	MediaType string
	URL       string
	CreatedAt string
}
