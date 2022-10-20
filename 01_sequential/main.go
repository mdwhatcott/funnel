package main

import (
	"fmt"
	"time"

	"github.com/mdwhatcott/funnel"
)

func main() {
	started := time.Now()
	defer func() { fmt.Println(time.Since(started)) }()
	
	for _, address := range funnel.Addresses {
		fmt.Println(funnel.ScrapeTitle(address))
	}
}