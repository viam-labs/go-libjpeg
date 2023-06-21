// Package util contains utility code for demonstrations of go-libjpeg.
package util

import (
	"bufio"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

// SubsampledImages contains example images which have subsampling colors.
var SubsampledImages = []string{
	"checkerboard_444.jpg",
	"checkerboard_440.jpg",
	"checkerboard_422.jpg",
	"checkerboard_420.jpg",
}

// ResolveFile returns the path of the given file relative to the root
// of the codebase. For example, if this file currently
// lives in utils/file.go and ./foo/bar/baz is given, then the result
// is foo/bar/baz. This is helpful when you don't want to relatively
// refer to files when you're not sure where the caller actually
// lives in relation to the target file.
func ResolveFile(fn string) string {
	//nolint:dogsled
	_, thisFilePath, _, _ := runtime.Caller(0)
	thisDirPath, err := filepath.Abs(filepath.Dir(thisFilePath))
	if err != nil {
		panic(err)
	}
	return filepath.Join(thisDirPath, "../..", fn)
}

// GetExFilePath returns the path of specified example file.
func GetExFilePath(name string) string {
	return ResolveFile("./test/images/" + name)
}

// GetOutFilePath returns the path of specified out file.
func GetOutFilePath(name string) string {
	return ResolveFile("./test/out/" + name)
}

// OpenFile opens specified example file
func OpenFile(name string) (io io.Reader) {
	io, err := os.Open(GetExFilePath(name))
	if err != nil {
		panic(err)
	}
	return
}

// ReadFile reads and returns data bytes of specified example file.
func ReadFile(name string) (data []byte) {
	data, err := ioutil.ReadFile(GetExFilePath(name))
	if err != nil {
		panic(err)
	}
	return
}

// CreateFile opens specified example file
func CreateFile(name string) (f *os.File) {
	f, err := os.Create(GetOutFilePath(name))
	if err != nil {
		panic(err)
	}
	return
}

// WritePNG encodes and writes image into PNG file.
func WritePNG(img image.Image, name string) {
	f, err := os.Create(GetOutFilePath(name))
	if err != nil {
		panic(err)
	}
	b := bufio.NewWriter(f)
	defer func() {
		b.Flush()
		f.Close()
	}()

	if err := png.Encode(b, img); err != nil {
		panic(err)
	}
	return
}
