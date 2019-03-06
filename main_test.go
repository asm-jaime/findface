package main

import (
	"image"
	"image/jpeg"
	"os"
	"strconv"
	"testing"
)

func imgToFile(img *image.Image, name string) error {
	out, err := os.Create(name)
	if err != nil {
		return err
	}
	err = jpeg.Encode(out, *img, nil)
	return err
}

func TestGetdataDlib(t *testing.T) {
	dts, err := getdataDlib("recognizer/std.webm")
	if err != nil {
		t.Error(err)
		return
	}
	for i, d := range dts {
		err := imgToFile(&d.Img, "data/"+strconv.Itoa(i)+".jpeg")
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func TestGetClasters(t *testing.T) {
	dts, err := getdataDlib("recognizer/std.webm")
	if err != nil {
		t.Error(err)
		return
	}
	cls := getClasters(dts)
	for i := range cls {
		err = makeGif(cls[i], RECOGNIZER_DIR+"/"+strconv.Itoa(i)+".gif")
		if err != nil {
			t.Error(err)
			return
		}
	}
}
