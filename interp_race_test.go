// +build race

package interp_test

func init() {
	gorootTestSkips["atomicload.go"] = "unsupport race"
	gorootTestSkips["deferfin.go"] = "unsupport race"
	gorootTestSkips["mallocfin.go"] = "unsupport race"
	gorootTestSkips["peano.go"] = "stack overflow on race"
	gorootTestSkips["chan/goroutines.go"] = "race: limit on 8128 simultaneously alive goroutines is exceeded, dying"
}
