// lcd.go
package scenes

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
    "multi/shared"
    "fmt"
)

var (
    lcdModal  *ebiten.Image
    btrCerrar *ebiten.Image
)

var (
    MostrarLcd bool = false
)

func InitLcd() {
    lcdModal = shared.LoadImage("public/lcd.png")
    btrCerrar = shared.LoadImage("public/btnCerrarModal.png")
}

func DrawModalLcd(screen *ebiten.Image) {
    if !MostrarLcd {
        return
    }

    // Fondo negro translúcido
    w, h := screen.Size()
    overlay := ebiten.NewImage(w, h)
    overlay.Fill(color.RGBA{0, 0, 0, 150})
    screen.DrawImage(overlay, nil)

    // Modal centrado
    if lcdModal != nil {
        shared.DrawCenteredImage(screen, lcdModal, 0.3)
    }

    // Botón cerrar EXACTO como indicaste
    if btrCerrar != nil {
        shared.DrawImagen(screen, btrCerrar, 0.2, 683, 106)
    }
}

func DetectarClickCerrar() {
    if !MostrarLcd || btrCerrar == nil {
        return
    }

    if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        x, y := ebiten.CursorPosition()

        escala := 8.0
        ancho := int(float64(btrCerrar.Bounds().Dx()) * escala)
        alto  := int(float64(btrCerrar.Bounds().Dy()) * escala)

        // Ajuste porque DrawImagen centra la imagen en (x,y)
        inicioX := 808 - (ancho / 2) +420
        inicioY := 90 - (alto / 2) + 420
        finX := inicioX + ancho
        finY := inicioY + alto

        fmt.Printf("Área: X[%d-%d] Y[%d-%d] Cursor: (%d,%d)\n",
            inicioX, finX, inicioY, finY, x, y)

        if x >= inicioX && x <= finX && y >= inicioY && y <= finY {
            fmt.Println("✅ Click dentro del área, cerrando modal")
            MostrarLcd = false
        } else {
            fmt.Println("❌ Click fuera del área")
        }
    }
}
