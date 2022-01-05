package scripts

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func Generate() {
	key := make([]byte, 64)

	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
	}

	stringB64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(stringB64)
}
