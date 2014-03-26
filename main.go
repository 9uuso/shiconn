package main

import(
	"os/exec"
	"net"
	"time"
	"log"
	"runtime"
)

func main() {
	log.Println("Shiconn - shibes checking your connection")
	log.Println("This program shuts down your Windows/Mac/Linux computer in case of losing Internet connection.")
	for {
		_, err := net.Dial("tcp", "google.com:80")
		if err != nil {
			log.Println("Connection lost - shutting down the computer...")
			if runtime.GOOS == "windows" {
				cmd := exec.Command("shutdown", "-s")
			}
			if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
				cmd := exec.Command("sudo", "shutdown")
			}
			err = cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Println(time.Now(), "connection is working!")
			time.Sleep(1 * time.Minute)
		}
	}
}