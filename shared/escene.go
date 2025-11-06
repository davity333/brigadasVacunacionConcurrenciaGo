// shared/scene.go
package shared

import "github.com/hajimehoshi/ebiten/v2"

// Scene es la interfaz que todas las escenas deben implementar
type Scene interface {
    Update() error
    Draw(screen *ebiten.Image)
}

// CurrentScene es la escena activa del juego
var CurrentScene Scene
