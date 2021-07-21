package uuid_test

import (
	"gopkg.in/dtapps/go-library.v2/uuid"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	log.Println(uuid.GenUUID())
}
