//go:build !dynamic

package jpeg

/*
Note the contents of include/ are extracted from jpeg-turbo 2.1.5.1.
Contents of lib/ have been built from that source as well.
Libraries are actually copies of libturbojpeg.a as built for each architecture.
*/

/*
#cgo linux,darwin CFLAGS: -I${SRCDIR}/include
#cgo windows CFLAGS: -I${SRCDIR}/include/include_windows

#cgo linux,amd64 LDFLAGS: ${SRCDIR}/lib/libturbojpeg_linux_amd64.a
#cgo linux,arm64 LDFLAGS: ${SRCDIR}/lib/libturbojpeg_linux_arm64.a
#cgo darwin,amd64 LDFLAGS: ${SRCDIR}/lib/libturbojpeg_darwin_amd64.a
#cgo darwin,arm64 LDFLAGS: ${SRCDIR}/lib/libturbojpeg_darwin_arm64.a
#cgo windows LDFLAGS: ${SRCDIR}/lib/libturbojpeg_windows.a
*/
import "C"
