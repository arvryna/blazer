package main

import (
	"fmt"
	"time"
)

type Job struct {
	Id              int
	Message         string
	TotalCount      int
	RetryVerdict    bool
	DownloadSuccess bool
}

var successTracker = make(chan Job)
var retryQueue = make(chan Job)
var verictQueue = make(chan Job)

func main() {
	start := time.Now()
	initf()
	fmt.Println("Download finished in: ", time.Since(start))
}

func initf() {
	go successQ()
	go retryQ()
	go retryQ()
	go retryQ()
	go retryQ()
	go retryQ()
	go retryQ()

	// Dispatch jobs
	go jobScheduler(450)

	for job := range verictQueue {
		if job.RetryVerdict == false {
			fmt.Println("download incomplete, closing")
			return
		} else if job.DownloadSuccess {
			fmt.Printf("download done")
			return
		}
	}
}

func jobScheduler(count int) {
	for i := 1; i <= count; i++ {
		job := Job{i, "", count, true, false}
		go downloadFile(job)
	}
}

func downloadFile(job Job) {
	time.Sleep(300)
	if job.Id%3 == 0 {
		fmt.Println("download failed retrying: ", job.Id)
		retryQueue <- job
	} else {
		successTracker <- job
	}
}

func successQ() {
	downloadCount := 0
	for job := range successTracker {
		downloadCount++
		fmt.Println("Download finished ", job.Id)
		if downloadCount == job.TotalCount {
			job.DownloadSuccess = true
			verictQueue <- job
		}
	}
}

func retryQ() {
	for job := range retryQueue {
		time.Sleep(1 * time.Second)
		job.Id = job.Id + 1
		downloadFile(job)
	}
}
