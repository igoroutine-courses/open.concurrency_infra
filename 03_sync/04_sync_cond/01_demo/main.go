package main

import (
	"log"
	"sync"
	"time"
)

func subscribe(name string, data map[string]string, c *sync.Cond) {
	c.L.Lock()
	defer c.L.Unlock()

	// not spurious wakeup!
	for len(data) == 0 {
		c.Wait()
	}

	log.Printf("[%s] %s\n", name, data["key"])
}

func publish(name string, data map[string]string, c *sync.Cond) {
	time.Sleep(time.Second)

	c.L.Lock()
	data["key"] = "value"
	c.L.Unlock()

	log.Printf("[%s] data publisher\n", name)
	c.Broadcast()
}

func main() {
	data := map[string]string{}
	cond := sync.NewCond(&sync.Mutex{})

	wg := new(sync.WaitGroup)
	wg.Go(func() {
		subscribe("subscriber_1", data, cond)
	})

	wg.Go(func() {
		subscribe("subscriber_2", data, cond)
	})

	wg.Go(func() {
		publish("publisher", data, cond)
	})

	wg.Wait()
}
