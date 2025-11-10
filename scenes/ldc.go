// lcd.go
package scenes

import (
    "fmt"
    "image/color"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/text"
    "golang.org/x/image/font"
    "golang.org/x/image/font/opentype"
    "multi/logic" 
    "multi/assets" // aquí importas tu paquete assets con embed
    "multi/shared"
)

var (
    lcdModal  *ebiten.Image
    btrCerrar *ebiten.Image
    digiFont  font.Face
)

var (
    MostrarLcd bool = false
    numA int = 123
    numB int = 456
    numC int = 789
)

// Inicializa imágenes y fuente embebida
func InitLcd() {
    lcdModal = shared.LoadImage("public/lcd.png")
    btrCerrar = shared.LoadImage("public/btnCerrarModal.png")

    // Cargar fuente desde embed
    tt, err := opentype.Parse(assets.DigiTTF)
    if err != nil {
        panic(err)
    }
    digiFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
        Size:    48,
        DPI:     72,
        Hinting: font.HintingFull,
    })
    if err != nil {
        panic(err)
    }
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

    // Botón cerrar
    if btrCerrar != nil {
        shared.DrawImagen(screen, btrCerrar, 0.2, 683, 106)
    }

    // Color verde #3B7D23
    verdeLCD := color.RGBA{59, 125, 35, 255}

    // Dibujar los tres valores en vertical, un poco más a la derecha
    if digiFont != nil {
        baseX := w/2 + 110 // mantener la misma posición
        baseY := h/2 - 30  // mantener la misma posición inicial

        // Mostrar en orden: Interior, Exterior, Humedad
        text.Draw(screen, fmt.Sprintf("%d°C INT", logic.TempInterior), digiFont, baseX, baseY, verdeLCD)
        text.Draw(screen, fmt.Sprintf("%d°C EXT", logic.TempExterior), digiFont, baseX, baseY+50, verdeLCD)
        text.Draw(screen, fmt.Sprintf("%d%% HUM", logic.Humedad), digiFont, baseX, baseY+100, verdeLCD)
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
        inicioX := 808
        inicioY := 90
        finX := inicioX + ancho
        finY := inicioY + alto

        fmt.Printf("Área: X[%d-%d] Y[%d-%d] Cursor: (%d,%d)\n",
            inicioX, finX, inicioY, finY, x, y)

        if x >= inicioX && x <= finX && y >= inicioY && y <= finY {
            fmt.Println("Click dentro del área, cerrando modal")
            MostrarLcd = false
        } else {
            fmt.Println("Click fuera del área")
        }
    }
}
