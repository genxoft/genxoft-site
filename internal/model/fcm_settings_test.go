package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFcmSettings_BeforeCreate(t *testing.T) {
	f := TestFcmSettings(t)
	f.BeforeCreate()
	assert.GreaterOrEqual(t, f.CreatedAt.Unix(), int64(0))
	assert.GreaterOrEqual(t, f.UpdatedAt.Unix(), int64(0))
}

func TestFcmSettings_BeforeUpdate(t *testing.T) {
	f := TestFcmSettings(t)
	f.BeforeUpdate()
	assert.GreaterOrEqual(t, f.UpdatedAt.Unix(), int64(0))
}

func TestFcmSettings_Validate(t *testing.T) {
	type fields struct {
		Token     string
		Timezone  string
		IP        string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "expect success",
			fields: fields{
				Token:     "test-token",
				Timezone:  "Asia/Tbilisi",
				IP:        "192.168.1.1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: false,
		},
		{
			name: "invalid token",
			fields: fields{
				Token:     "",
				Timezone:  "Asia/Tbilisi",
				IP:        "192.168.1.1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: true,
		},
		{
			name: "invalid timezone",
			fields: fields{
				Token:     "test-token",
				Timezone:  "Arrakis/Arrakeen",
				IP:        "192.168.1.1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FcmSettings{
				Token:     tt.fields.Token,
				Timezone:  tt.fields.Timezone,
				IP:        tt.fields.IP,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if tt.wantErr {
				assert.NotNil(t, f.Validate())
			} else {
				assert.Nil(t, f.Validate())
			}
		})
	}
}
