package main

import (
	"errors"
	"fmt"
	"reflect"
)

func merge(a, b interface{}) (interface{}, error) {
	if a == nil && b == nil {
		return nil, errors.New("both arguments are nil")
	}
	if a == nil {
		return b, nil
	}
	if b == nil {
		return a, nil
	}

	var res []interface{}

	if reflect.TypeOf(a).Kind() == reflect.Slice {
		slice := reflect.ValueOf(a)
		for i := 0; i < slice.Len(); i++ {
			if reflect.ValueOf(slice.Index(i).Interface()).Kind() == reflect.Slice { // if slice have nested slice
				nestedSlice, err := merge(res, slice.Index(i).Interface()) // recursive call for nested slice
				if err != nil {
					panic(err)
				}
				res = nestedSlice.([]interface{})
			} else { // if slice item is not slice then append to res directly
				res = append(res, slice.Index(i).Interface())
			}
		}
	} 

	// If 'b' is a slice, append its individual elements to the result
	if reflect.TypeOf(b).Kind() == reflect.Slice {
		slice := reflect.ValueOf(b)
		for i := 0; i < slice.Len(); i++ {
			if reflect.ValueOf(slice.Index(i).Interface()).Kind() == reflect.Slice {
				nestedSlice, err := merge(res, slice.Index(i).Interface())
				if err != nil {
					panic(err)
				}
				res = nestedSlice.([]interface{})
			} else {
				res = append(res, slice.Index(i).Interface())
			}
		}
	} else {
		// If 'b' is not a slice, directly append it to the result
		res = append(res, b)
	}
	return res, nil
}

func main() {
	a := []interface{}{6, 7, 8, []int{8, 9}, []interface{}{"a", "xfy", 20.0}}
	b := []interface{}{1, 2, 3, 5, []int{3, 4}, []interface{}{true, "abh", 2}}
	
	merged, err := merge(a, b)
	// merged, err := merge(nil, nil)
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	fmt.Printf("\nSlice a : %v\nSlice b : %v", a, b)
	fmt.Printf("\nMerged Slice : %v\n\n",merged)
}
