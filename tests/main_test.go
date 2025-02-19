package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/YulyaY/go_final_project.git/internal/config"
	"github.com/golang-jwt/jwt"
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	setup()
	defer teardown()

	return m.Run()
}

func setup() {
	appConfig, err := config.LoadAppConfig()
	if err != nil {
		panic(fmt.Sprintf("load config error: %v", err))
	}

	Token = genToken(appConfig.Secret)
}

func teardown() {
}

func genToken(secret string) string {
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		panic(fmt.Sprintf("gen token error: %v", err))
	}

	return token
}
