//assets/assets.go
package assets

import "embed"

//go:embed public/*
var Assets embed.FS

//go:embed public/Roboto-Regular.ttf
var RobotoTTF []byte

//go:embed public/DS-DIGIB.TTF
var DigiTTF []byte