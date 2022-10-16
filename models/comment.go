package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `gorm:"not null;" json:"message" form:"message" valid:"required~Message is required"`
	UserId  int    `json:"user_id" form:"user_id"`
	PhotoId int    `json:"photo_id" form:"photo_id"`
	Photo   *Photo
	User    *User
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
