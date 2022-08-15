package store

import "genxoft.dev/internal/model"

type FcmRepository interface {
	Find(token string) (*model.FcmSettings, error)
	Create(f *model.FcmSettings) error
	Update(f *model.FcmSettings) error
}
