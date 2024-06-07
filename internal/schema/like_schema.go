package schema

import "gorm.io/gorm"

type LikeReq struct {
	gorm.Model
	liker    uint
	likee    uint
	momentID uint
}
