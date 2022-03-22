package zabbix

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNewClient() {
	client, err := NewClient("127.0.0.1:10051", "localhost")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(client.Hostname)
	// Output: localhost
}

func ExampleClient_Add() {
	client, err := NewClient("127.0.0.1:10051", "localhost")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	client.Add("test1", 1)
	client.Add("test2", 2)

	retro := time.Date(2019, 5, 1, 14, 34, 11, 0, time.Local)

	m := client.Add("test3", 3)
	m.Host = "host1.domain.com"
	m.Clock = retro.Unix()
	m.NS = retro.Nanosecond()

	client.Send()
}

func ExampleClient_Num() {
	client, err := NewClient("127.0.0.1:10051", "localhost")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	client.Add("test1", 1)
	client.Add("test2", 2)
	client.Add("test3", 3)

	fmt.Printf("Items in stack: %d", client.Num())
	// Output: Items in stack: 3
}

func ExampleClient_Clear() {
	client, err := NewClient("127.0.0.1:10051", "localhost")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	client.Add("test1", 1)
	client.Add("test2", 2)
	client.Add("test3", 3)

	client.Clear()

	fmt.Printf("Items in stack: %d", client.Num())
	// Output: Items in stack: 0
}

func ExampleClient_Send() {
	client, err := NewClient("127.0.0.1:10051", "localhost")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	client.Add("test1", 1)
	client.Add("test2", 2)
	client.Add("test3", 3)

	resp, err := client.Send()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf(
		"Metrics sended (processed: %d | failed: %d | total: %d)",
		resp.Processed, resp.Failed, resp.Total,
	)
}
