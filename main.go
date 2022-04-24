package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/lrth06/go-concurrency-example/scrapers"
	"github.com/lrth06/go-concurrency-example/utils"
)

func main() {
	start := time.Now()
	utils.ResetFiles()
	var wg sync.WaitGroup
	wg.Add(6)
	go func() {
		defer wg.Done()
		scrapers.ScrapePhotos()
			fmt.Println("Scraped Photos in: ", time.Since(start))
	}()
	go func() {
		defer wg.Done()
		scrapers.ScrapeUsers()
			fmt.Println("Scraped Users in: ", time.Since(start))
	}()
	go func() {
		defer wg.Done()
		scrapers.ScrapePosts()
			fmt.Println("Scraped Posts in: ", time.Since(start))
	}()
	go func() {
		defer wg.Done()
		scrapers.ScrapeComments()
			fmt.Println("Scraped Comments in: ", time.Since(start))
	}()
	go func() {
		defer wg.Done()
		scrapers.ScrapeAlbums()
			fmt.Println("Scraped Albums in: ", time.Since(start))
	}()
	go func() {
		defer wg.Done()
		scrapers.ScrapeTodos()
			fmt.Println("Scraped Todos in: ", time.Since(start))
	}()
	wg.Wait()
	fmt.Println("Scraped all in: ", time.Since(start))
	os.Exit(0)
}
