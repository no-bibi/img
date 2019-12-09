package img

import (
	"image"
	"image/color"
)

//convert img to gray
//Gray = (Red * 0.3 + Green * 0.59 + Blue * 0.11)
func Gray(img image.Image) image.Image {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	newRgba := image.NewRGBA(bounds)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			colorRgb := img.At(x, y)
			rr, gg, bb, aa := colorRgb.RGBA()
			floatG := float64(rr>>8)*0.3 + float64(gg>>8)*0.59 + float64(bb>>8)*0.11
			newG := uint8(floatG)
			newA := uint8(aa >> 8)
			newRgba.SetRGBA(x, y, color.RGBA{R: newG, G: newG, B: newG, A: newA})
		}
	}
	return newRgba
}

//convert img to binaryzation
// rbg >= threshold is black
// rbg < threshold is white
func Binary(img image.Image, threshold uint8) image.Image {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	newRgba := image.NewRGBA(bounds)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			imageColor := img.At(x, y)
			rr, gg, bb, _ := imageColor.RGBA()
			rgb := uint8(((rr >> 8) + (gg >> 8) + (bb >> 8)) / 3)

			if rgb >= threshold {
				rgb = 255
			} else {
				rgb = 0
			}
			newRgba.SetRGBA(x, y, color.RGBA{R: rgb, G: rgb, B: rgb, A: 255})
		}
	}
	return newRgba
}

//convert img to reverse color
func Reverse(img image.Image) image.Image {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	newRgba := image.NewRGBA(bounds)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			imageColor := img.At(x, y)
			rr, gg, bb, aa := imageColor.RGBA()
			r := uint8(255 - (rr >> 8))
			g := uint8(255 - (gg >> 8))
			b := uint8(255 - (bb >> 8))
			a := uint8(aa >> 8)
			newRgba.SetRGBA(x, y, color.RGBA{R: r, G: g, B: b, A: a})
		}
	}
	return newRgba
}
