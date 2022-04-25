package scrapers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/lrth06/go-concurrency-example/types"
	"github.com/lrth06/go-concurrency-example/utils"
)

func ScrapePhotos(){
	http.DefaultClient.Timeout = 10 * time.Second

	imgResponse, err := utils.GetResponse("https://jsonplaceholder.typicode.com/photos/")
	if err != nil {
		panic(err)
	}
	// unbuffered := make(chan types.ImgData)
	header := "id,title,url"
	utils.WriteCSV("output/photos.csv", []string{header})
	var data []types.ImgData
	json.Unmarshal([]byte(imgResponse), &data)
	dataLength := len(data)
	var resData []types.ImgData
	var wg sync.WaitGroup
	wg.Add(dataLength)
	for i := 0; i < dataLength; i++ {
		line := fmt.Sprintf("%d,%s,%s", data[i].ID, data[i].Title, data[i].URL)
		utils.WriteCSV("output/photos.csv", []string{line})
		go func(i int) {
			defer wg.Done()
			resData = append(resData, data[i])
			res ,err:= http.Get(data[i].URL)
				if err != nil {
					panic(err)
				}
			file, err := os.Create(fmt.Sprintf("output/photos/%d.jpg", data[i].ID))
			if err != nil {
				panic(err)
			}
			defer file.Close()
			image , err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			_, err = file.Write(image)
			if err != nil {
				panic(err)
			}
			}(i)
		}
	wg.Wait()
}

func ScrapeUsers(){
	userResponse, err := utils.GetResponse("https://jsonplaceholder.typicode.com/users/")
	if err != nil {
		panic(err)
	}
	var userData []types.UserData
	json.Unmarshal([]byte(userResponse), &userData)
	header := "id,name,username,email,address,phone,website,company"
	utils.WriteCSV("output/users.csv", []string{header})
	for _, v := range userData {
		address := fmt.Sprintf("\"%s,%s,%s,%s,%s,%s\"", v.Address.Street, v.Address.Suite, v.Address.City, v.Address.Zipcode, v.Address.Geo.Lat, v.Address.Geo.Lng)
		company := fmt.Sprintf("\"%s,%s,%s\"", v.Company.Name, v.Company.CatchPhrase, v.Company.Bs)
		line := fmt.Sprintf("%d,%s,%s,%s,%s,%s,%s,%s", v.ID, v.Name, v.Username, v.Email, address, v.Phone, v.Website, company)
		utils.WriteCSV("output/users.csv", []string{line})
	}
}

func ScrapePosts(){
	postResponse, err := utils.GetResponse("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		panic(err)
	}
	var postData []types.PostData
	json.Unmarshal([]byte(postResponse), &postData)
	header := "id,userId,title,body"
	utils.WriteCSV("output/posts.csv", []string{header})
	for _, v := range postData {
		body := strings.Replace(v.Body, "\n", "\\n", -1)
		line := fmt.Sprintf("%d,%d,%s,%s", v.ID, v.UserId, v.Title, body)
		utils.WriteCSV("output/posts.csv", []string{line})
	}
}

func ScrapeComments(){
	commentsResponse, err := utils.GetResponse("https://jsonplaceholder.typicode.com/comments/")
	if err != nil {
		panic(err)
	}
	var commentsData []types.CommentData
	json.Unmarshal([]byte(commentsResponse), &commentsData)
	header := "id,postId,name,email,body"
	utils.WriteCSV("output/comments.csv", []string{header})
	for _, v := range commentsData {
		body := strings.Replace(v.Body, "\n", "\\n", -1)
		line := fmt.Sprintf("%d,%d,%s,%s,%s", v.ID, v.PostId, v.Name, v.Email, body)
		utils.WriteCSV("output/comments.csv", []string{line})
	}
}

func ScrapeAlbums(){
	albumsResponse, err := utils.GetResponse("https://jsonplaceholder.typicode.com/albums/")
	if err != nil {
		panic(err)
	}
	var albumsData []types.AlbumData
	json.Unmarshal([]byte(albumsResponse), &albumsData)
	header := "id,userId,title"
	utils.WriteCSV("output/albums.csv", []string{header})
	for _, v := range albumsData {
		line := fmt.Sprintf("%d,%d,%s", v.ID, v.UserId, v.Title)
		utils.WriteCSV("output/albums.csv", []string{line})
	}
}

func ScrapeTodos(){
	todosResponse, err := utils.GetResponse("https://jsonplaceholder.typicode.com/todos/")
	if err != nil {
		panic(err)
	}
	var todosData []types.TodoData
	json.Unmarshal([]byte(todosResponse), &todosData)
	header := "id,userId,title,completed"
	utils.WriteCSV("output/todos.csv", []string{header})
	for _, v := range todosData {
		line := fmt.Sprintf("%d,%d,%s,%t", v.ID, v.UserId, v.Title, v.Completed)
		utils.WriteCSV("output/todos.csv", []string{line})
	}
}