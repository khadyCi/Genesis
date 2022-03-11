package main

//"github.com/chromedp/chromedp/kb"
import (
	"flag"
	"log"

	//"gitlab.i/vacolba/Genesis/config"
	"github.com/khadyCi/Genesis/config"
)

var (
	ruta *string
	cfg  config.Config
)

func setFlags() {
	ruta = flag.String("c", "../config.yaml", "")
}

func main() {
	err := config.LoadConfig(&cfg, *ruta)
	if err != nil {
		log.Fatal(err)
	}
}
