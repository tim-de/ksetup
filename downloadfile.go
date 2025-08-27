package main

import (
	"fmt"
	"encoding/json"
	"errors"
	"os"
	"net/http"
	"io"
)

const buflen int = 2048

type TargetInfo struct {
	filename string
	downloadurl string
}

func getDownloadInfo(reponame string) (TargetInfo, error) {
	var info TargetInfo
	var data map[string]any
	var rawdata []byte
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", reponame)

	response, err := http.Get(url)
	if err != nil {
		return TargetInfo{}, err
	}
	defer response.Body.Close()

	for {
		buffer := make([]byte, buflen)
		nread, err := response.Body.Read(buffer)
		if err != nil && err != io.EOF {
			return TargetInfo{}, err
		}
		rawdata = append(rawdata, buffer[:nread]...)
		if nread < buflen || err == io.EOF {
			break
		}
	}

	if err := json.Unmarshal(rawdata, &data); err != nil {
		return TargetInfo{}, err
	}
	info.filename = data["assets"].([]any)[0].(map[string]any)["name"].(string)
	info.downloadurl = data["assets"].([]any)[0].(map[string]any)["browser_download_url"].(string)
	return info, nil
}

func downloadFile(url string, filepath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, response.Body)
	return err
}

func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)
}
