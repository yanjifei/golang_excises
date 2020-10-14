// Create a progress bar for applications that can keep track of a download in
// progress. The progress bar will be on a separate thread and will communicate
// with the main thread using delegates.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type downloader struct {
	totalSize   int
	currentSize int
}

func (d *downloader) start() {
	rand.Seed(88)
	var size int
	for d.currentSize < d.totalSize {
		time.Sleep(800 * time.Millisecond)
		size = rand.Intn(20)
		if d.currentSize+size <= d.totalSize {
			d.currentSize += size
		} else {
			d.currentSize = d.totalSize
		}
		//fmt.Printf("Current Size = %d\r\n", d.currentSize)
	}
	//fmt.Printf("Finished\r\n")
}

func main() {
	d := downloader{totalSize: 200, currentSize: 0}

	go d.start()

	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\r")
		if d.currentSize < d.totalSize {
			fmt.Printf("Progress: %f%%",
				100.0*float32(d.currentSize)/float32(d.totalSize))
		} else {
			fmt.Printf("Progress: %f%%",
				100.0*float32(d.currentSize)/float32(d.totalSize))
			fmt.Println("")
			fmt.Println("End")
			break
		}
	}

}
