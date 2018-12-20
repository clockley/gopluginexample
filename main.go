package main

import (
	"alliance/packages/shared"
	"log"
	"os"
	"plugin"
)

func main() {
	p, _ := plugin.Open(os.Args[1])
	symName, err := p.Lookup("PluginInfo")
	if err != nil {
		log.Fatal(err)
	}
	var info *shared.PluginInfo

	var infoFunc = symName.(func() *shared.PluginInfo)

	info = infoFunc()
	if info.Init != nil {
		info.Init(nil, nil)
	}

	if info.MethodMap != nil {
		if info.MethodMap[shared.SessionUserLogin] != nil {
			f := info.MethodMap[shared.SessionUserLogin].(func())
			f()
		}
		if info.MethodMap[shared.SessionUserLogoff] != nil {
			f := info.MethodMap[shared.SessionUserLogoff].(func())
			f()
		}
	}
	print(info.ID + "\n")
}
