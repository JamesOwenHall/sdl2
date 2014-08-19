package sdl2

// #include <SDL2/SDL.h>
import "C"

type KeyCode C.int

// Events
var (
	EventInit    = func() {}
	EventQuit    = func() {}
	EventDraw    = func() {}
	EventUpdate  = func() {}
	EventKeyDown = func(KeyCode) {}
	EventKeyUp   = func(KeyCode) {}
)

// KeyCodes
const (
	K_Right = C.SDLK_RIGHT
	K_Left  = C.SDLK_LEFT
	K_Down  = C.SDLK_DOWN
	K_Up    = C.SDLK_UP
)

//export eventQuit
func eventQuit() {
	EventQuit()
}

//export eventDraw
func eventDraw() {
	EventUpdate()
	EventDraw()
	C.SDL_RenderPresent(renderer)
}

//export eventKeyDown
func eventKeyDown(key C.int) {
	EventKeyDown(KeyCode(key))
}

//export eventKeyUp
func eventKeyUp(key C.int) {
	EventKeyUp(KeyCode(key))
}
