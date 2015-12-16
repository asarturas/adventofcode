package solution

import (
	"testing"
	"github.com/golang/mock/gomock"
)

func Test_it_returns_mined_coin_hash_if_hash_starts_with_five_zeroes(t *testing.T) {
	controller := gomock.NewController(t);
	hasher := NewMockHasher(controller);
	hasher.EXPECT().Hash("a", 1).Times(1).Return("00000a");
	miner := &Miner{"00000", hasher};
	coin, number := miner.Mine("a");
	if number != 1 {
		t.Errorf("Didn't return expected coin:", coin);
	}
}

func Test_it_returns_first_mined_coin(t *testing.T) {
	controller := gomock.NewController(t);
	hasher := NewMockHasher(controller);
	hasher.EXPECT().Hash("a", 1).Times(1).Return("10000a");
	hasher.EXPECT().Hash("a", 2).Times(1).Return("20000a");
	hasher.EXPECT().Hash("a", 3).Times(1).Return("00100b");
	hasher.EXPECT().Hash("a", 4).Times(1).Return("30000a");
	miner := &Miner{"00100", hasher}
	coin, number := miner.Mine("a");
	if number != 3 {
		t.Errorf("Didn't return expected coin:", coin);
	}
}
