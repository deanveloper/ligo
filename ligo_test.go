package ligo_test

import (
	"encoding/base32"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"testing/quick"

	"github.com/deanveloper/ligo"
)

type Input struct {
	Website string
	CodeLen int
}

func (Input) Generate(r *rand.Rand, size int) reflect.Value {
	buf := make([]byte, size)
	r.Read(buf)
	b32 := base32.StdEncoding.EncodeToString(buf)

	return reflect.ValueOf(Input{
		Website: b32,
		CodeLen: r.Intn(size * 2),
	})
}

func TestWebsiteToCode_table(t *testing.T) {

	tests := []struct {
		website string
		code    string
	}{
		{
			website: "https://google.com",
			code:    "IllIlIIIIlllIlIIIlllIlIIIlllIIIIIlllIIllIIlllIlIIIlIllllIIlIllllIllIIlllIllIllllIllIllllIllIIlllIllIllIIIllIIlIlIIlIlllIIllIIIllIllIllllIllIllIl",
		},
		{
			website: "https://mozilla.org",
			code:    "IllIlIIIIlllIlIIIlllIlIIIlllIIIIIlllIIllIIlllIlIIIlIllllIIlIllllIllIllIlIllIllllIllllIlIIllIlIIlIllIllIIIllIllIIIllIIIIlIIlIlllIIllIllllIlllIIlIIllIIlll",
		},
	}

	for _, test := range tests {
		out := ligo.WebsiteToCode(test.website, 0)
		if test.code != out {
			t.Errorf("Test failed for %v.\nExpected:\n\t%v\ngot\n\t%v", test.website, test.code, out)
		}
	}

}

func TestWebsiteToCode_alwaysGreaterOrEqual(t *testing.T) {
	valid := func(in Input) bool {
		code := ligo.WebsiteToCode(in.Website, in.CodeLen)
		return len(code) >= in.CodeLen
	}
	err := quick.Check(valid, &quick.Config{})
	if err != nil {
		t.Error(err)
	}
}

func TestWebsiteToCode_alwaysSamePrefix(t *testing.T) {
	removedPadding := func(in Input) string {
		code := ligo.WebsiteToCode(in.Website, in.CodeLen)
		zeroIndex := strings.Index(code, "IIIIIIII")
		if zeroIndex >= 0 {
			return code[:zeroIndex]
		}
		return code
	}
	withoutPadding := func(in Input) string {
		return ligo.WebsiteToCode(in.Website, 0)
	}
	err := quick.CheckEqual(removedPadding, withoutPadding, nil)
	if err != nil {
		t.Error(err)
	}
}
