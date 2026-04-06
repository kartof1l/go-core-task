package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
    m := New()
    
    m.Add("key1", 10)
    m.Add("key2", 20)
    m.Add("key3", 30)
    
    val, exists := m.Get("key1")
switch {
case !exists || val != 10:
    t.Errorf("Add failed: expected key1=10, got %d, exists=%v", val, exists)
}

val, exists = m.Get("key2")
switch {
case !exists || val != 20:
    t.Errorf("Add failed: expected key2=20, got %d, exists=%v", val, exists)
}

val, exists = m.Get("key3")
switch {
case !exists || val != 30:
    t.Errorf("Add failed: expected key3=30, got %d, exists=%v", val, exists)
}
}

func TestAddOverwrite(t *testing.T) {
    m := New()
    
    m.Add("key1", 10)
    
    m.Add("key1", 20)
    
    val, exists := m.Get("key1")
    if !exists || val != 20 {
        t.Errorf("Add overwrite failed: expected 20, got %d", val)
    }
}

func TestRemove(t *testing.T) {
    m := New()
    
    m.Add("key1", 10)
    m.Add("key2", 20)
    
    m.Remove("key1")
    
    if m.Exists("key1") {
        t.Error("Remove failed: key1 still exists")
    }
    
    if !m.Exists("key2") {
        t.Error("Remove failed: key2 was unexpectedly removed")
    }
    
    m.Remove("nonexistent")
}

func TestCopy(t *testing.T) {
    m := New()
    
    m.Add("key1", 10)
    m.Add("key2", 20)
    m.Add("key3", 30)
    
    copiedMap := m.Copy()
    
    if len(copiedMap) != 3 {
        t.Errorf("Copy failed: expected size 3, got %d", len(copiedMap))
    }
    
    if val, exists := copiedMap["key1"]; !exists || val != 10 {
        t.Errorf("Copy failed: expected key1=10, got %d", val)
    }
    
    if val, exists := copiedMap["key2"]; !exists || val != 20 {
        t.Errorf("Copy failed: expected key2=20, got %d", val)
    }
    
    if val, exists := copiedMap["key3"]; !exists || val != 30 {
        t.Errorf("Copy failed: expected key3=30, got %d", val)
    }
    
    copiedMap["key1"] = 100
    originalVal, _ := m.Get("key1")
    if originalVal == 100 {
        t.Error("Copy failed: modifying copy affected original map")
    }
}

func TestExists(t *testing.T) {
    m := New()
    
    m.Add("key1", 10)
    
    if !m.Exists("key1") {
        t.Error("Exists failed: key1 should exist")
    }
    
    if m.Exists("key2") {
        t.Error("Exists failed: key2 should not exist")
    }
    
    m.Add("key3", 30)
    m.Remove("key3")
    
    if m.Exists("key3") {
        t.Error("Exists failed: key3 should not exist after removal")
    }
}

func TestGet(t *testing.T) {
    m := New()
    
    m.Add("key1", 10)
    m.Add("key2", 20)
    
    val, exists := m.Get("key1")
    if !exists {
        t.Error("Get failed: key1 should exist")
    }
    if val != 10 {
        t.Errorf("Get failed: expected 10, got %d", val)
    }
    
    val, exists = m.Get("nonexistent")
    if exists {
        t.Error("Get failed: nonexistent key should not exist")
    }
    if val != 0 {
        t.Errorf("Get failed: expected zero value 0, got %d", val)
    }
}

func TestIntegration(t *testing.T) {
    m := New()
    
    m.Add("apple", 5)
    m.Add("banana", 7)
    m.Add("cherry", 9)
    
    if val, _ := m.Get("apple"); val != 5 {
        t.Errorf("Integration failed: apple should be 5")
    }
    
    if !m.Exists("banana") {
        t.Errorf("Integration failed: banana should exist")
    }
    
    m.Remove("banana")
    
    if m.Exists("banana") {
        t.Errorf("Integration failed: banana should be removed")
    }
    
    copyMap := m.Copy()
    if len(copyMap) != 2 {
        t.Errorf("Integration failed: copy should have 2 elements")
    }
    
    m.Add("date", 11)
    
    if !m.Exists("date") {
        t.Errorf("Integration failed: date should exist")
    }
    
    if _, exists := copyMap["date"]; exists {
        t.Errorf("Integration failed: copy should not have date")
    }
}

func TestEmptyMap(t *testing.T) {
    m := New()
    
    if m.Exists("any") {
        t.Error("Empty map: Exists should return false")
    }
    
    if _, exists := m.Get("any"); exists {
        t.Error("Empty map: Get should return false")
    }
    
    copyMap := m.Copy()
    if len(copyMap) != 0 {
        t.Error("Empty map: Copy should return empty map")
    }
    
    m.Remove("any")
}