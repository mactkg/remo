package remo

import "fmt"

var MAJOR=0
var MINOR=2
var PATCH=1
var VERSION=fmt.Sprintf("%d.%d.%d", MAJOR, MINOR, PATCH)

func GetCurrentVersion() string {
	return VERSION
}
