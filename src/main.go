package main

import (
	"log"
	"os"
	"sync"
)

func main() {
	URL := os.Getenv("HTTP_NODE")
	PRV_KEY := os.Getenv("PRV_KEY")
	DEST_ADDRESS := os.Getenv("DEST_ADDRESS")

	mainEx, err := NewExecutor(
		URL,
		PRV_KEY,
	)
	if err != nil {
		log.Fatalln(err)
	}
	claimer := &Claimer{
		Executor: *mainEx,
	}

	err = claimer.buildDistributor()
	if err != nil {
		log.Fatalln(err)
	}
	err = claimer.buildToken()
	if err != nil {
		log.Fatalln(err)
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for {
			tx, err := claimer.claim()
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println(tx)
			break
		}
		wg.Done()
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for {
			tx, err := claimer.withdrawTokens(DEST_ADDRESS, 625.0)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println(tx)
			break
		}
		wg.Done()
	}(wg)

	wg.Wait()
}
