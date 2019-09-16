package sms_test

import (
	"os"

	"github.com/aiio/qiniu/auth"

	"github.com/aiio/qiniu/sms"
)

var manager *sms.Manager

func init() {
	accessKey := os.Getenv("accessKey")
	secretKey := os.Getenv("secretKey")

	mac := auth.New(accessKey, secretKey)
	manager = sms.NewManager(mac)
}
