package screen
var ScreenX = 680 //Public
var ScreenY = 720 //Public
var width = 1080

const ColorDepth = 16 //Public
const velocity = 8

func ResizeScreen(x, y int) { //Public
	ScreenX, ScreenY = x, y
}

func changeWidth(w int) {
	width = w
}
