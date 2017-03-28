package main

import (
	"fmt"
	"go_test/jump_hash/util"
)

func main() {
	ips := []string{"234.23.62.76",
		"45.12.82.44",
		"43.21.44.67",
		"123.23.1.243",
		"65.71.42.33",
		"29.33.43.87",
		"34.66.74.62",
		"153.46.32.88",
		"74.36.46.75",
		"25.55.47.82",
		"33.67.82.55"}

	for _, ip := range ips {
		ip2uint32 := MyNet.Ip2uint32(ip)
		fmt.Println(ip2uint32)

		hash_res := MyNet.JumpHash(uint64(ip2uint32), 12)
		fmt.Println(hash_res)
		fmt.Println("")
	}
}
