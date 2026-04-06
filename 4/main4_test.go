package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
    tests := []struct{
        name string
        s1, s2, want []string
    }{
        {"basic", []string{"apple","banana","cherry","date","43","lead","gno1"}, []string{"banana","date","fig"}, []string{"apple","cherry","43","lead","gno1"}},
        {"empty s1", []string{}, []string{"a","b"}, []string{}},
        {"empty s2", []string{"a","b"}, []string{}, []string{"a","b"}},
        {"all common", []string{"a","b"}, []string{"a","b"}, []string{}},
        {"no common", []string{"a","b"}, []string{"c","d"}, []string{"a","b"}},
        {"with duplicates", []string{"a","b","a","c"}, []string{"b"}, []string{"a","a","c"}},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) { 
            if got := FindDifference(tt.s1, tt.s2); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}