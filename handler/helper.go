package handler

import "github.com/gopherjs/gopherjs/js"

var localStorage *js.Object

func NewStorage(newStorage *js.Object) {
	localStorage = newStorage
}

// Save val into localStorage under key
func SetItem(key string, val string) {
	localStorage.Call("setItem", key, val)
}

// Return "" when no key found
func GetItem(key string) string {
	obj := localStorage.Call("getItem", key)
	if obj == nil {
		return ""
	}
	return obj.String()
}
