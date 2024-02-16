package middleware

import (
	"io"
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
)

func UsePosterRouter(api fiber.Router) {
	api.Static("/poster", "/tmp", fiber.Static{
		CacheDuration: 0,
		Index:         "defaultposter.jpg",
		Browse:        true,
	})
}

func CopyImage(source string, destination ...string) error {
	wd, _ := os.Getwd()
	target := path.Join(destination...)
	target = path.Join(wd, target)

	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(target)
	if err != nil {
		return err
	}
	defer dst.Close()

	io.Copy(dst, src)

	return nil
}
