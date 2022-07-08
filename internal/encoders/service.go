package encoders

import (
	"image"
	"io"
)

// Service creates encoder instances
type ServiceInterface interface {
	NewEncoder(codec VideoCodec, size image.Point, frameRate int) (EncoderInterface, error)
	Supports(codec VideoCodec) bool
}

// Encoder takes an image/frame and encodes it
type EncoderInterface interface {
	io.Closer
	Encode(*image.RGBA) ([]byte, error)
	VideoSize() (image.Point, error)
}

//VideoCodec can be either h264 or vp8
type VideoCodec = int

const (
	//NoCodec "zero-value"
	NoCodec VideoCodec = iota
	//H264Codec h264
	H264Codec
	//VP8Codec vp8
	VP8Codec
)
