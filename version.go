package remo

import "fmt"

var MAJOR=0
var MINOR=3
var PATCH=0
var VERSION=fmt.Sprintf("%d.%d.%d", MAJOR, MINOR, PATCH)

func GetCurrentVersion() string {
	return VERSION
}
