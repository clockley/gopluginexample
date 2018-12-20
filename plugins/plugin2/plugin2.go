package main

import (
	"alliance/packages/shared"
)

func testH() {
	print("Hello plugin2\n")
}

func PluginInfo() *shared.PluginInfo {
	var info shared.PluginInfo
	info.ID = "Name"
	info.MethodMap = make(map[shared.Events]interface{})
	info.MethodMap[shared.SessionUserLogoff] = testH
	return &info
}
