package utils

import (
	"os"
)

func ResetFiles() {
	// if _, err := os.Stat("output/photos.csv"); err == nil {
	// 	os.Remove("output/photos.csv")
	// }
	// if _, err := os.Stat("users.csv"); err == nil {
	// 	os.Remove("users.csv")
	// }
	// f, err := os.Create("output/photos.csv")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer f.Close()
	// f2, err := os.Create("output/users.csv")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer f2.Close()
	// if _, err := os.Stat("output/photos.csv"); err == nil {
	// 	os.Remove("output/photos.csv")
	// }
	// f3, err := os.Create("output/posts.csv")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer f3.Close()
	// if _, err := os.Stat("output/comments.csv"); err == nil {
	// 	os.Remove("output/comments.csv")
	// }
	// f4, err := os.Create("output/comments.csv")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer f4.Close()
	// if _, err := os.Stat("output/photos"); err == nil {
	// 	os.RemoveAll("output/photos")
	// }
	if _, err := os.Stat("output"); err == nil {
		os.RemoveAll("output")
	}
	os.MkdirAll("output/photos", 0777)
}