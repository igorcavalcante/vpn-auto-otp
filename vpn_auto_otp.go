package main

import (
	"fmt"
	"github.com/xlzd/gotp"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func main2() {
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
	fmt.Printf("otp code %s \n", twofactor)

	fmt.Println("Stop existing vpnc process")
	err = exec.Command("sudo", "vpnc-disconnect").Run()
	if exitError, ok := err.(*exec.ExitError); ok {
		if exitError.ExitCode() != 1 {
			log.Fatalf("Failed to stop vpnc, err: %s", exitError.Error())
		}
	} else if err != nil {
		log.Fatalf("Failed to stop vpnc, err: %s", err.Error())
	}
	fmt.Println("Start vpnc process")
	err = exec.Command("sudo", "vpnc", "/home/igorcavalcante/Private/vpn/globo.com.conf", "--password", password+twofactor).Run()
	if err != nil {
		log.Fatalf("Failed to start vpnc, err: %s", err.Error())
	}
	fmt.Println("Setup DNS")
	err = exec.Command("cp", "-Rf", "/etc/resolv.conf.vpn", "/etc/resolv.conf").Run()
	if err != nil {
		log.Fatalf("Failed to setup resolv.conf, err: %s", err.Error())
	}
	fmt.Println("VPN setup finished")
}
