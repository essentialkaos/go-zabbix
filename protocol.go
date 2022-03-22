package zabbix

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Response response is parsed Zabbix server response data
type Response struct {
	Status       string
	Processed    int
	Failed       int
	Total        int
	SecondsSpent float64
}

// Metric contains information about one metric
type Metric struct {
	id int

	Clock int64
	NS    int
	Host  string
	Key   string
	Value string
}

// ////////////////////////////////////////////////////////////////////////////////// //

// symbols is slice with symbols used for session ID generation
var symbols = "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890"

// zabbixHeader is Zabbix header
var zabbixHeader = []byte("ZBXD\x01")

// ////////////////////////////////////////////////////////////////////////////////// //

// encodeMetrics encodes metrics to zabbix protocol payload
// https://www.zabbix.com/documentation/4.2/manual/appendix/protocols/header_datalen
func encodeMetrics(metrics []*Metric) []byte {
	payload := marshalMetrics(metrics)
	return encodePayload(payload)
}

// marshalMetrics marshals data to JSON object
func marshalMetrics(metrics []*Metric) []byte {
	var buf bytes.Buffer

	now := time.Now()
	totalItems := len(metrics)

	buf.WriteString("{")
	buf.WriteString(`"request":"sender data",`)
	buf.WriteString(fmt.Sprintf(`"session":"%s",`, genSessionID()))
	buf.WriteString(fmt.Sprintf(`"clock":%d,`, now.Unix()))
	buf.WriteString(fmt.Sprintf(`"ns":%d,`, now.Nanosecond()))

	buf.WriteString(`"data":[`)

	for index, metric := range metrics {
		marshalMetric(&buf, metric)

		if index+1 < totalItems {
			buf.WriteRune(',')
		}
	}

	buf.WriteString("]}")

	return buf.Bytes()
}

// marshalMetric marshal metric data to JSON object
func marshalMetric(buf *bytes.Buffer, metric *Metric) {
	buf.WriteString("{")
	buf.WriteString(fmt.Sprintf(`"host":%s,`, strconv.Quote(metric.Host)))
	buf.WriteString(fmt.Sprintf(`"key":%s,`, strconv.Quote(metric.Key)))
	buf.WriteString(fmt.Sprintf(`"value":%s,`, strconv.Quote(metric.Value)))
	buf.WriteString(fmt.Sprintf(`"id":%d,`, metric.id))
	buf.WriteString(fmt.Sprintf(`"clock":%d,`, metric.Clock))
	buf.WriteString(fmt.Sprintf(`"ns":%d`, metric.NS))
	buf.WriteString("}")
}

// genSessionID generates unique session ID
func genSessionID() string {
	result := make([]byte, 32)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 32; i++ {
		result[i] = symbols[rand.Intn(62)]
	}

	return string(result)
}

// encodePayload encodes payload
func encodePayload(payload []byte) []byte {
	size := uint64(len(payload))

	var buf bytes.Buffer

	sizeBuf := make([]byte, 8)
	binary.LittleEndian.PutUint64(sizeBuf, size)

	buf.Write(zabbixHeader)
	buf.Write(sizeBuf)
	buf.Write(payload)

	return buf.Bytes()
}

// decodeResponse decodes Zabbix server response
func decodeResponse(resp []byte) (Response, error) {
	data := strings.Trim(string(resp), "{}")

	rs := strings.Index(data, `":"`)
	re := strings.Index(data, `","`)

	if rs == -1 || re == -1 || rs+3 >= re {
		return Response{}, fmt.Errorf("Can't decode response status")
	}

	status := data[rs+3 : re]
	data = data[re+3:]

	rs = strings.Index(data, `":"`)

	if rs == -1 || rs >= len(data)-4 {
		return Response{}, fmt.Errorf("Can't decode response info")
	}

	info := data[rs+3 : len(data)-1]

	processed, failed, total, spent, err := parseResponseInfo(info)

	if err != nil {
		return Response{}, fmt.Errorf("Can't decode response info: %v", err)
	}

	return Response{
		Status:       status,
		Processed:    processed,
		Failed:       failed,
		Total:        total,
		SecondsSpent: spent,
	}, nil
}

// parseResponseInfo parses response processing info
func parseResponseInfo(data string) (int, int, int, float64, error) {
	items := strings.Fields(data)

	if len(items) != 9 {
		return -1, -1, -1, 0.0, fmt.Errorf("Wrong number of items")
	}

	processed, err := strconv.Atoi(strings.Trim(items[1], ";"))

	if err != nil {
		return -1, -1, -1, 0.0, fmt.Errorf("Can't parse processed value: %v", err)
	}

	failed, err := strconv.Atoi(strings.Trim(items[3], ";"))

	if err != nil {
		return -1, -1, -1, 0.0, fmt.Errorf("Can't parse failed value: %v", err)
	}

	total, err := strconv.Atoi(strings.Trim(items[5], ";"))

	if err != nil {
		return -1, -1, -1, 0.0, fmt.Errorf("Can't parse total value: %v", err)
	}

	spent, err := strconv.ParseFloat(items[8], 64)

	if err != nil {
		return -1, -1, -1, 0.0, fmt.Errorf("Can't parse spent value: %v", err)
	}

	return processed, failed, total, spent, nil
}
