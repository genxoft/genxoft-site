package model

import "testing"

func TestFcmSettings(t *testing.T) FcmSettings {
	return FcmSettings{
		Token:    "test-token",
		Timezone: "Asia/Tbilisi",
		IP:       "192.168.1.1",
	}
}
