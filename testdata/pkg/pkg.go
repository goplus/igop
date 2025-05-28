package pkg

import "fmt"

func Add(i, j int) int {
	return i + j
}

func Println(a ...interface{}) {
	fmt.Println(a...)
}

type App struct {
	init bool
}

func (p *App) initApp() {
	p.init = true
}

func (p *App) IsInit() bool {
	println("initApp")
	return p.init
}

func RunApp(a interface{ initApp() }) {
	a.initApp()
	a.(interface{ MainEntry() }).MainEntry()
}
