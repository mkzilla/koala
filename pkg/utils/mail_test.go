package utils

import (
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	err := SendMail("0.0.0.0:10025", "no-reply@mkzilla.com", "792400644@qq.com", "欢迎注册 mkzilla koala", registerTemplate)
	if err != nil {
		t.Fatal(err)
	}
}
