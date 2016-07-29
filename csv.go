package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/structs"
)

func getHeader(data interface{}) string {
	return strings.Join(structs.New(data).Names(), ",")
}

func getValues(data interface{}) string {
	var values []string

	for _, v := range structs.New(data).Values() {

		switch t := v.(type) {
		case string:
			values = append(values, t)
		case bool:
			values = append(values, func() string {
				if t {
					return "1"
				}
				return "0"
			}())
		case int:
			values = append(values, strconv.Itoa(t))
		case int8:
			values = append(values, strconv.FormatInt(int64(t), 10))
		case int16:
			values = append(values, strconv.FormatInt(int64(t), 10))
		case int32:
			values = append(values, strconv.FormatInt(int64(t), 10))
		case int64:
			values = append(values, strconv.FormatInt(int64(t), 10))
		case float32:
			values = append(values, fmt.Sprintf("%.10f", t))
		case float64:
			values = append(values, fmt.Sprintf("%.10f", t))
		case nil:
			values = append(values, "")
		default:
			values = append(values, "???")
		}
	}

	return strings.Join(values, ",")
}

func jsonToCsv(outputFile string, data interface{}) {
	var _ *bufio.Reader
	var _ *csv.Writer
}