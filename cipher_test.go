package cipher

import "testing"

const testString1 = "Hello, World!"
const testString2 = "Goodbye, World!"
const testString3 = "Has this test gone foobar, or fizzbuzz?"
const testString4 = "3l173 5up3r h4ck3r!!!11!"

func TestAtbash(t *testing.T) {
	answerMap := make(map[string]string)
	answerMap[testString1] = "Svool, Dliow!"
	answerMap[testString2] = "Tllwybv, Dliow!"
	answerMap[testString3] = "Szh gsrh gvhg tlmv ullyzi, li uraayfaa?"
	answerMap[testString4] = "3o173 5fk3i s4xp3i!!!11!"

	for k, v := range answerMap {
		tval := Atbash(k)

		if tval != v {
			t.Errorf("expected Atbash(\"%v\") == \"%v\" but got \"%v\"", k, v, tval)
		}
	}
}
