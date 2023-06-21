//go:build dynamic

package jpeg

/*
#cgo windows LDFLAGS: -ljpeg
#cgo !windows pkg-config: libjpeg
*/
import "C"
