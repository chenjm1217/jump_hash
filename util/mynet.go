package MyNet

import (
	//"fmt"
	"strconv"
	"strings"
)

func Ip2uint32(ip string) uint32 {
	bits := strings.Split(ip, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return uint32(sum)
}

/*
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
	for _, v := range ips {
		fmt.Println(Ip2uint32(v))
	}
}
*/

/////////////////////////////////////
func JumpHash(key uint64, bucket int) int64 {
	b := int64(-1)
	j := int64(0)

	for {
		if j >= int64(bucket) {
			break
		}
		b = j
		key = key*uint64(2862933555777941757) + 1
		j = int64(float64(b+1) * (float64(1<<31) / float64((key>>33)+1)))
	}
	return b
}
