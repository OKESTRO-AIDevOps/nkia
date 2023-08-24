package promquery

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"

	gabs "github.com/Jeffail/gabs/v2"
)

func PromQueryPost(query string) ([]byte, error) {

	var tmp []byte

	pq_req := url.Values{
		"query": {query},
	}

	resp, err := http.Post(PROM_COMM_URL, "application/x-www-form-urlencoded", strings.NewReader(pq_req.Encode()))

	if err != nil {
		return tmp, fmt.Errorf("Error requesting prometheus: %s", err.Error())
	}

	defer resp.Body.Close()

	body_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body_bytes, nil

}

func PromQueryStandardizer(raw_query_bytes []byte, metric_key string) ([]PQOutputFormat, error) {

	var ret_batch []PQOutputFormat

	json_parsed, err := gabs.ParseJSON(raw_query_bytes)

	if err != nil {

		return ret_batch, fmt.Errorf("received query result format corrupted, abort: %s", err.Error())

	}

	data_result := json_parsed.Path("data.result.*")

	for i, metric := range data_result.Children() {

		metric_values := metric.Path("values.*")

		metric_pod := metric.Path("metric." + metric_key).String()

		numbering_i := strconv.Itoa(i)

		metric_pod = strings.ReplaceAll(metric_pod, "\"", "")

		metric_pod = strings.ReplaceAll(metric_pod, ".", "-")

		metric_pod = strings.ReplaceAll(metric_pod, ":", "-")

		id_metric_pod := numbering_i + "-" + metric_pod

		pq_record := PQOutputFormat{}

		pq_record.DataLabel = id_metric_pod

		for _, vv := range metric_values.Children() {

			unix_timestamp := vv.Path("0").String()

			parsed_timestamp, _ := strconv.ParseInt(unix_timestamp, 10, 64)

			tm := time.Unix(parsed_timestamp, 0).UTC()

			tm_str := tm.String()

			prom_value := vv.Path("1").String()

			prom_value = strings.ReplaceAll(prom_value, "\"", "")

			prom_vfloat, _ := strconv.ParseFloat(prom_value, 64)

			pq_record.Timestamp = append(pq_record.Timestamp, tm_str)

			pq_record.Values = append(pq_record.Values, prom_vfloat)

		}

		ret_batch = append(ret_batch, pq_record)

	}

	return ret_batch, nil

}

func SanitizePQ(check_q string) error {

	for _, c := range check_q {

		if !unicode.IsLower(c) || c == '-' {

			return fmt.Errorf(": %s", "sanitization failed")

		}

	}

	return nil
}
