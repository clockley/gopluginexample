package shared

type Events string

const (
	SessionUserLogin  = "SessionUserLogin"
	SessionUserLogoff = "SessionUserLogoff"
)

type PluginInfo struct {
	ID           string                         // machine-name or ID of the plugin (must be unique)
	Version      string                         // plugin version (SemVer)
	Name         string                         // human-readble name of the plugin
	Description  string                         // describes the purpose of the plugin and functionality
	MethodMap    map[Events]interface{}         //Hashtable of methods and events
	Dependencies *PluginInfo                    // Plugin infos of plugins this one depends on (requires)
	Init         func(interface{}, interface{}) //Pointer to init function
}
