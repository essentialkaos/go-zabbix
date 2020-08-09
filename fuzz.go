// +build gofuzz

package zabbix

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

func Fuzz(data []byte) int {
	_, err := decodeResponse(data)

	if err != nil {
		return 0
	}

	return 1
}
