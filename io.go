package img

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
)

//read file
func Read(fp string) (img image.Image, err error) {
	file, err := os.Open(fp)
	if err != nil {
		return
	}
	defer file.Close()

	img, _, err = image.Decode(file)
	return
}

//save file
func Encode(img image.Image, fp string, o ...interface{}) (err error) {

	file, err := os.Create(fp)
	if err != nil {
		return
	}
	defer file.Close()

	ext := path.Ext(fp)

	switch ext {
	case `.png`:
		return png.Encode(file, img)
	case `.jpg`:
		fallthrough
	case `.jpeg`:
		options := &jpeg.Options{Quality: 50}
		if len(o) > 0 {
			options = o[0].(*jpeg.Options)
		}
		return jpeg.Encode(file, img, options)
	case `.gif`:
		options := &gif.Options{
			NumColors: 127,
			Quantizer: nil,
			Drawer:    nil,
		}
		if len(o) > 0 {
			options = o[0].(*gif.Options)
		}

		return gif.Encode(file, img, options)
	}
	return errors.New(`not support image type : ` + ext)
}
