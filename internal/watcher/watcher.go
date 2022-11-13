package watcher

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
)

type OperationType int

const (
	Create OperationType = iota
	Remove
	Modify
)

type NotifyMessage struct {
	Operation OperationType
	FilePath  string
}

func Watch(directory string, notifyCh chan<- NotifyMessage) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			log.Printf("goroutine file watch alive.")
			select {
			case event, ok := <-watcher.Events:
				log.Printf("receive event: %+v", event)
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					notifyCh <- NotifyMessage{
						Operation: Modify,
						FilePath:  event.Name,
					}
				} else if event.Op&fsnotify.Create == fsnotify.Create {
					notifyCh <- NotifyMessage{
						Operation: Create,
						FilePath:  event.Name,
					}
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					err = watcher.Add(directory)
					if err != nil {
						log.Fatal(err)
					}
					notifyCh <- NotifyMessage{
						Operation: Remove,
						FilePath:  event.Name,
					}
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(directory)
	if err != nil {
		log.Fatal(err)
	}
	<-done
	fmt.Println("watch done")
}
