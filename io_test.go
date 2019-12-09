package img

import (
	"image"
	"testing"
)

func TestEncode(t *testing.T) {
	bounds := image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{200, 200},
	}
	imgPng := image.NewNRGBA(bounds)
	imgJpg := image.NewNRGBA(bounds)
	imgGif := image.NewNRGBA(bounds)

	type args struct {
		img image.Image
		fp  string
		o   []interface{}
	}
	var tests = []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "png",
			args:    args{img: imgPng, fp: `./build/imgpng.png`},
			wantErr: false,
		},
		{
			name:    "jpg",
			args:    args{img: imgJpg, fp: `./build/imgjpg.jpg`},
			wantErr: false,
		},
		{
			name:    "gif",
			args:    args{img: imgGif, fp: `./build/imgjif.gif`},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Encode(tt.args.img, tt.args.fp, tt.args.o...); (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRead(t *testing.T) {
	type args struct {
		fp string
	}
	tests := []struct {
		name    string
		args    args
		wantImg image.Image
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "first",
			args:    args{`./source/enter.png`},
			wantImg: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Read(tt.args.fp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
