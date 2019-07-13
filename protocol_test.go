package zabbix

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"encoding/binary"
	"encoding/json"
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type request struct {
	Request string    `json:"request"`
	Session string    `json:"session"`
	Clock   int64     `json:"clock"`
	NS      int       `json:"ns"`
	Data    []*metric `json:"data"`
}

type metric struct {
	Host  string `json:"host"`
	Key   string `json:"key"`
	Value string `json:"value"`
	ID    int    `json:"id"`
	Clock int64  `json:"clock"`
	NS    int    `json:"ns"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

type ZabbixSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ZabbixSuite{})

var exampleResponse = []byte(`{
    "response":"success",
    "info":"processed: 6; failed: 2; total: 8; seconds spent: 0.003571"
}`)

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ZabbixSuite) TestClient(c *C) {
	client, err := NewClient("127.0.", "localhost")

	c.Assert(client, IsNil)
	c.Assert(err, NotNil)

	client, err = NewClient("127.0.0.0:10051", "localhost")

	c.Assert(client, NotNil)
	c.Assert(err, IsNil)

	client.Add("test1", 8381794)
	client.Add("test2", 7.391348924)

	c.Assert(client.Num(), Equals, 2)

	client.Clear()

	c.Assert(client.Num(), Equals, 0)
	c.Assert(client.data, IsNil)

	resp, err := client.Send()

	c.Assert(resp.Status, Equals, "nothing to send")
}

func (s *ZabbixSuite) TestEncoder(c *C) {
	client, _ := NewClient("127.0.0.0:10051", "localhost")

	m1 := client.Add("test1", 8381794)
	m2 := client.Add("test2", 7.391348924)
	m3 := client.Add("test3", "ABCD")
	m4 := client.Add("test4", float32(3.11836103))

	c.Assert(m1.Value, Equals, "8381794")
	c.Assert(m2.Value, Equals, "7.391349")
	c.Assert(m3.Value, Equals, "ABCD")
	c.Assert(m4.Value, Equals, "3.118361")

	c.Assert(client.data, HasLen, 4)

	payload := encodeMetrics(client.data)

	c.Assert(payload[:5], DeepEquals, zabbixHeader)

	payloadSize := binary.LittleEndian.Uint64(payload[5:13])

	c.Assert(payloadSize, Equals, uint64(487))

	req := &request{}
	err := json.Unmarshal(payload[13:], req)

	c.Assert(err, IsNil)
	c.Assert(req.Request, Equals, "agent data")
	c.Assert(req.Data[0].Value, Equals, "8381794")
}

func (s *ZabbixSuite) TestResponseDecoder(c *C) {
	resp, err := decodeResponse(exampleResponse)

	c.Assert(err, IsNil)
	c.Assert(resp.Status, Equals, "success")
	c.Assert(resp.Processed, Equals, 6)
	c.Assert(resp.Failed, Equals, 2)
	c.Assert(resp.Total, Equals, 8)
	c.Assert(resp.SecondsSpent, Equals, 0.003571)

	_, err = decodeResponse([]byte(`{EXAMPLE}`))

	c.Assert(err, NotNil)

	_, err = decodeResponse([]byte(`{"":"", ""}`))

	c.Assert(err, NotNil)

	_, err = decodeResponse([]byte(`{"":"", "":"abcd"}`))

	c.Assert(err, NotNil)

	_, err = decodeResponse([]byte(`":", 0`))

	c.Assert(err, NotNil)

	_, err = decodeResponse([]byte(`":"0", ":"`))

	c.Assert(err, NotNil)

	_, _, _, _, err = parseResponseInfo("processed: V; failed: 2; total: 8; seconds spent: 0.003571")

	c.Assert(err, NotNil)

	_, _, _, _, err = parseResponseInfo("processed: 6; failed: V; total: 8; seconds spent: 0.003571")

	c.Assert(err, NotNil)

	_, _, _, _, err = parseResponseInfo("processed: 6; failed: 2; total: V; seconds spent: 0.003571")

	c.Assert(err, NotNil)

	_, _, _, _, err = parseResponseInfo("processed: 6; failed: 2; total: 8; seconds spent: V")

	c.Assert(err, NotNil)
}
