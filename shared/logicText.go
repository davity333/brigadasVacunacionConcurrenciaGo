package shared
import (
    "golang.org/x/image/font/opentype"
    "golang.org/x/image/font"
)

var(
    lsdNumero font.Face
)

func FontLoadNumero(err error, fontData []byte, size float64) {
    if err != nil {
        panic(err)
    }

    ttf, err := opentype.Parse(fontData)
    if err != nil {
        panic(err)
    }

    lsdNumero, err = opentype.NewFace(ttf, &opentype.FaceOptions{
        Size:    size,
        DPI:     72,
        Hinting: font.HintingFull,
    })
    if err != nil {
        panic(err)
    }
}