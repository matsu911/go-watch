package gowatch

import (
	"log"

	"golang.org/x/exp/fsnotify"
)

// Watch will watch a path for changes and run the set of commands against them when changes happen.
func Watch(path string) {
	ch := make(chan bool, 1)
	startWatching(path, ch)

	go func() {
		for {
			select {
			case event := <-ch:
				debug("Event happened! %v", event)
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
					debug("event:", ev)
				}
			case err := <-watcher.Error:
				debug("error:", err)
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
