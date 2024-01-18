package routineid

import (
	"fmt"
	"reflect"
	"unsafe"
)

var (
	offsetGoid       uintptr
	offsetParentGoid uintptr
)

func init() {
	gt := getgt()
	offsetGoid = offset(gt, "goid")
	offsetParentGoid = offset(gt, "parentGoid")
}

type g struct {
	goid       uint64
	parentGoid uint64
}

// getg returns current coroutine struct.
func getg() g {
	gp := getgp()
	if gp == nil {
		panic("Failed to get gp from runtime natively.")
	}
	return g{
		goid:       *(*uint64)(add(gp, offsetGoid)),
		parentGoid: *(*uint64)(add(gp, offsetParentGoid)),
	}
}

// offset returns the offset of the specified field.
func offset(t reflect.Type, f string) uintptr {
	field, found := t.FieldByName(f)
	if found {
		return field.Offset
	}
	panic(fmt.Sprintf("No such field '%v' of struct '%v.%v'.", f, t.PkgPath(), t.Name()))
}

// add pointer addition operation.
func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

func GetRoutineIds() (curRoutineId uint64, parentRoutineId uint64) {
	runtime_g := getg()
	curRoutineId = runtime_g.goid
	parentRoutineId = runtime_g.parentGoid
	return
}
