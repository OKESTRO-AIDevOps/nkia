package utils

import (
	"encoding/json"
	"fmt"
	"os"

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
