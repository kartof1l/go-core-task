package main

import (
	"fmt"
	"reflect"
	"testing"
)
func TestSliceExample(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "обычный слайс с четными и нечетными",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "только четные числа",
			input:    []int{2, 4, 6, 8, 10},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "только нечетные числа",
			input:    []int{1, 3, 5, 7, 9},
			expected: []int{},
		},
		{
			name:     "пустой слайс",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "с отрицательными числами",
			input:    []int{-3, -2, -1, 0, 1, 2},
			expected: []int{-2, 0, 2},
		},
		{
			name:     "с нулем",
			input:    []int{0, 1, 2, 3},
			expected: []int{0, 2},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceExample(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sliceExample(%v) = %v, ожидалось %v", 
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		value    int
		expected []int
	}{
		{
			name:     "добавление в конец слайса",
			input:    []int{1, 2, 3},
			value:    4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "добавление в пустой слайс",
			input:    []int{},
			value:    5,
			expected: []int{5},
		},
		{
			name:     "добавление отрицательного числа",
			input:    []int{1, 2, 3},
			value:    -1,
			expected: []int{1, 2, 3, -1},
		},
		{
			name:     "добавление нуля",
			input:    []int{10, 20},
			value:    0,
			expected: []int{10, 20, 0},
		},
		{
			name:     "добавление в слайс с одним элементом",
			input:    []int{42},
			value:    100,
			expected: []int{42, 100},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.input))
			copy(original, tt.input)
			
			result := addElements(tt.input, tt.value)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("addElements(%v, %d) = %v, ожидалось %v", 
					tt.input, tt.value, result, tt.expected)
			}
		
			if !reflect.DeepEqual(tt.input, original) {
				t.Errorf("оригинальный слайс изменился: было %v, стало %v", 
					original, tt.input)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "обычный слайс",
			input: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "пустой слайс",
			input: []int{},
		},
		{
			name:  "слайс с одним элементом",
			input: []int{42},
		},
		{
			name:  "слайс с отрицательными числами",
			input: []int{-1, -2, -3, -4},
		},
		{
			name:  "слайс с повторяющимися числами",
			input: []int{1, 1, 2, 2, 3, 3},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := copySlice(tt.input)

			if !reflect.DeepEqual(result, tt.input) {
				t.Errorf("copySlice(%v) = %v, ожидалось %v", 
					tt.input, result, tt.input)
			}
		
			if len(tt.input) > 0 && &result[0] == &tt.input[0] {
				t.Error("копия и оригинал используют один и тот же массив")
			}
			
		
			if len(result) > 0 {
				result[0] = 999
		
				if tt.input[0] == 999 {
					t.Error("изменение копии повлияло на оригинал")
				}
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		index    int
		expected []int
	}{
		{
			name:     "удаление из середины",
			input:    []int{1, 2, 3, 4, 5},
			index:    2,
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "удаление первого элемента",
			input:    []int{1, 2, 3, 4, 5},
			index:    0,
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "удаление последнего элемента",
			input:    []int{1, 2, 3, 4, 5},
			index:    4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "удаление из слайса с одним элементом",
			input:    []int{42},
			index:    0,
			expected: []int{},
		},
		{
			name:     "отрицательный индекс (возврат копии)",
			input:    []int{1, 2, 3},
			index:    -1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "индекс больше длины (возврат копии)",
			input:    []int{1, 2, 3},
			index:    5,
			expected: []int{1, 2, 3},
		},
		{
			name:     "удаление второго элемента",
			input:    []int{10, 20, 30, 40},
			index:    1,
			expected: []int{10, 30, 40},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.input))
			copy(original, tt.input)
			
			result := removeElement(tt.input, tt.index)
			

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("removeElement(%v, %d) = %v, ожидалось %v", 
					tt.input, tt.index, result, tt.expected)
			}
			
			if !reflect.DeepEqual(tt.input, original) {
				t.Errorf("оригинальный слайс изменился: было %v, стало %v", 
					original, tt.input)
			}
			
			
			if tt.index >= 0 && tt.index < len(original) && len(result) > 0 {
			
				result[0] = 999
				if original[0] == 999 {
					t.Error("изменение результата повлияло на оригинал")
				}
			}
		})
	}
}
func TestRemoveElementEdgeCases(t *testing.T) {
	original := []int{10, 20, 30, 40, 50}

	for i := 0; i < len(original); i++ {
		t.Run(fmt.Sprintf("удаление_индекса_%d", i), func(t *testing.T) {
			result := removeElement(original, i)
			

			if len(result) != len(original)-1 {
				t.Errorf("длина результата %d, ожидалось %d", len(result), len(original)-1)
			}
			for _, v := range result {
				if v == original[i] {
					t.Errorf("элемент %d не должен быть в результате", original[i])
				}
			}
		})
	}
}