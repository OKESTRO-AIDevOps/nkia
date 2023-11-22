package runtimefs

import (
	"encoding/json"
	"fmt"
	"os"
)

func MakePluginJSON() (Plugins, error) {

	ERR_MSG := "failed to load plugin json: %s"

	var plugins Plugins

	if _, err := os.Stat(".npia/install/HEAD"); err != nil {

		return plugins, fmt.Errorf(ERR_MSG, err.Error())

	}

	head_b, err := os.ReadFile(".npia/install/HEAD")

	if err != nil {
		return plugins, fmt.Errorf(ERR_MSG, err.Error())
	}

	head_str := string(head_b)

	if head_str != "control-plane" {
		return plugins, fmt.Errorf(ERR_MSG, "not the control plane")
	}

	if _, err := os.Stat(".npia/install/control-plane"); err != nil {

		return plugins, fmt.Errorf(ERR_MSG, err.Error())

	}

	if _, err := os.Stat(".npia/install/plugins.json"); err != nil {

		init_pl := []PluginJSON{

			{
				Name: "_INIT",
				Keyval: map[string]string{
					"_INIT": "_INIT",
				},
			},
		}

		jb, _ := json.Marshal(init_pl)

		_ = os.WriteFile(".npia/install/plugins.json", jb, 0644)

	}

	file_b, err := os.ReadFile(".npia/install/plugins.json")

	if err != nil {
		return plugins, fmt.Errorf(ERR_MSG, err.Error())
	}

	err = json.Unmarshal(file_b, &plugins)

	if err != nil {
		return plugins, fmt.Errorf(ERR_MSG, err.Error())
	}

	return plugins, nil

}

func GetMapFromPluginJSON(plugins Plugins, name string) (bool, map[string]string, error) {

	var ret_kv map[string]string

	var found bool = false

	for _, pj := range plugins {

		if pj.Name != name {
			continue
		} else {

			found = true

			ret_kv = pj.Keyval

		}

	}

	return found, ret_kv, nil
}

func AddMapToPluginJSON(plugins Plugins, pj PluginJSON) (Plugins, error) {

	var ret_plugins Plugins

	new_name := pj.Name

	for idx, pj := range plugins {

		if pj.Name == new_name {

			_, ret_plugins = PopPluginJSONFromSliceByIndex(plugins, idx)

		} else {
			continue
		}

	}

	ret_plugins = append(ret_plugins, pj)

	return plugins, nil
}

func SavePluginJSON(plugins Plugins) error {

	ERR_MSG := "failed to unload plugin json: %s"

	plugins_b, err := json.Marshal(plugins)

	if err != nil {
		return fmt.Errorf(ERR_MSG, err.Error())
	}

	err = os.WriteFile(".npia/install/plugins.json", plugins_b, 0644)

	if err != nil {
		return fmt.Errorf(ERR_MSG, err.Error())
	}

	return nil
}
