package repository

import (
	"github.com/boni/golang-mongo/src/module/profile/model"
)

type ProfileRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindByID(string) (*model.Profile, error)
	FindAll() (model.Profiles, error)
}
