package main

import (
	"github.com/danigunawan/gocelery-lib"
	"time"
)

// Celery Task
func add(a int, b int) int {
	return a + b
}

func main() {

	// create AMQP & Redis connection pool
	url_amqp := "amqp://guest:guest@localhost:5672/"
	// url_redis :=

	// create broker and backend
	// celeryBroker := gocelery_lib.NewRedisCeleryBroker("redis://localhost:6379")
	// celeryBackend := gocelery_lib.NewRedisCeleryBackend("redis://localhost:6379")

	//use AMQP instead
	celeryBroker := gocelery_lib.NewAMQPCeleryBroker(url_amqp)
	celeryBackend := gocelery_lib.NewAMQPCeleryBackend(url_amqp)

	// initialize celery client
	cli, _ := gocelery_lib.NewCeleryClient(celeryBroker, celeryBackend, 5) // number of workers

	// register task
	cli.Register("worker_go.add", add)

	// start workers (non-blocking call)
	cli.StartWorker()

	// wait for client request
	time.Sleep(100000 * time.Second)

	// stop workers gracefully (blocking call)
	cli.StopWorker()
}
