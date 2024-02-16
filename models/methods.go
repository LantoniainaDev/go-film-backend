package models

import (
	"fmt"
	"path"
)

func (f *Film) Validate() error {
	//validation du film
	f.Title = validateTitle(f.Title)

	// validation de l'URL
	f.Url = validateUrl(f.Url)

	// validation du poster
	f.Poster = validatePoster(f.Url)

	// validation du cover
	if f.Cover == "" {
		f.Cover = f.Poster
	}

	// validation de la source
	if err := validaterSource(f.Source, false); err != nil {
		return err
	}

	// pas de validation du Date
	return nil
}

func (f *Film) PreSave(skip ...bool) error {
	if len(skip) != 0 {
		if skip[0] {
			return nil
		}
	}
	// creation de l'image
	pathToVideo := f.Source
	fmt.Println("presaving film", pathToVideo)
	output := path.Join("film", f.Poster)

	err := thumble(pathToVideo, output)

	return err
}

// compiles

func (f *Compile) Validate() error {
	//validation de la compile
	f.Title = validateTitle(f.Title)

	// validation de l'URL
	f.Url = validateUrl(f.Url)

	// validation du poster
	f.Poster = validatePoster(f.Poster)

	// validation du cover
	if f.Cover == "" {
		f.Cover = f.Poster
	}

	// validation de la source
	if err := validaterSource(f.Source, true); err != nil {
		return err
	}

	// recherche et listing des episodes
	f.Episodes = vaidateEpisodes(f.Source)

	return nil
}

func (f *Compile) PreSave(param string) error {
	var err error
	// parcours du tableau episode
	for i := 0; i < len(f.Episodes); i++ {
		pathToVideo := path.Join(f.Source, f.Episodes[i])
		fmt.Println("path to video", pathToVideo)

		output := path.Join(param, f.Url, fmt.Sprintf("%d.jpeg", i))

		// creation de l'image
		err = thumble(pathToVideo, output)
		if err != nil {
			break
		}

	}

	return err
}
