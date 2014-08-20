package zero

// #include <stdlib.h>
// #include <SDL2/SDL.h>
// #include "game.h"
import "C"
import (
	"errors"
	"unsafe"
)

type Texture struct {
	tex *C.SDL_Texture
}

var Textures map[string]Texture

func init() {
	Textures = make(map[string]Texture)
}

func LoadBMP(path, name string) error {
	// Make sure this name isn't already in use
	if _, ok := Textures[name]; ok {
		return errors.New("Texture with name " + name + " already in use")
	}

	// Convert the path to a c-string
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	// Load the surface
	surface := C.loadBMP(cpath)
	if surface == nil {
		return errors.New("SDL_LoadBMP")
	}
	defer C.SDL_FreeSurface(surface)

	// Create the texture
	tex := C.SDL_CreateTextureFromSurface(renderer, surface)
	if tex == nil {
		return errors.New("SDL_CreateTextureFromSurface")
	}

	Textures[name] = Texture{tex: tex}

	return nil
}

func destroyTextures() {
	for _, t := range Textures {
		C.SDL_DestroyTexture(t.tex)
	}
}
