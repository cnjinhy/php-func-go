package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"testing"
)

func TestEmpty(t *testing.T) {
	var test1 int
	test1 = 1
	result := php.Empty(test1)
	expected := false
	if result != expected {
		t.Errorf("TestEmpty(test1) returned %t, expected %t", result, expected)
	}

	var test2 int64
	test2 = 0
	result = php.Empty(test2)
	//php.VarDump(result)
	expected = true
	if result != expected {
		t.Errorf("TestEmpty(test2) returned %t, expected %t", result, expected)
	}

	var test21 int
	test21 = 0
	result = php.Empty(test21)
	//php.VarDump(result)
	expected = true
	if result != expected {
		t.Errorf("TestEmpty(test21) returned %t, expected %t", result, expected)
	}

	var test22 int64
	test22 = 1
	result = php.Empty(test22)
	//php.VarDump(result)
	expected = false
	if result != expected {
		t.Errorf("TestEmpty(test22) returned %t, expected %t", result, expected)
	}

	test3 := ""
	result = php.Empty(test3)
	expected = true
	if result != expected {
		t.Errorf("TestEmpty(test3) returned %t, expected %t", result, expected)
	}

	test4 := "0"
	result = php.Empty(test4)
	expected = true
	if result != expected {
		t.Errorf("TestEmpty(test4) returned %t, expected %t", result, expected)
	}

	test5 := false
	result = php.Empty(test5)
	expected = true
	if result != expected {
		t.Errorf("TestEmpty(test5) returned %t, expected %t", result, expected)
	}

	test6 := true
	result = php.Empty(test6)
	expected = false
	if result != expected {
		t.Errorf("TestEmpty(test6) returned %t, expected %t", result, expected)
	}

	test7 := "01"
	result = php.Empty(test7)
	expected = false
	if result != expected {
		t.Errorf("TestEmpty(test7) returned %t, expected %t", result, expected)
	}

	test8 := make(map[string]interface{})
	test8["a"] = "ValueA"
	result = php.Empty(test8)
	expected = false
	if result != expected {
		t.Errorf("TestEmpty(test8) returned %t, expected %t", result, expected)
	}

	test81 := make(map[string]interface{})
	test81["a"] = "ValueA"
	result = php.Empty(test81["b"])
	expected = true
	if result != expected {
		t.Errorf("TestEmpty(test81) returned %t, expected %t", result, expected)
	}

	test82 := make(map[string]interface{})
	test82["a"] = "ValueA"
	result = php.Empty(test81["b"])
	expected = true
	if result != expected {
		t.Errorf("TestEmpty(test82) returned %t, expected %t", result, expected)
	}

	test83 := make(map[string]map[string]int)
	test83["a"] = make(map[string]int)
	test83["a"]["s"] = 1
	test83["a"]["ss"] = 0
	result = php.Empty(test83["b"])
	expected = true
	if result != expected {
		t.Errorf("TestEmpty(test83) returned %t, expected %t", result, expected)
	}

	result = php.Empty(test83["a"])
	expected = false
	if result != expected {
		t.Errorf("TestEmpty(test83-a) returned %t, expected %t", result, expected)
	}

	result = php.Empty(test83["a"]["t"])
	expected = true
	if result != expected {
		t.Errorf("TestEmpty(test83-a-t) returned %t, expected %t", result, expected)
	}

	result = php.Empty(test83["a"]["s"])
	expected = false
	if result != expected {
		t.Errorf("TestEmpty(test83-a-s) returned %t, expected %t", result, expected)
	}

	result = php.Empty(test83["a"]["ss"])
	expected = true
	if result != expected {
		t.Errorf("TestEmpty(test83-a-ss) returned %t, expected %t", result, expected)
	}

	test9 := make(map[string]interface{})
	result = php.Empty(test9)
	expected = true
	if result != expected {
		t.Errorf("TestEmpty(test9) returned %t, expected %t", result, expected)
	}

}
