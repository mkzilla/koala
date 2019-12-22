package config

import (
	"fmt"
	"time"
)

type Token struct {
	ISS string
	IAT int64
	EXP int64
	AUD int64
}

func (t Token) Valid() error {
	if t.ISS != "koala" {
		return fmt.Errorf("not koala token")
	}
	if time.Now().Unix() < t.IAT {
		return fmt.Errorf("the token has been tampered with")
	}
	if t.EXP < time.Now().Unix() {
		return fmt.Errorf("the token has expired")
	}
	return nil
}
