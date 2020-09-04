package table

import "testing"

func TestDoMath(t *testing.T) {
	result, err := DoMath(2, 2, "+")
	if result != 4 {
		t.Error("Should have been 4, got", result)
	}
	if err != nil {
		t.Error("Should have been nil error, got", err)
	}
	result2, err2 := DoMath(2, 2, "-")
	if result2 != 0 {
		t.Error("Should have been 0, got", result2)
	}
	if err2 != nil {
		t.Error("Should have been nil error, got", err2)
	}
	result3, err3 := DoMath(2, 2, "*")
	if result3 != 4 {
		t.Error("Should have been 4, got", result3)
	}
	if err3 != nil {
		t.Error("Should have been nil error, got", err3)
	}
	result4, err4 := DoMath(2, 2, "/")
	if result4 != 1 {
		t.Error("Should have been 1, got", result4)
	}
	if err4 != nil {
		t.Error("Should have been nil error, got", err4)
	}
}

func TestDoMathTable(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   string
	}{
		{"addition", 2, 2, "+", 4, ""},
		{"subtraction", 2, 2, "-", 0, ""},
		{"multiplication", 2, 2, "*", 4, ""},
		{"division", 2, 2, "/", 1, ""},
		{"bad_division", 2, 0, "/", 0, `division by zero`},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error message `%s`, got `%s`",
					d.errMsg, errMsg)
			}
		})
	}
}
