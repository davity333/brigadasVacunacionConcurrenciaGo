// shared/logicaImagenes.go
package shared

import (
    "image"
    "log"
    _ "image/jpeg"
    "github.com/hajimehoshi/ebiten/v2"
    "multi/assets"
)

// LoadImage carga una imagen embebida desde assets
func LoadImage(path string) *ebiten.Image {
    f, err := assets.Assets.Open(path)
    if err != nil {
        log.Fatalf("Error al abrir %s: %v", path, err)
    }
    defer f.Close()

    img, _, err := image.Decode(f)
    if err != nil {
        log.Fatalf("Error al decodificar %s: %v", path, err)
    }
    return ebiten.NewImageFromImage(img)
}

// DrawCenteredImage dibuja una imagen centrada en pantalla con escala
func DrawCenteredImage(screen *ebiten.Image, img *ebiten.Image, escala float64) {
    w, h := img.Bounds().Dx(), img.Bounds().Dy()
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Scale(escala, escala)
    op.GeoM.Translate(-float64(w)*escala/2, -float64(h)*escala/2)
    x := float64(screen.Bounds().Dx()) / 2
    y := float64(screen.Bounds().Dy()) / 2
    op.GeoM.Translate(x, y)
    screen.DrawImage(img, op)
}

// IsMouseOverArea verifica si el cursor está sobre un área
func IsMouseOverArea(x, y, w, h float64) bool {
    mx, my := ebiten.CursorPosition()
    return float64(mx) >= x && float64(mx) <= x+w &&
        float64(my) >= y && float64(my) <= y+h
}

