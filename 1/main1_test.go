package main

import (
	"testing"
)

func TestDetectType(t *testing.T) {
	test := []struct {
		name   string
		input  any
		expect string
	}{
		{"int", 42, "int"},
		{"float64", 3.14, "float64"},
		{"string", "hello", "string"},
		{"bool", true, "bool"},
		{"complex64", complex64(1 + 2i), "complex64"},
	}
	for _, te := range test {
		t.Run(te.name, func(t *testing.T) {
			result := DetectType(te.input)
			if result != te.expect {
				t.Errorf("DetectType(%v) =  %s, а   должно   было  %s \n", te.input, result, te.expect)
			}
		})
	}
}
func TestConverter(t *testing.T) {
	test := []struct {
		name   string
		input  []any
		expect string
	}{
		{"числа", []any{42, 3.14}, "423.14"},
		{"смеси", []any{42, "hello", true}, "42hellotrue"},
		{"пустота", []any{}, ""},
	}
	for _, te := range test {
		t.Run(te.name, func(t *testing.T) {
			result := Converter(te.input...)
			if result != te.expect {
				t.Errorf("Converter(%v) =  %s, а   должно   было  %s \n", te.input, result, te.expect)
			}
		})
	}
}
func TestRuneMaker(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []rune
	}{
		{
			name:     "простая строка",
			input:    "hello",
			expected: []rune{'h', 'e', 'l', 'l', 'o'},
		},
		{
			name:     "русские символы",
			input:    "привет",
			expected: []rune{'п', 'р', 'и', 'в', 'е', 'т'},
		},
		{
			name:     "пустая строка",
			input:    "",
			expected: []rune{},
		},
	}

	for _, te := range tests {
		t.Run(te.name, func(t *testing.T) {
			result := RuneMaker(te.input)
			if len(result) != len(te.expected) {
				t.Errorf("длина результата %d, ожидалось %d", len(result), len(te.expected))
			}
			for i := range result {
				if result[i] != te.expected[i] {
					t.Errorf("позиция %d: %c, ожидалось %c", i, result[i], te.expected[i])
				}
			}
		})
	}
}
func TestHashWithSalt(t *testing.T) {
	tests := []struct {
		name     string
		runes    []rune
		salt     string
		expected string
	}{
		{
			name:     "простой тест",
			runes:    []rune("hello"),
			salt:     "go-2024",
			expected: "проверяем что хэш не пустой",
		},
	}

	for _, te := range tests {
		t.Run(te.name, func(t *testing.T) {
			result := SaltyHashing(te.runes, te.salt)
			if len(result) != 64 { // 64 символа в hex
				t.Errorf("длина хэша %d, ожидалось 64", len(result))
			}
			if result == "" {
				t.Error("хэш не должен быть пустым")
			}
		})
	}
}
