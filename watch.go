package gowatch

import (
	"fmt"
	"log"
	"time"

	"github.com/beefsack/go-rate"
	"github.com/fatih/color"
	"golang.org/x/exp/fsnotify"
)

// Watch will watch a path for changes and run the set of commands against them when changes happen.
func Watch(path string) {
	ch := make(chan bool, 1)
	startWatching(path, ch)

	// Counter to run in between each session.
	counter := intSeq()

	go func() {
		for {
			select {
			case event := <-ch:
				debug("Event happened! %v", event)
				run(counter())
			}
		}
	}()

	// Block indefinitely.
	select {}
}

func startWatching(path string, ch chan bool) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		rl := rate.New(1, time.Second/4)
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
	fmt.Println(header("1", color.FgMagenta))
	ch <- true

	// Run commands every time a file changes.
	err = watcher.Watch(path)
	if err != nil {
		log.Fatal(err)
	}
}

// rateLimit is a leaky bucket which allows one execution per quarter second and does not queue.
func rateLimit(ch chan bool, rl *rate.RateLimiter) {
	if ok, remaining := rl.Try(); ok {
		ch <- true
	} else {
		debug("Rate limited for another %s seconds.\n", remaining)
	}
}

func intSeq() func() int {
	i := 1
	return func() int {
		i++
		return i
	}
}
