package model

import (
	"errors"
	"time"
)

type FcmSettings struct {
	Token     string    `json:"token"`
	Timezone  string    `json:"timezone"`
	IP        string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (f *FcmSettings) Validate() error {
	if len(f.Token) < 1 || len(f.Token) > 255 {
		return errors.New("token length invalid")
	}
	_, err := time.LoadLocation(f.Timezone)
	if err != nil {
		return errors.New("invalid timezone")
	}

	return nil
}

func (f *FcmSettings) BeforeCreate() {
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
}

func (f *FcmSettings) BeforeUpdate() {
	f.UpdatedAt = time.Now()
}
