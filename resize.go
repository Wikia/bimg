// +build go1.7

package bimg

import (
	"runtime"
)

// Resize is used to transform a given image as byte buffer
// with the passed options.
func Resize(buf []byte, o Options) ([]byte, error) {
	// Required in order to prevent premature garbage collection. See:
	// https://github.com/h2non/bimg/pull/162
	defer runtime.KeepAlive(buf)
	return resizer(buf, o)
}

// Simulate the window-crop-fixed operation of Vignette
// see https://github.com/Wikia/vignette/#window-crop-fixed
func WindowCropFixed(buf []byte, o Options) ([]byte, error) {
	// Required in order to prevent premature garbage collection. See:
	// https://github.com/h2non/bimg/pull/162
	defer runtime.KeepAlive(buf)
	return windowcropfixed(buf, o)
}
