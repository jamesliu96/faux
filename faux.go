package faux

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"image"
	"io"

	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
)

func Faux(r io.Reader, w io.Writer, key []byte) (err error) {
	m, _, err := image.Decode(r)
	if err != nil {
		return
	}
	bounds := m.Bounds()
	nm := image.NewRGBA(bounds)
	s := sha512.Sum384(key)
	b, err := aes.NewCipher(s[:32])
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(nil)
	sw := cipher.StreamWriter{S: cipher.NewCTR(b, s[32:32+aes.BlockSize]), W: buf}
	width, height := bounds.Dx(), bounds.Dy()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := color.RGBAModel.Convert(m.At(x, y)).(color.RGBA)
			sw.Write([]byte{c.R, c.G, c.B})
			buf.Write([]byte{c.A})
		}
	}
	nm.Pix = buf.Bytes()
	err = png.Encode(w, nm)
	return
}
