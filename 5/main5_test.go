package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
    tests := []struct{
        name string
        a, b []int
        wantBool bool
        wantSlice []int
    }{
        {"basic", []int{65,3,58,678,64}, []int{64,2,3,43}, true, []int{3,64}},
        {"no intersection", []int{1,2,3}, []int{4,5,6}, false, []int{}},
        {"empty a", []int{}, []int{1,2,3}, false, []int{}},
        {"empty b", []int{1,2,3}, []int{}, false, []int{}},
        {"both empty", []int{}, []int{}, false, []int{}},
        {"duplicates", []int{1,1,2,3}, []int{1,2,2,4}, true, []int{1,2}},
        {"all intersect", []int{1,2,3}, []int{1,2,3}, true, []int{1,2,3}},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotBool, gotSlice := Intersection(tt.a, tt.b)
            if gotBool != tt.wantBool || !reflect.DeepEqual(gotSlice, tt.wantSlice) {
                t.Errorf("got (%v, %v), want (%v, %v)", gotBool, gotSlice, tt.wantBool, tt.wantSlice)
            }
        })
    }
}