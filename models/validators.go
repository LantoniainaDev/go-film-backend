package models

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

type Img struct {
	Film
}

func validateUrl(url string) string {

	// trim
	url = strings.Trim(url, "")
	url = strings.Replace(url, " ", "-", -3)

	if url == "" {
		url = "pipol"
	}

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

func validatePoster(poster string) string {
	poster = slash(poster)
	if path.Ext(poster) == ".jpeg" {
		return poster
	} else {
		return poster + ".jpeg"
	}
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

func vaidateEpisodes(source string) []string {
	var episodes []string

	episodesEntries, err := os.ReadDir(source)
	if err != nil {
		fmt.Println("une erreur es survenue")
		return episodes
	}

	for i := 0; i < len(episodesEntries); i++ {
		entry := episodesEntries[i]
		if entry.IsDir() {
			continue
		}
		episodes = append(episodes, entry.Name())
	}

	return episodes
}
