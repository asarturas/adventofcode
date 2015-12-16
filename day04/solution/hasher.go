package solution
import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

type Hasher interface {
	Hash(key string, number int) string;
}

type Md5Hasher struct {}

func (this Md5Hasher) Hash(key string, number int) string {
	hash := md5.Sum([]byte(key + strconv.Itoa(number)));
	return hex.EncodeToString(hash[:]);
}
