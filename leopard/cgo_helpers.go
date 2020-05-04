// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Mon, 04 May 2020 02:25:44 CEST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package leopard

/*
#cgo LDFLAGS: -L./ -llibleopard -lstdc++
#include "leopard.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = unsafe.Pointer(p)
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - uintptr(h.Data))
	}
	return
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(h.Data), C.int(h.Len))
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// allocPUint8TMemory allocates memory for type *C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPUint8TMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPUint8TValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPUint8TValue = unsafe.Sizeof([1]*C.uint8_t{})

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackArgSSByte transforms a sliced Go data structure into plain C format.
func unpackArgSSByte(x [][]byte) (unpacked **C.uint8_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.uint8_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPUint8TMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: mem0,
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.uint8_t)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.uint8_t)(h.Data)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.uint8_t)(h.Data)
	return
}

// packSSByte reads sliced Go data structure out from plain C format.
func packSSByte(v [][]byte, ptr0 **C.uint8_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.uint8_t)(unsafe.Pointer(ptr0)))[i0]
		hxfc4425b := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxfc4425b.Data = unsafe.Pointer(ptr1)
		hxfc4425b.Cap = 0x7fffffff
		// hxfc4425b.Len = ?
	}
}
