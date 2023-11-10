package utils

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"math"
)

func CreateSharelink() string {

	buff := make([]byte, int(math.Ceil(float64(32)/2)))
	_, err := rand.Read(buff)
	if err != nil {
		log.Println("Error creating random string for sharelink: ", err.Error())
	}
	str := hex.EncodeToString(buff)

	return str[:32]
}
