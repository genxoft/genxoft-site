package sqlitestore_test

import (
	"genxoft.dev/internal/model"
	"genxoft.dev/internal/store/sqlitestore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFcmRepository_Create(t *testing.T) {
	s, teardown := sqlitestore.TestStore(t)
	defer teardown("fcm_settings")

	f := model.TestFcmSettings(t)

	assert.Nil(t, s.FcmToken().Create(&f))
	assert.NotNil(t, f)
}

func TestFcmRepository_Find(t *testing.T) {
	s, teardown := sqlitestore.TestStore(t)
	defer teardown("fcm_settings")

	f := model.TestFcmSettings(t)

	assert.Nil(t, s.FcmToken().Create(&f))
	assert.NotNil(t, f)

	ft, err := s.FcmToken().Find(model.TestFcmSettings(t).Token)

	assert.Nil(t, err)
	assert.NotNil(t, ft)
}

func TestFcmRepository_Update(t *testing.T) {
	s, teardown := sqlitestore.TestStore(t)
	defer teardown("fcm_settings")

	f := model.TestFcmSettings(t)

	assert.Nil(t, s.FcmToken().Create(&f))
	assert.NotNil(t, f)

	f.Timezone = "Asia/Jakarta"

	assert.Nil(t, s.FcmToken().Update(&f))
	assert.NotNil(t, f)

	ft, err := s.FcmToken().Find(model.TestFcmSettings(t).Token)

	assert.Nil(t, err)
	assert.NotNil(t, ft)

	assert.Equal(t, "Asia/Jakarta", ft.Timezone)
}
