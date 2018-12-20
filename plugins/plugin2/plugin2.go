package main

import (
	"alliance/packages/shared"
)

func testH() {
	print("Hello plugin2\n")
}

func pluginInit(argument1 interface{}, argument2 interface{}) {
	println("Plugin2 Init function")
}

func PluginInfo() *shared.PluginInfo {
	var info shared.PluginInfo
	info.ID = "Name"
	info.MethodMap = make(map[shared.Events]interface{})
	info.MethodMap[shared.SessionUserLogoff] = testH
	info.Init = pluginInit
	return &info
}
