package time

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	log.Println(GetCurrentDate())
	log.Println(GetCurrentUnix())
	log.Println(GetCurrentMilliUnix())
	log.Println(GetCurrentNanoUnix())
	log.Println(GetCurrentWjDate())
}
