package converter

import "testing"

func TestFormatingTime(t *testing.T) {
	result := FormatingTime("2024-07-13T10:26")
	expected := "2024-07-13 13:26:00:00"
	if result != expected {
		t.Error("Expected", expected, "got", result)
	} else {
		t.Log("OK")
	}

	result = FormatingTime("2024-07-13T23:26")
	expected = "2024-07-14 02:26:00:00"
	if result != expected {
		t.Error("Expected", expected, "got", result)
	} else {
		t.Log("OK")
	}

	result = FormatingTime("2024-07-31T23:26")
	expected = "2024-08-01 02:26:00:00"
	if result != expected {
		t.Error("Expected", expected, "got", result)
	} else {
		t.Log("OK")
	}

	result = FormatingTime("2024-02-28T23:26")
	expected = "2024-02-29 02:26:00:00"
	if result != expected {
		t.Error("Expected", expected, "got", result)
	} else {
		t.Log("OK")
	}

	result = FormatingTime("2024-12-31T23:26")
	expected = "2025-01-01 02:26:00:0"
	if result != expected {
		t.Error("Expected", expected, "got", result)
	} else {
		t.Log("OK")
	}
}
