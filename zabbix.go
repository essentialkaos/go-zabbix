// Package zabbix provides client for sending metrics to Zabbix Server 3+
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
	"io"
	"net"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Client is Zabbix client
type Client struct {
	Hostname string

	ConnectTimeout time.Duration
	WriteTimeout   time.Duration
	ReadTimeout    time.Duration

	dialer *net.Dialer
	addr   *net.TCPAddr
	data   []*Metric
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewClient creates new client
func NewClient(address, hostname string) (*Client, error) {
	addr, err := net.ResolveTCPAddr("tcp4", address)

	if err != nil {
		return nil, err
	}

	dialer := &net.Dialer{Timeout: time.Second * 5}

	return &Client{addr: addr, dialer: dialer, Hostname: hostname}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Add adds new metric to stack
func (c *Client) Add(key string, value interface{}) *Metric {
	now := time.Now()
	metric := &Metric{
		id: len(c.data),

		Clock: now.Unix(),
		NS:    now.Nanosecond(),
		Host:  c.Hostname,
		Key:   key,
		Value: formatValue(value),
	}

	c.data = append(c.data, metric)

	return metric
}

// Num returns number of metrics in stack
func (c *Client) Num() int {
	return len(c.data)
}

// Clear clears all metrics in stack
func (c *Client) Clear() {
	c.data = nil
}

// Send sends data to Zabbix server
func (c *Client) Send() (Response, error) {
	if len(c.data) == 0 {
		return Response{"nothing to send", 0, 0, 0, 0.0}, nil
	}

	conn, err := connectToServer(c)

	if err != nil {
		return Response{}, err
	}

	defer conn.Close() // Zabbix doesn't support persistent connections

	err = writeToConnection(conn, encodeMetrics(c.data), c.WriteTimeout)

	if err != nil {
		return Response{}, err
	}

	c.data = nil

	respMeta := make([]byte, 13)
	err = readFromConnection(conn, respMeta, c.ReadTimeout)

	if err != nil {
		return Response{}, err
	}

	if !bytes.Equal(respMeta[:5], zabbixHeader) {
		return Response{}, fmt.Errorf("Wrong header format")
	}

	respSize := binary.LittleEndian.Uint64(respMeta[5:])
	respBuf := make([]byte, respSize)

	err = readFromConnection(conn, respBuf, c.ReadTimeout)

	if err != nil {
		return Response{}, err
	}

	return decodeResponse(respBuf)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// connectToServer makes connetion to Zabbix server
func connectToServer(c *Client) (*net.TCPConn, error) {
	if c.ConnectTimeout > 0 && c.dialer.Timeout != c.ConnectTimeout {
		c.dialer.Timeout = c.ConnectTimeout
	}

	conn, err := c.dialer.Dial(c.addr.Network(), c.addr.String())

	if err != nil {
		return nil, err
	}

	return conn.(*net.TCPConn), nil
}

// readFromConnection reads data fron connection
func readFromConnection(conn *net.TCPConn, buf []byte, timeout time.Duration) error {
	if timeout > 0 {
		conn.SetReadDeadline(time.Now().Add(timeout))
	}

	_, err := io.ReadFull(conn, buf)

	return err
}

// writeToConnection writes data into connection
func writeToConnection(conn *net.TCPConn, data []byte, timeout time.Duration) error {
	if timeout > 0 {
		conn.SetWriteDeadline(time.Now().Add(timeout))
	}

	_, err := conn.Write(data)

	return err
}

// formatValue convert value to string
func formatValue(v interface{}) string {
	switch t := v.(type) {
	case float32:
		return fmt.Sprintf("%.6f", t)
	case float64:
		return fmt.Sprintf("%.6f", t)
	default:
		return fmt.Sprint(t)
	}
}
