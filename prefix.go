package common

import "reflect"

var PREFIX = reflect.TypeOf(prefix).PkgPath()

func prefix() {}
