package jpeg_test

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/viam-labs/go-libjpeg/jpeg"
	"github.com/viam-labs/go-libjpeg/test/util"
	"go.viam.com/test"
	"image"
	"image/color"
	nativeJPEG "image/jpeg"
	"io/ioutil"
	"os"
	"testing"
)

var naturalImageFiles = []string{
	"cosmos.jpg",
	"kinkaku.jpg",
}

var subsampledImageFiles = []string{
	"checkerboard_444.jpg",
	"checkerboard_440.jpg",
	"checkerboard_422.jpg",
	"checkerboard_420.jpg",
}

func TestMain(m *testing.M) {
	result := m.Run()
	if jpeg.SourceManagerMapLen() > 0 {
		fmt.Println("sourceManager leaked")
		result = 2
	}
	if jpeg.DestinationManagerMapLen() > 0 {
		fmt.Println("destinationManager leaked")
		result = 2
	}
	os.Exit(result)
}

func delta(u0, u1 uint32) int {
	d := int(u0) - int(u1)
	if d < 0 {
		return -d
	}
	return d
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, file := range naturalImageFiles {
			io := util.OpenFile(file)
			img, err := jpeg.Decode(io, &jpeg.DecoderOptions{})
			if img == nil {
				b.Error("Got nil")
			}
			if err != nil {
				b.Errorf("Got Error: %v", err)
			}
		}
	}
}

func BenchmarkDecodeIntoRGB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, file := range naturalImageFiles {
			io := util.OpenFile(file)
			img, err := jpeg.DecodeIntoRGB(io, &jpeg.DecoderOptions{})
			if img == nil {
				b.Error("Got nil")
			}
			if err != nil {
				b.Errorf("Got Error: %v", err)
			}
		}
	}
}

func BenchmarkDecodeWithNativeJPEG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, file := range naturalImageFiles {
			io := util.OpenFile(file)
			img, err := nativeJPEG.Decode(io)
			if img == nil {
				b.Error("Got nil")
			}
			if err != nil {
				b.Errorf("Got Error: %v", err)
			}
		}
	}
}

func TestDecode(t *testing.T) {
	for _, file := range naturalImageFiles {
		io := util.OpenFile(file)
		fmt.Printf(" - test: %s\n", file)

		img, err := jpeg.Decode(io, &jpeg.DecoderOptions{})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, img, test.ShouldNotBeNil)
		util.WritePNG(img, fmt.Sprintf("TestDecode_%s.png", file))
	}
}

func TestDecodeScaled(t *testing.T) {
	for _, file := range naturalImageFiles {
		io := util.OpenFile(file)
		fmt.Printf(" - test: %s\n", file)

		img, err := jpeg.Decode(io, &jpeg.DecoderOptions{ScaleTarget: image.Rect(0, 0, 100, 100)})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, img.Bounds().Dx(), test.ShouldEqual, 256)
		test.That(t, img.Bounds().Dy(), test.ShouldEqual, 192)
		test.That(t, img, test.ShouldNotBeNil)
		util.WritePNG(img, fmt.Sprintf("TestDecodeScaled_%s.png", file))
	}
}

func TestDecodeIntoRGBA(t *testing.T) {
	if jpeg.SupportRGBA() != true {
		t.Skipf("This build is not support DecodeIntoRGBA.")
		return
	}
	for _, file := range naturalImageFiles {
		io := util.OpenFile(file)
		fmt.Printf(" - test: %s\n", file)

		img, err := jpeg.DecodeIntoRGBA(io, &jpeg.DecoderOptions{})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, img, test.ShouldNotBeNil)
		util.WritePNG(img, fmt.Sprintf("TestDecodeIntoRGBA_%s.png", file))
	}
}

func TestDecodeScaledIntoRGBA(t *testing.T) {
	if jpeg.SupportRGBA() != true {
		t.Skipf("This build is not support DecodeIntoRGBA.")
		return
	}
	for _, file := range naturalImageFiles {
		io := util.OpenFile(file)
		fmt.Printf(" - test: %s\n", file)

		img, err := jpeg.DecodeIntoRGBA(io, &jpeg.DecoderOptions{ScaleTarget: image.Rect(0, 0, 100, 100)})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, img, test.ShouldNotBeNil)

		test.That(t, img.Bounds().Dx(), test.ShouldEqual, 256)
		test.That(t, img.Bounds().Dy(), test.ShouldEqual, 192)

		util.WritePNG(img, fmt.Sprintf("TestDecodeIntoRGBA_%s.png", file))
	}
}

func TestDecodeScaledIntoRGB(t *testing.T) {
	for _, file := range naturalImageFiles {
		io := util.OpenFile(file)
		fmt.Printf(" - test: %s\n", file)

		img, err := jpeg.DecodeIntoRGB(io, &jpeg.DecoderOptions{ScaleTarget: image.Rect(0, 0, 100, 100)})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, img, test.ShouldNotBeNil)

		test.That(t, img.Bounds().Dx(), test.ShouldEqual, 256)
		test.That(t, img.Bounds().Dy(), test.ShouldEqual, 192)

		util.WritePNG(img, fmt.Sprintf("TestDecodeIntoRGB_%s.png", file))
	}
}

func TestDecodeSubsampledImage(t *testing.T) {
	for _, file := range subsampledImageFiles {
		io := util.OpenFile(file)
		fmt.Printf(" - test: %s\n", file)

		img, err := jpeg.Decode(io, &jpeg.DecoderOptions{})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, img, test.ShouldNotBeNil)

		util.WritePNG(img, fmt.Sprintf("TestDecodeSubsampledImage_%s.png", file))
	}
}

func TestDecodeAndEncode(t *testing.T) {
	for _, file := range naturalImageFiles {
		io := util.OpenFile(file)
		fmt.Printf(" - test: %s\n", file)

		img, err := jpeg.Decode(io, &jpeg.DecoderOptions{})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, img, test.ShouldNotBeNil)

		// Create output file
		f, err := os.Create(util.GetOutFilePath(fmt.Sprintf("TestDecodeAndEncode_%s", file)))
		if err != nil {
			panic(err)
		}
		w := bufio.NewWriter(f)
		defer func() {
			w.Flush()
			f.Close()
		}()

		err = jpeg.Encode(w, img, &jpeg.EncoderOptions{Quality: 90})
		test.That(t, err, test.ShouldBeNil)
	}
}

func TestDecodeAndEncodeSubsampledImages(t *testing.T) {
	for _, file := range subsampledImageFiles {
		r := util.OpenFile(file)
		fmt.Printf(" - test: %s\n", file)

		img, err := jpeg.Decode(r, &jpeg.DecoderOptions{})
		test.That(t, err, test.ShouldBeNil)

		// Create output file
		f, err := os.Create(util.GetOutFilePath(fmt.Sprintf("TestDecodeAndEncodeSubsampledImages_%s", file)))
		test.That(t, err, test.ShouldBeNil)
		w := bufio.NewWriter(f)
		defer func() {
			w.Flush()
			f.Close()
		}()

		err = jpeg.Encode(w, img, &jpeg.EncoderOptions{Quality: 90})
		test.That(t, err, test.ShouldBeNil)
	}
}

func TestEncodeGrayImage(t *testing.T) {
	w, h := 400, 200
	img := image.NewGray(image.Rect(0, 0, w, h))

	// make gradient
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.SetGray(x, y, color.Gray{uint8(float64(x*y) / float64(w*h) * 255)})
		}
	}

	// encode gray gradient
	f, err := os.Create(util.GetOutFilePath(fmt.Sprintf("TestEncodeGrayImage_%dx%d.jpg", w, h)))
	test.That(t, err, test.ShouldBeNil)
	wr := bufio.NewWriter(f)
	defer func() {
		wr.Flush()
		f.Close()
	}()
	err = jpeg.Encode(wr, img, &jpeg.EncoderOptions{Quality: 90})
	test.That(t, err, test.ShouldBeNil)
	wr.Flush()

	// rewind to first
	f.Seek(0, 0)

	// decode file
	decoded, err := jpeg.Decode(f, &jpeg.DecoderOptions{})
	test.That(t, err, test.ShouldBeNil)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			r, g, b, _ := decoded.At(x, y).RGBA()
			ref := uint32(float64(x*y) / float64(w*h) * 255)
			if delta((r>>8), ref) > 1 || delta((g>>8), ref) > 1 || delta((b>>8), ref) > 1 {
				t.Errorf("(%d, %d): got (%d, %d, %d) want %v", x, y, r, g, b, ref)
			}
		}
	}
}

func TestDecodeConfig(t *testing.T) {
	for _, file := range naturalImageFiles {
		r := util.OpenFile(file)
		fmt.Printf(" - test: %s\n", file)

		config, err := jpeg.DecodeConfig(r)
		test.That(t, err, test.ShouldBeNil)

		test.That(t, config.ColorModel, test.ShouldEqual, color.YCbCrModel)
		test.That(t, config.Width, test.ShouldEqual, 1024)
		test.That(t, config.Height, test.ShouldEqual, 768)
	}
}

func TestNewYCbCrAlignedWithLandscape(t *testing.T) {
	img := jpeg.NewYCbCrAligned(image.Rect(0, 0, 125, 25), image.YCbCrSubsampleRatio444)

	test.That(t, len(img.Y), test.ShouldResemble, 6912)
	test.That(t, len(img.Cb), test.ShouldResemble, 6912)
	test.That(t, len(img.Cr), test.ShouldResemble, 6912)
	test.That(t, img.YStride, test.ShouldResemble, 144)
	test.That(t, img.CStride, test.ShouldResemble, 144)
}

func TestNewYCbCrAlignedWithPortrait(t *testing.T) {
	img := jpeg.NewYCbCrAligned(image.Rect(0, 0, 25, 125), image.YCbCrSubsampleRatio444)

	test.That(t, len(img.Y), test.ShouldResemble, 6912)
	test.That(t, len(img.Cb), test.ShouldResemble, 6912)
	test.That(t, len(img.Cr), test.ShouldResemble, 6912)
	test.That(t, img.YStride, test.ShouldResemble, 48)
	test.That(t, img.CStride, test.ShouldResemble, 48)
}

func TestDecodeFailsWithBlankFile(t *testing.T) {
	blank := bytes.NewBuffer(nil)
	_, err := jpeg.Decode(blank, &jpeg.DecoderOptions{})
	test.That(t, err, test.ShouldNotBeNil)
}

func TestEncodeFailsWithEmptyImage(t *testing.T) {
	dummy := &image.YCbCr{}
	w := bytes.NewBuffer(nil)
	err := jpeg.Encode(w, dummy, &jpeg.EncoderOptions{})
	test.That(t, err, test.ShouldNotBeNil)
}
func newRGBA() *image.RGBA {
	rgba := image.NewRGBA(image.Rect(0, 0, 4, 8))
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			rgba.SetRGBA(i, j, color.RGBA{255, 0, 0, 255})
		}
		for j := 4; j < 8; j++ {
			rgba.SetRGBA(i, j, color.RGBA{0, 0, 255, 255})
		}
	}
	return rgba
}

func TestEncodeRGBA(t *testing.T) {
	rgba := newRGBA()
	w := bytes.NewBuffer(nil)

	err := jpeg.Encode(w, rgba, &jpeg.EncoderOptions{
		Quality: 100,
	})
	test.That(t, err, test.ShouldBeNil)

	decoded, err := jpeg.Decode(w, &jpeg.DecoderOptions{})
	test.That(t, err, test.ShouldBeNil)
	//test.That(t, decoded, test.ShouldNotBeNil)

	diff, err := util.MatchImage(rgba, decoded, 1)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, diff, test.ShouldBeNil)
	if err != nil {
		t.Errorf("match image: %v", err)
		util.WritePNG(rgba, "TestEncodeRGBA.want.png")
		util.WritePNG(decoded, "TestEncodeRGBA.got.png")
		util.WritePNG(diff, "TestEncodeRGBA.diff.png")
	}
}

// See: https://github.com/viam-labs/go-libjpeg/issues/36
func TestDecodeAndEncodeRGBADisableFancyUpsampling(t *testing.T) {
	src := image.NewRGBA(image.Rect(0, 0, 3000, 2000))

	w, err := ioutil.TempFile("", "jpeg_test_")
	test.That(t, err, test.ShouldBeNil)
	name := w.Name()
	defer os.Remove(w.Name())

	err = jpeg.Encode(w, src, &jpeg.EncoderOptions{Quality: 95})
	w.Close()
	test.That(t, err, test.ShouldBeNil)

	r, err := os.Open(name)
	test.That(t, err, test.ShouldBeNil)
	defer r.Close()

	_, err = jpeg.DecodeIntoRGBA(r, &jpeg.DecoderOptions{
		DisableBlockSmoothing:  true,
		DisableFancyUpsampling: true,
	})
	test.That(t, err, test.ShouldBeNil)
}
