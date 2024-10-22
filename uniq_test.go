package main

import (
	"reflect"
	"testing"
)

func TestCountOccurrences(t *testing.T) {
	opts := Options{count: true}
	input := []string{"line1", "line1", "line2", "line3", "line3", "line3"}
	expected := []string{"2 line1", "1 line2", "3 line3"}

	result := FormResult(input, opts)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindDuplicates(t *testing.T) {
	opts := Options{deplicate: true}
	input := []string{"line1", "line1", "line2", "line3", "line3"}
	expected := []string{"line1", "line3"}

	result := FormResult(input, opts)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindUnique(t *testing.T) {
	opts := Options{uniqStrings: true}
	input := []string{"line1", "line1", "line2", "line3", "line3"}
	expected := []string{"line2"}

	result := FormResult(input, opts)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSkipFields(t *testing.T) {
	input := "field1 field2 field3"
	numFields := 1
	expected := "field2 field3"

	result := SkipFields(input, numFields)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestSkipChars(t *testing.T) {
	input := "Hello, World!"
	numChars := 7
	expected := "World!"

	result := SkipChars(input, numChars)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestIgnoreCase(t *testing.T) {
	opts := Options{ignoreRegistr: true}
	input := []string{"Line1", "line1", "LINE2", "line2"}
	expected := []string{"Line1", "LINE2"}

	// Применяем форматирование с игнорированием регистра
	result := FormResult(input, opts)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
