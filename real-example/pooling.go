package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Result struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}
type Job struct {
	number int
}

var jobs = make(chan Job, 100)
var results = make(chan Result, 100)
var resultCollection []Result

const URL = "https://xkcd.com"

func fetch(n int) (*Result, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	url := strings.Join([]string{URL, fmt.Sprintf("%d", n), "info.0.json"}, "/")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http err: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status not ok")
	}

	data := &Result{}
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		return nil, fmt.Errorf("decode err: %v", err)
	}

	res.Body.Close()
	return data, nil
}
func allocateJob(jobsNo int) {
	for i := 0; i < jobsNo; i++ {
		fmt.Printf("job %v allocated \n", i+1)
		jobs <- Job{i + 1}
	}
	close(jobs)
}
func worker(wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	for job := range jobs {
		result, err := fetch(job.number)
		if err != nil {
			log.Printf("fetch error: %v", err)
		}
		fmt.Printf("job %d fetched by worker %d \n", job.number, workerID)
		results <- *result
	}
}
func createWorkerPool(workersNo int) {
	var wg sync.WaitGroup
	for i := 0; i < workersNo; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
	close(results)
}
func getResults(done chan bool) {
	fmt.Println("results channel is ready to get")
	for result := range results {
		resultCollection = append(resultCollection, result)
	}
	done <- true
}
func writeInFile(data []byte) error {
	f, err := os.Create("output.json")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
func main() {
	//allocate job
	jobsNo := 20
	go allocateJob(jobsNo)

	//get results
	done := make(chan bool)
	go getResults(done)

	//create worker pool
	workersNo := 10
	createWorkerPool(workersNo)

	//wait to get all results
	<-done

	//write result to file
	res, err := json.Marshal(resultCollection)
	if err != nil {
		log.Printf("json err: %v", err)
	}

	err = writeInFile(res)
	if err != nil {
		log.Printf("write err: %v", err)
	}
}
