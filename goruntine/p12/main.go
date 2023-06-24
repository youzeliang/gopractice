package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// goroutine的线程池
type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	response    string
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) (int, string) {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}

	url := "https://www.cqyti.com/sy.htm" // 替换为你要请求的网址
	resp, err := http.Get(url)
	if err != nil {
		return sum, fmt.Sprintf("Error: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sum, fmt.Sprintf("Error: %s", err.Error())
	}

	return sum, string(body)
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		sumOfDigits, response := digits(job.randomno)
		output := Result{
			job:         job,
			response:    response,
			sumofdigits: sumOfDigits,
		}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
		fmt.Printf("Response: %s\n", result.response)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 10000000
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10000
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Total time taken:", diff.Seconds(), "seconds")
}
