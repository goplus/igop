package pkg

import (
	"fmt"
)

func Add(i, j int) int {
	return i + j
}

func Println(a ...interface{}) {
	fmt.Println(a...)
}

type App struct {
	init bool
	test bool
}

func (p *App) initApp() {
	p.init = true
}

func (p *App) initTest(b bool) {
	p.test = b
}

func (p *App) IsInit() bool {
	return p.init
}

func (p *App) IsTest() bool {
	return p.test
}

func RunApp(a interface{ initApp() }) {
	a.initApp()
	a.(interface{ MainEntry() }).MainEntry()
}

func RunTest(a interface{ initTest(b bool) }) {
	a.initTest(true)
	a.(interface{ MainEntry() }).MainEntry()
}
