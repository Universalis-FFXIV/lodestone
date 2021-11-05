package main

import (
	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone/v2"
)

func main() {
	_ = godestone.NewScraper(bingode.New(), godestone.EN)
}
