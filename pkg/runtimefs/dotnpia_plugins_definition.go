package runtimefs

type Plugins []PluginJSON

type PluginJSON struct {
	Name string `json:"name"`

	Keyval map[string]string `json:"keyval"`
}

func PopPluginJSONFromSliceByIndex(slice []PluginJSON, idx int) (PluginJSON, []PluginJSON) {

	pop_val := slice[idx]

	return pop_val, append(slice[:idx], slice[idx+1:]...)

}
