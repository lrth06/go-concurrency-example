package main

import (
	"fmt"
	"time"

	"github.com/lrth06/go-concurrency-example/scrapers"
	"github.com/lrth06/go-concurrency-example/utils"
)





func main() {
	start := time.Now()
	utils.ResetFiles()
	
	scrapers.ScrapePhotos()
	photoTime := time.Since(start)
	fmt.Println("Scraped Photos in: ", photoTime)

	scrapers.ScrapeUsers()
	userTime := time.Since(start) - photoTime
	fmt.Println("Scraped Users in: ", userTime)
	scrapers.ScrapePosts()
	postTime := time.Since(start) - photoTime - userTime

	fmt.Println("Scraped Posts in: ", postTime)
	scrapers.ScrapeComments()
	commentTime := time.Since(start) - photoTime - userTime - postTime
	fmt.Println("Scraped Comments in: ", commentTime)

	fmt.Println("Total Time taken:", time.Since(start))
}
