package php

import "hash/crc32"

func Crc32(s string) uint32 {
	table := crc32.MakeTable(crc32.IEEE)
	return crc32.Checksum([]byte(s), table)
}
