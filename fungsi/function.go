package fungsi

import (
	"crypto/md5"
	"encoding/hex"
	"hash/crc32"
	"strconv"
	"strings"
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

func CategoryParseEngine(Klasifikasi string) string {
	classifiers := strings.Split(Klasifikasi, ",")
	var Category string
	for x, class := range classifiers {
		c, _ := strconv.ParseFloat(class, 32)
		if c == 4 {
			if x == 0 {
				Category = Category + "Data processing, computer science"
			} else {
				Category = Category + ", Data processing, computer science"
			}
		} else if c >= 200 && c < 300 {
			if x == 0 {
				Category = Category + "Religion"
			} else {
				Category = Category + ", Religion"
			}
		} else if c >= 300 && c < 400 {
			if x == 0 {
				Category = Category + "Social science"
			} else {
				Category = Category + ", Social science"
			}
		} else if c >= 400 && c < 500 {
			if x == 0 {
				Category = Category + "Language"
			} else {
				Category = Category + ", Language"
			}
		} else if c >= 500 && c < 600 {
			if x == 0 {
				Category = Category + "Natural science"
			} else {
				Category = Category + ", Natural science"
			}
		} else if c >= 620 && c < 621 {
			if x == 0 {
				Category = Category + "Engineering & allied operation"
			} else {
				Category = Category + ", Engineering & allied operation"
			}
		} else if c >= 621.3 && c < 621.4 {
			if x == 0 {
				Category = Category + "Electric, electronic, magnetic, communications"
			} else {
				Category = Category + ", Electric, electronic, magnetic, communications"
			}
		} else if c >= 623 && c < 624 {
			if x == 0 {
				Category = Category + "Ship"
			} else {
				Category = Category + ", Ship"
			}
		} else if c >= 604 && c < 605 {
			if x == 0 {
				Category = Category + "Technical drawing"
			} else {
				Category = Category + ", Technical drawing"
			}
		} else if c >= 621.4 && c < 621.9 {
			if x == 0 {
				Category = Category + "Heat engineering, Pump, Pneumatic, Machine engineering"
			} else {
				Category = Category + ", Heat engineering, Pump, Pneumatic, Machine engineering"
			}
		} else if c >= 658 && c < 659 {
			if x == 0 {
				Category = Category + "General Management"
			} else {
				Category = Category + ", General Management"
			}
		} else if c >= 629 && c < 629.9 {
			if x == 0 {
				Category = Category + "Other branches of engineering; automatic control engineering"
			} else {
				Category = Category + ", Other branches of engineering; automatic control engineering"
			}
		} else if c >= 669 && c < 672 {
			if x == 0 {
				Category = Category + "Metallurgy, manufacturing, metalworking processes"
			} else {
				Category = Category + ", Metallurgy, manufacturing, metalworking processes"
			}
		} else if c >= 700 && c < 701 {
			if x == 0 {
				Category = Category + "The Arts"
			} else {
				Category = Category + ", The Arts"
			}
		} else if c >= 800 && c < 1000 {
			if x == 0 {
				Category = Category + "Literature, history and geography"
			} else {
				Category = Category + ", Literature, history and geography"
			}
		} else {
			if x == 0 {
				Category = Category + "Unknown Category"
			} else {
				Category = Category + ", Unknown Category"
			}
		}
	}
	return Category
}
