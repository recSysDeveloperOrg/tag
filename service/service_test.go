package service

import (
	"math/rand"
	"tag/config"
	"tag/model"
	"testing"
	"time"
)

const (
	testUser  = "800000000000000000000002"
	testMovie = "100000000000000000000005"
)

func TestMain(m *testing.M) {
	if err := config.InitConfig(config.CfgFileNested); err != nil {
		panic(err)
	}
	if err := model.InitModel(); err != nil {
		panic(err)
	}
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func generateStr() string {
	rand.Seed(time.Now().Unix())
	s := ""
	for i := 0; i < 23; i++ {
		s += string(rune('0' + rand.Intn(10)))
	}

	return "9" + s
}
