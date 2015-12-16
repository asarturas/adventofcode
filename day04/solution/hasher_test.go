package solution

import (
	"testing"
)

var hasher = Md5Hasher{};

func Test_it_generates_hash_of_a_key_and_a_number(t *testing.T) {
	if hasher.Hash("a", 1) != "8a8bb7cd343aa2ad99b7d762030857a2" {
		t.Error("Unexpected hash result: ", hasher.Hash("a", 1));
	}
}
