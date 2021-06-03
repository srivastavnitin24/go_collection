package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id        int
	randomNum int
}
type Result struct {
	job         Job
	sumOfDigits int
}

var jobchan = make(chan Job, 10)
var resultchan = make(chan Result, 10)

func dosumdigits(num int) int {
	sum := 0
	n := num
	for n != 0 {
		x := n % 10
		sum += x
		n = n / 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobchan {
		output := Result{job, dosumdigits(job.randomNum)}
		resultchan <- output
	}
	wg.Done()
}

func crateWorkerPools(noofworkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noofworkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(resultchan)
}

func allocate(noofjobs int) {
	for i := 0; i < noofjobs; i++ {
		num := rand.Intn(999)
		job := Job{
			id:        i,
			randomNum: num,
		}
		jobchan <- job
	}
	close(jobchan)
}

func readresults(done chan bool) {
	for res := range resultchan {
		fmt.Printf("\nJob id %v , randomint = %v , sum of digits = %v ", res.job.id+1, res.job.randomNum, res.sumOfDigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 100

	go allocate(noOfJobs)

	done := make(chan bool)
	go readresults(done)

	noofworkers := 10
	crateWorkerPools(noofworkers)
	<-done

	endtime := time.Now()
	diff := startTime.Sub(endtime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")

}
