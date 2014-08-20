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
	// Letters
	K_A = C.SDLK_a
	K_B = C.SDLK_b
	K_C = C.SDLK_c
	K_D = C.SDLK_d
	K_E = C.SDLK_e
	K_F = C.SDLK_f
	K_G = C.SDLK_g
	K_H = C.SDLK_h
	K_I = C.SDLK_i
	K_J = C.SDLK_j
	K_K = C.SDLK_k
	K_L = C.SDLK_l
	K_M = C.SDLK_m
	K_N = C.SDLK_n
	K_O = C.SDLK_o
	K_P = C.SDLK_p
	K_Q = C.SDLK_q
	K_R = C.SDLK_r
	K_S = C.SDLK_s
	K_T = C.SDLK_t
	K_U = C.SDLK_u
	K_V = C.SDLK_v
	K_W = C.SDLK_w
	K_X = C.SDLK_x
	K_Y = C.SDLK_y
	K_Z = C.SDLK_z

	// Arrows
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
