package main

import (
	gf "github.com/Kagami/go-face"
	"image"
)

type data struct {
	Frame     int
	Rectangle image.Rectangle
	Vector    gf.Descriptor
	Img       image.Image
}

type clasters [][]*data
