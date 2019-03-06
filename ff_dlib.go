package main

import (
	"bytes"
	"errors"
	gf "github.com/Kagami/go-face"
	"gocv.io/x/gocv"
	"image"
	"image/jpeg"
	//"log"
	//"strconv"
)

func fixBound(main, sub *image.Rectangle) {
	if sub.Min.X < 0 {
		sub.Min.X = 0
	}
	if sub.Min.Y < 0 {
		sub.Min.Y = 0
	}
	if sub.Max.X > main.Max.X {
		sub.Max.X = main.Max.X
	}
	if sub.Max.Y > main.Max.Y {
		sub.Max.Y = main.Max.Y
	}
}

func getdataDlib(link string) (dts []data, err error) {
	cap, err := gocv.OpenVideoCapture(link)
	if err != nil {
		return dts, errors.New("can not capture " + link)
	}
	defer cap.Close()

	img := gocv.NewMat()
	defer img.Close()

	rec, err := gf.NewRecognizer(RECOGNIZER_DIR)
	if err != nil {
		return dts, err
	}
	defer rec.Close()

	numberFrames := int(cap.Get(gocv.VideoCaptureFrameCount))

	for fr := 0; fr < numberFrames; fr++ {
		if ok := cap.Read(&img); !ok {
			break
		}
		buf := new(bytes.Buffer)
		cvImg, _ := img.ToImage()
		jpeg.Encode(buf, cvImg, nil)
		jb := buf.Bytes()

		mainBound := cvImg.Bounds()

		gdatas, err := rec.Recognize(jb)
		if err != nil {
			return dts, err
		}
		for _, f := range gdatas {
			d := data{
				Frame:     fr,
				Vector:    f.Descriptor,
				Rectangle: f.Rectangle,
			}
			fixBound(&mainBound, &d.Rectangle)

			cropMat := gocv.NewMat()
			crop := img.Region(d.Rectangle)
			crop.CopyTo(&cropMat)
			d.Img, err = cropMat.ToImage()
			cropMat.Close()
			if err != nil {
				return dts, err
			}

			dts = append(dts, d)
		}
	}

	return dts, err
}
