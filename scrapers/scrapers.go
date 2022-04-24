package scrapers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)
type ImgData struct {
	ID int `json:"id"`
	Title string `json:"title"`
	URL string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
	AlbumID int `json:"albumId"`
}
type UserData struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Address struct {
		Street string `json:"street"`
		Suite string `json:"suite"`
		City string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs string `json:"bs"`
	} `json:"company"`
}
type PostData struct {
	UserId int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}
type CommentData struct {
	PostId int `json:"postId"`
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Body string `json:"body"`
}


func GetResponse(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	return string(body), nil
}

func writeCSV(file string, data []string) error {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, v := range data {
		_, err := f.WriteString(v + "\n")
		f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func ScrapePhotos(){
	imgResponse, err := GetResponse("https://jsonplaceholder.typicode.com/photos/")
	if err != nil {
		panic(err)
	}
	unbuffered := make(chan ImgData)
	header := "id,title,url"
	writeCSV("output/photos.csv", []string{header})
	var data []ImgData
	json.Unmarshal([]byte(imgResponse), &data)
	dataLength := len(data)
	var resData []ImgData
	var wg sync.WaitGroup
	wg.Add(dataLength)
	for i := 0; i < dataLength; i++ {
		line := fmt.Sprintf("%d,%s,%s", data[i].ID, data[i].Title, data[i].URL)
		writeCSV("output/photos.csv", []string{line})
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
		_, err = file.Write(image)
		file.Close()
		if err != nil {
			panic(err)
		}
		}(i)
	}
	close(unbuffered)
	wg.Wait()
}
func ScrapeUsers(){
	userResponse, err := GetResponse("https://jsonplaceholder.typicode.com/users/")
	if err != nil {
		panic(err)
	}
	var userData []UserData
	json.Unmarshal([]byte(userResponse), &userData)
	header := "id,name,username,email,address,phone,website,company"
	writeCSV("output/users.csv", []string{header})
	for _, v := range userData {
		address := fmt.Sprintf("\"%s,%s,%s,%s,%s,%s\"", v.Address.Street, v.Address.Suite, v.Address.City, v.Address.Zipcode, v.Address.Geo.Lat, v.Address.Geo.Lng)
		company := fmt.Sprintf("\"%s,%s,%s\"", v.Company.Name, v.Company.CatchPhrase, v.Company.Bs)
		line := fmt.Sprintf("%d,%s,%s,%s,%s,%s,%s,%s", v.ID, v.Name, v.Username, v.Email, address, v.Phone, v.Website, company)
		writeCSV("output/users.csv", []string{line})
	}
}

func ScrapePosts(){
	postResponse, err := GetResponse("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		panic(err)
	}
	var postData []PostData
	json.Unmarshal([]byte(postResponse), &postData)
	header := "id,userId,title,body"
	writeCSV("output/posts.csv", []string{header})
	for _, v := range postData {
		body := strings.Replace(v.Body, "\n", "\\n", -1)
		line := fmt.Sprintf("%d,%d,%s,%s", v.ID, v.UserId, v.Title, body)
		writeCSV("output/posts.csv", []string{line})
	}
}

func ScrapeComments(){
	commentsResponse, err := GetResponse("https://jsonplaceholder.typicode.com/comments/")
	if err != nil {
		panic(err)
	}
	var commentsData []CommentData
	json.Unmarshal([]byte(commentsResponse), &commentsData)
	header := "id,postId,name,email,body"
	writeCSV("output/comments.csv", []string{header})
	for _, v := range commentsData {
		body := strings.Replace(v.Body, "\n", "\\n", -1)
		line := fmt.Sprintf("%d,%d,%s,%s,%s", v.ID, v.PostId, v.Name, v.Email, body)
		writeCSV("output/comments.csv", []string{line})
	}
}