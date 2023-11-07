package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/promquery"

	agph "github.com/guptarohit/asciigraph"
)

func CheckIfSliceContains[T comparable](slice []T, ele T) bool {

	hit := false

	for i := 0; i < len(slice); i++ {

		if slice[i] == ele {

			hit = true

			return hit
		}

	}

	return hit

}

func RenderASCIIGraph(render_target string) {

	var render_data promquery.PQOutputFormat

	file_byte, _ := os.ReadFile(render_target)

	json.Unmarshal(file_byte, &render_data)

	out := agph.Plot(render_data.Values)

	fmt.Println(out)

}

func PopFromSliceByIndex[T comparable](slice []T, idx int) (T, []T) {

	pop_val := slice[idx]

	return pop_val, append(slice[:idx], slice[idx+1:]...)

}

func InsertToSliceByIndex[T comparable](slice []T, idx int, val T) []T {

	return append(slice[:idx], append([]T{val}, slice[idx:]...)...)
}

func SplitStrict(content string) map[string]string {
	out := map[string]string{}
	for _, line := range strings.Split(content, "\n") {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := parts[0]
		if len(key) == 0 || key[0] == '#' {
			continue
		}

		value := parts[1]
		if len(value) > 2 && value[0] == '"' && value[len(value)-1] == '"' {

			var err error
			value, err = strconv.Unquote(value)
			if err != nil {
				continue
			}
		}
		out[key] = value
	}
	return out
}

func MakeOSReleaseLinux() map[string]string {

	var osRelease map[string]string

	if osRelease == nil {

		osRelease = map[string]string{}
		if bytes, err := os.ReadFile("/etc/os-release"); err == nil {
			osRelease = SplitStrict(string(bytes))
		}
	}
	return osRelease
}

func SliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
