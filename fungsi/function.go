package fungsi

import (
	"crypto/md5"
	"encoding/hex"
	"hash/crc32"
	"strconv"
)

func ToMd5(string string) string {
	data := []byte(string)
	b := md5.Sum(data)
	pass := hex.EncodeToString(b[:])
	return pass
}

func ToCRC32(string string) string {
	data := []byte(string)
	crc32InUint32 := crc32.ChecksumIEEE([]byte(data))
	crc32InString := strconv.FormatUint(uint64(crc32InUint32), 16)
	return crc32InString
}