package utils

import (
	"io/ioutil"
	"net/http"
	"os"
)

func ResetFiles() {
	if _, err := os.Stat("output"); err == nil {
		os.RemoveAll("output")
	}
	os.MkdirAll("output/photos", 0777)
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

func WriteCSV(file string, data []string) error {
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