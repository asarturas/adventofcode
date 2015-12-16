package solution
import "strings"

type Miner struct {
	CoinPrefix string
	Hash Hasher
}

func (this *Miner) Mine(key string) (string, int) {
	for i := 1; true; i++ {
		coin := this.Hash.Hash(key, i);
		if strings.HasPrefix(coin, this.CoinPrefix) {
			return coin, i;
		}
	}
	return "", 0
}
