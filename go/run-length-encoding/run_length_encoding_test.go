package encode

import (
	"testing"
)

func TestRunLengthEncode(t *testing.T) {
	for _, test := range encodeTests {
		if actual := RunLengthEncode(test.input); actual != test.expected {
			t.Errorf("FAIL %s - RunLengthEncode(%s) = %q, expected %q.",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS RunLengthEncode - %s", test.description)
	}
}

func TestRunLengthDecode(t *testing.T) {
	for _, test := range decodeTests {
		if actual := RunLengthDecode(test.input); actual != test.expected {
			t.Errorf("FAIL %s - RunLengthDecode(%s) = %q, expected %q.",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS RunLengthDecode - %s", test.description)
	}
}

func TestRunLengthEncodeDecode(t *testing.T) {
	for _, test := range encodeDecodeTests {
		if actual := RunLengthDecode(RunLengthEncode(test.input)); actual != test.expected {
			t.Errorf("FAIL %s - RunLengthDecode(RunLengthEncode(%s)) = %q, expected %q.",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS %s", test.description)
	}
}

func TestEncode(t *testing.T) {

	if Encode('X', 0) != "" {
		t.Errorf("should be emtpy")
	}
	if Encode('X', 1) != "X" {
		t.Errorf("should be X")
	}
	if Encode('X', 3) != "3X" {
		t.Errorf("should be 3X")
	}

}

func TestDecode(t *testing.T) {

	if Decode('X', []rune{}) != "X" {
		t.Errorf("should be X")
	}

	if Decode('X', []rune{'1', '0'}) != "XXXXXXXXXX" {
		t.Errorf("should be XXXXXXXXXX")
	}
}
