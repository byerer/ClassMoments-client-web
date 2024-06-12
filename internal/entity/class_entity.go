package entity

type Class struct {
	ClassID          uint `gorm:"primary_key;auto_increment"`
	ClassName        string
	SchoolName       string
	ClassDescription string
}
