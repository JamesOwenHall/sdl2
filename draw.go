package sdl2

// #include <stdlib.h>
// #include <SDL2/SDL.h>
// #include "game.h"
import "C"

type Color struct {
	R, G, B uint8
}

func DrawTexture(tex Texture, x, y, w, h int) {
	dest := C.struct_SDL_Rect{x: C.int(x), y: C.int(y), w: C.int(w), h: C.int(h)}
	C.SDL_RenderCopy(renderer, tex.tex, nil, &dest)
}

func SetColor(c Color) {
	C.SDL_SetRenderDrawColor(renderer, C.Uint8(c.R), C.Uint8(c.G), C.Uint8(c.B), C.Uint8(255))
}

func Clear() {
	C.SDL_RenderClear(renderer)
}
