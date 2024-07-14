package converter

import "testing"

func TestFormatingTime(t *testing.T) {
	result := FormatingTime("2024-07-13T10:26")
	expected := "2024-07-13 07:26:00.000000"
	if result != expected {
		t.Error("Expected", expected, "got", result)
	} else {
		t.Log("OK")
	}

	result = FormatingTime("2024-07-14T02:26")
	expected = "2024-07-13 23:26:00.000000"
	if result != expected {
		t.Error("Expected", expected, "got", result)
	} else {
		t.Log("OK")
	}

	result = FormatingTime("2024-07-01T02:26")
	expected = "2024-06-30 23:26:00.000000"
	if result != expected {
		t.Error("Expected", expected, "got", result)
	} else {
		t.Log("OK")
	}

	result = FormatingTime("2024-01-01T02:26")
	expected = "2023-12-31 23:26:00.000000"
	if result != expected {
		t.Error("Expected", expected, "got", result)
	} else {
		t.Log("OK")
	}
}
