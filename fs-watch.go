package gowatch

import (
	"log"
	"time"

	"github.com/beefsack/go-rate"
	"golang.org/x/exp/fsnotify"
)

// Watch will watch a path for changes and run the set of commands against them when changes happen.
func Watch(path string) {
	ch := make(chan int, 1)
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

	// Block indefinitely.
	select {}
}

func startWatching(path string, ch chan int) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	rl := rate.New(1, time.Second/4)
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if ev.IsCreate() {
					rateLimit(ch, rl)
					debug("event:", ev)
				}
			case err := <-watcher.Error:
				debug("error:", err)
			}
		}
	}()

	// Run commands on first run.
	header("")
	ch <- 1

	// Run commands every time a file changes.
	err = watcher.Watch(path)
	if err != nil {
		log.Fatal(err)
	}
}

func rateLimit(ch chan int, rl *rate.RateLimiter) {
	if ok, remaining := rl.Try(); ok {
		ch <- 2
	} else {
		debug("Spam filter triggered, please wait %s\n", remaining)
	}
}
