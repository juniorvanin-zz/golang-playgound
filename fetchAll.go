package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		os.Exit(1)
	}

	ch := make(chan string)

	for _, url :=  range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		_, err = f.WriteString(<-ch + "\n")

		if err != nil {
			fmt.Println("error writing output.txt file")
		}
	}

}

func fetch(url string,ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}