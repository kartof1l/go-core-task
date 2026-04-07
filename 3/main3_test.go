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
    t.Errorf("ошибка добавления ожидалось key1=10, получено %d, exists=%v", val, exists)
}

val, exists = m.Get("key2")
switch {
case !exists || val != 20:
    t.Errorf("ошибка добавления ожидалось key2=20, получено %d, exists=%v", val, exists)
}

val, exists = m.Get("key3")
switch {
case !exists || val != 30:
    t.Errorf("ошибка добавления ожидалось key3=30, получено %d, exists=%v", val, exists)
}
}

func TestAddOverwrite(t *testing.T) {
    m := New()
    
    m.Add("key1", 10)
    
    m.Add("key1", 20)
    
    val, exists := m.Get("key1")
    if !exists || val != 20 {
        t.Errorf("ошибка перезаписи ожидалось 20, получено %d", val)
    }
}

func TestRemove(t *testing.T) {
    m := New()
    
    m.Add("key1", 10)
    m.Add("key2", 20)
    
    m.Remove("key1")
    
    if m.Exists("key1") {
        t.Error("ошибка удаления key1 все еще существует")
    }
    
    if !m.Exists("key2") {
        t.Error("ошибка удаления key2 был неожиданно удален")
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
        t.Errorf("ошибка копирования ожидался размер 3, получено %d", len(copiedMap))
    }
    
    if val, exists := copiedMap["key1"]; !exists || val != 10 {
        t.Errorf("ошибка копирования ожидалось key1=10, получено %d", val)
    }
    
    if val, exists := copiedMap["key2"]; !exists || val != 20 {
        t.Errorf("ошибка копирования ожидалось key2=20, получено %d", val)
    }
    
    if val, exists := copiedMap["key3"]; !exists || val != 30 {
        t.Errorf("ошибка копирования ожидалось key3=30, получено %d", val)
    }
    
    copiedMap["key1"] = 100
    originalVal, _ := m.Get("key1")
    if originalVal == 100 {
        t.Error("ошибка копирования изменение копии повлияло на оригинальную карту")
    }
}

func TestExists(t *testing.T) {
    m := New()
    
    m.Add("key1", 10)
    
    if !m.Exists("key1") {
        t.Error("ключ1 должен быть")
    }
    
    if m.Exists("key2") {
        t.Error("ключа два не должно быть")
    }
    
    m.Add("key3", 30)
    m.Remove("key3")
    
    if m.Exists("key3") {
        t.Error("ключа три после ремува недолжно быть")
    }
}

func TestGet(t *testing.T) {
    m := New()
    
    m.Add("key1", 10)
    m.Add("key2", 20)
    
    val, exists := m.Get("key1")
    if !exists {
        t.Error("получение ключа провалено")
    }
    if val != 10 {
        t.Errorf("получение провалено, ожидалось 10, а вышло %d", val)
    }
    
    val, exists = m.Get("nonexistent")
    if exists {
        t.Error("получение провалено nonexistent недолжно быть")
    }
    if val != 0 {
        t.Errorf("ожидалось нулевое значение 0, а получилось %d", val)
    }
}

func TestIntegration(t *testing.T) {
    m := New()
    
    m.Add("apple", 5)
    m.Add("banana", 7)
    m.Add("cherry", 9)
    
    if val, _ := m.Get("apple"); val != 5 {
        t.Errorf("должно быть 5")
    }
    
    if !m.Exists("banana") {
        t.Errorf("должно быть банана")
    }
    
    m.Remove("banana")
    
    if m.Exists("banana") {
        t.Errorf("банана не должно быть после ремува")
    }
    
    copyMap := m.Copy()
    if len(copyMap) != 2 {
        t.Errorf("должно два элемента после копирования")
    }
    
    m.Add("date", 11)
    
    if !m.Exists("date") {
        t.Errorf("должно содердать date")
    }
    
    if _, exists := copyMap["date"]; exists {
        t.Errorf("у копи не должно быть data")
    }
}

func TestEmptyMap(t *testing.T) {
    m := New()
    
    if m.Exists("any") {
        t.Error("Пустая мапа на ЕКСисТ должно вернуть false")
    }
    
    if _, exists := m.Get("any"); exists {
        t.Error("Пустая мапа должно возвращать на ГЕТ false")
    }
    
    copyMap := m.Copy()
    if len(copyMap) != 0 {
        t.Error("Пустая мапа, должно возвращать на копи -  map")
    }
    
    m.Remove("any")
}