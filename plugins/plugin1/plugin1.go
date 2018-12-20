package main

import (
	"alliance/packages/shared"
)

func testH() {
	print("Hello plugin1\n")
}

func PluginInfo() (*shared.PluginInfo, error) {
	var info shared.PluginInfo
	info.ID = "Name"
	info.MethodMap = make(map[shared.Events]interface{})
	info.MethodMap[shared.SessionUserLogin] = testH
	return &info, nil
}
