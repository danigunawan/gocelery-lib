package main

import (
	"fmt"
	"github.com/danigunawan/gocelery-lib"
)

func get_value_from_multiple_return(val ...interface{}) []interface{} {
	return val
}

// Run Celery Worker First!
// celery -A worker worker --loglevel=debug --without-heartbeat --without-mingle
func main() {

	// // create RabbitMQ connection
	url := "amqp://guest:guest@localhost:5672/"

	// // initialize celery client and configure 1 workers
	celery_client, _ := gocelery_lib.NewCeleryClient(
		gocelery_lib.NewAMQPCeleryBroker(url),
		gocelery_lib.NewAMQPCeleryBackend(url),
		1, // number of workers
	)

	// prepare arguments
	taskName := "worker_py.add"
	argA := 2
	argB := 2

	// run task
	asyncResult, err := celery_client.Delay(taskName, argA, argB)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(asyncResult)
	// fmt.Println(celery_client)

	// TaskID := asyncResult.TaskID
	// fmt.Println(get_value_from_multiple_return(asyncResult.Ready())[0])

	// check if result is ready
	// isReady, _ := asyncResult.Ready()
	// fmt.Printf("Ready status: %v\n", isReady)

	// AsyncResultReady := asyncResult.Ready()
	// if AsyncResultReady != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(asyncResult.Ready()

	// fmt.Println(asyncResult.TaskID)
	// fmt.Println(asyncResult.GetResult("a70e8c97-a109-40f3-8acb-49ab749e58d3"))
	// fmt.Println(celery_client
	// get results from backend with timeout

	// res, err := AsyncResult.backend.GetResult(AsyncResult.TaskID)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res)
	// queueName := strings.Replace("efbaf301-2387-44b3-beb6-b5033cc12cfa", "-", "", -1)

	// get results from backend with timeout

	// run task
	// asyncResultID, err := gocelery_lib.Delay("09afe5c6-5005-4250-829c-93bf3856b93d")
	// if err != nil {
	// fmt.Println(err)
	// }

	// fmt.Println("HASIL BY ID")
	// out := gocelery_lib.AMQPCeleryBackend.GetResult("09afe5c6-5005-4250-829c-93bf3856b93d")
	// fmt.Println(out)
	// fmt.Println(asyncResultID)

	// fmt.Println("HASIL BY TIME")

	// res, err := asyncResult.Get(10 * time.Second)
	// if err != nil {
	// 	panic(err)
	// }

	// log.Printf("result: %+v of type %+v", res, reflect.TypeOf(res))

	// testCases := []struct {
	// 	name    string
	// 	backend *CeleryBackend
	// }{
	// 	{
	// 		name: "get result from redis backend",
	// 	},
	// }
	// for _, tc := range testCases {
	// 	taskID := uuid.Must(uuid.NewV4()).String()
	// 	// value must be float64 for testing due to json limitation
	// 	value := reflect.ValueOf(rand.Float64())
	// 	resultMessage := getReflectionResultMessage(&value)
	// 	messageBytes, err := json.Marshal(resultMessage)
	// 	if err != nil {

	// 	}
	// 	conn := tc.backend.Get()
	// 	defer conn.Close()
	// 	_, err = conn.Do("SETEX", fmt.Sprintf("celery-task-meta-%s", taskID), 86400, messageBytes)
	// 	if err != nil {

	// 	}
	// 	res, err := tc.backend.GetResult(taskID)
	// 	if err != nil {

	// 	}

}
