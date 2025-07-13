package animation

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAnimationPulse(t *testing.T) {
	got := AnimationPulse()
	b, _ := json.MarshalIndent(got, "", " ")
	fmt.Println(string(b))
}
