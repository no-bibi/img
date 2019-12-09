package img

import (
	"image"
	"reflect"
	"testing"
)

func TestBinary(t *testing.T) {
	type args struct {
		img       image.Image
		threshold uint8
	}
	tests := []struct {
		name string
		args args
		want image.Image
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Binary(tt.args.img, tt.args.threshold); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Binaryzation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGray(t *testing.T) {
	type args struct {
		img image.Image
	}
	tests := []struct {
		name string
		args args
		want image.Image
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gray(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		img image.Image
	}
	tests := []struct {
		name string
		args args
		want image.Image
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
