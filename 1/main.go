package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

func DetectType(v any) string {
	return reflect.TypeOf(v).String()
}
func Converter(values ...any) string {
	var build strings.Builder //строки копить в буфер
	for _, v := range values {
		build.WriteString(fmt.Sprintf("%v", v))
	}
	return build.String()
}
func RuneMaker(s string) []rune {
	return []rune(s)
}
func SaltyHashing(runes []rune, salt string) string {
	str := string(runes)
	middle := len(str)
	addSalt := str[:middle] + salt + str[middle:]
	hashed := sha256.Sum256([]byte(addSalt))
	return hex.EncodeToString(hashed[:])
}

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	fmt.Println("Типы переменных")
	fmt.Printf("numDecimal (%d) → тип: %s\n", numDecimal, DetectType(numDecimal))
	fmt.Printf("numOctal (%d) → тип: %s\n", numOctal, DetectType(numOctal))
	fmt.Printf("numHexadecimal (%d) → тип: %s\n", numHexadecimal, DetectType(numHexadecimal))
	fmt.Printf("pi (%f) → тип: %s\n", pi, DetectType(pi))
	fmt.Printf("name (%s) → тип: %s\n", name, DetectType(name))
	fmt.Printf("isActive (%t) → тип: %s\n", isActive, DetectType(isActive))
	fmt.Printf("complexNum (%v) → тип: %s\n", complexNum, DetectType(complexNum))
	fmt.Println()
	allVars := []interface{}{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}
	combined := Converter(allVars...) //перевод   типов в стринг
	fmt.Printf("Объединенная строка: %s\n", combined)
	fmt.Println()
	runes := RuneMaker(combined) //рунирование
	fmt.Printf("Срез рун (первые 10): %v\n", runes[:min(10, len(runes))])
	fmt.Printf("Длина среза рун: %d\n", len(runes))
	fmt.Println()
	salt := "go-2024" //соль
	hash := SaltyHashing(runes, salt)
	fmt.Printf("Соль: %s\n", salt)
	fmt.Printf("SHA256 хэш: %s\n", hash)
}
