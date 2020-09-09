package main

import (
	"fmt"
	"github.com/xlzd/gotp"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	totpKeyBytes, err := ioutil.ReadFile("/home/igorcavalcante/Private/vpn/vpn-otp")
	if err != nil {
		log.Fatalf("Failed to open vpn otp, err: %s", err.Error())
	}
	totpKey := strings.TrimSpace(string(totpKeyBytes))
	passwordBytes, err := ioutil.ReadFile("/home/igorcavalcante/Private/vpn/vpn-password")
	if err != nil {
		log.Fatalf("Failed to open vpn otp, err: %s", err.Error())
	}
	password := strings.TrimSpace(string(passwordBytes))
	totp := gotp.NewDefaultTOTP(string(totpKey))
	twofactor := totp.Now()
	//fmt.Printf("otp code %s \n", twofactor)
	fmt.Print(password, twofactor)
}
