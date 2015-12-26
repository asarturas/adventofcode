package solution

import (
	"testing"
)

func Test_it_is_off_with_zero_brightness(t *testing.T) {
	light := Light{}
	if light.IsOn() {
		t.Error("Expected light not to be on. Current brightness", light.Brightness);
	}
	if !light.IsOff() {
		t.Error("Expected light to be off. Current brightness", light.Brightness);
	}
}
