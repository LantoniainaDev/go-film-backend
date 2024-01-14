package models

import (
	"errors"
	"os"
	"strings"
)

type Img struct {
	Film
}

func validateUrl(url string) string {

	// trim
	url = strings.Trim(url, "")
	url = strings.Replace(url, " ", "-", -3)

	// case checking
	url = strings.ToLower(url)

	// checking for the slash
	url = slash(url)
	return url
}

func slash(str string) string {
	if strings.Index(str, "/") != 0 {
		str = "/" + str
	}

	return str
}

func validateTitle(title string) string {
	title = strings.Trim(title, "")

	return title
}

func validaterSource(path string, dir bool) error {
	var err error
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	ok := info.IsDir() == dir
	if !ok {
		err = errors.New("le chemin ne match pas le type recherché")
	}

	return err
}
