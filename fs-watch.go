package gowatch

import (
	"log"

	"golang.org/x/exp/fsnotify"
)

func Watch(path string) {
	ch := make(chan bool, 1)
	startWatching(path, ch)

	go func() {
		for {
			select {
			case event := <-ch:
				log.Printf("Event happened! %v", event)
				run()
			}
		}
	}()

	select {}
}

func startWatching(path string, ch chan bool) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if ev.IsCreate() {
					rateLimit(ch)
					log.Println("event:", ev)
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch(path)
	if err != nil {
		log.Fatal(err)
	}
}

func rateLimit(ch chan bool) {
	select {
	case ch <- true:
	default:
		return
	}

}
