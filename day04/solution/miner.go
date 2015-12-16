package solution
import "strings"

type Miner struct {
	Hash Hasher
}

func (this *Miner) Mine(key string) (string, int) {
	for i := 1; true; i++ {
		coin := this.Hash.Hash(key, i);
		if strings.HasPrefix(coin, "00000") {
			return coin, i;
		}
	}
	return "", 0
}
