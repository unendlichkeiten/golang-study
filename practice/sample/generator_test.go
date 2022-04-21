package sample

import (
	"bytes"
	"encoding/json"
	"testing"

	_ "practice/pb"
)

func TestNewLaptop(t *testing.T) {
	laptop := NewLaptop()
	laptopStr, _ := json.Marshal(laptop)

	var out bytes.Buffer
	json.Indent(&out, laptopStr, "	", "\t")
	t.Log(out.String())
}
