package sdl2

// #cgo CFLAGS: -Wno-implicit-function-declaration
// #cgo LDFLAGS: -framework SDL2
// #include <stdlib.h>
// #include <SDL2/SDL.h>
// #include "game.h"
import "C"
import (
	"errors"
	"unsafe"
)

var (
	window   *C.SDL_Window
	renderer *C.SDL_Renderer
)

func LaunchGame(title string, width, height int) error {
	// Convert the title to a c-string
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	// Initialize SDL
	if C.SDL_Init(C.SDL_INIT_VIDEO) < 0 {
		return errors.New("SDL_Init")
	}
	defer C.SDL_Quit()

	// Create the window
	window = C.SDL_CreateWindow(ctitle, C.SDL_WINDOWPOS_UNDEFINED, C.SDL_WINDOWPOS_UNDEFINED, C.int(width), C.int(height), C.SDL_WINDOW_SHOWN)
	if window == nil {
		return errors.New("SDL_CreateWindow")
	}
	defer C.SDL_DestroyWindow(window)

	// Create the renderer
	renderer = C.SDL_CreateRenderer(window, C.int(-1), C.SDL_RENDERER_ACCELERATED|C.SDL_RENDERER_PRESENTVSYNC)
	if renderer == nil {
		return errors.New("SDL_CreateRenderer")
	}
	defer C.SDL_DestroyRenderer(renderer)

	EventInit()

	C.gameLoop()

	destroyTextures()

	return nil
}

func QuitGame() {
	C.quitGame()
}
