package main

import (
	"github.com/andybons/gogif"
	"github.com/nfnt/resize"
	"image"
	"image/gif"

	"os"
)

func processImage(img image.Image) image.Image {
	return resize.Resize(150, 0, img, resize.Bilinear)
}

func imageToPaletted(img image.Image) *image.Paletted {
	pm, ok := img.(*image.Paletted)
	if !ok {
		b := img.Bounds()
		pm = image.NewPaletted(b, nil)
		q := &gogif.MedianCutQuantizer{NumColor: 256}
		q.Quantize(pm, b, img, image.ZP)
	}
	return pm
}

func makeGif(dts []*data, path string) (err error) {
	outGif := &gif.GIF{}
	for _, d := range dts {
		palettedImage := imageToPaletted(processImage(d.Img))

		// Add new frame to animated GIF
		outGif.Image = append(outGif.Image, palettedImage)
		outGif.Delay = append(outGif.Delay, 0)
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	gif.EncodeAll(f, outGif)

	return err
}
