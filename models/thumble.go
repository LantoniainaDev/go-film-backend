package models

import (
	"fmt"
	"os"
	"path"
	"time"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func thumble(target string, output string) error {
	var errChan chan error
	var done chan bool
	var err error

	go func() {
		if err = createImage(target, output, "300x200"); err != nil {
			errChan <- err
			done <- true
			return
		}

		output = output + "small.jpeg"
		if err = createImage(target, output, "15x10"); err != nil {
			errChan <- err
		}
		done <- true
	}()

	go func() {
		for {
			did := <-done
			if did {
				break
			}
			time.Sleep(time.Second * 5)
			fmt.Println("not already done")
		}
		err = <-errChan
		if err != nil {
			fmt.Println("successfully done")
		} else {
			fmt.Println("done with error:", err)
		}
	}()

	return err
}

func createImage(target string, output string, size string) error {
	var err error

	wd, _ := os.Getwd()

	output = path.Join(wd, "tmp", output)

	//s'assurer aue le dossier existe
	dossier := path.Dir(output)
	fmt.Println("dossier", dossier)
	if _, dontExist := os.ReadDir(dossier); dontExist != nil {
		err = os.MkdirAll(dossier, os.ModeDir)
		if err != nil {
			return err
		}
	}

	args := ffmpeg_go.KwArgs{
		"ss":      "4",
		"s":       size,
		"vframes": "1",
		"f":       "image2",
	}

	// le poster normal
	stream := ffmpeg_go.Input(target)
	err = stream.Output(output, args).OverWriteOutput().Run()

	return err
}
