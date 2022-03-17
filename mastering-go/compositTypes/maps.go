package main

import (
	"fmt"
	"sort"
)

func main() {
	test1()
	test2()
}

func test1() {
	// maps: a refernce to hash table. so map is reference type.
	// key type must be cmparable with ==

	// declare a map. // zero-value of map itself, is nil so m == nil
	var m map[int]string //m == nil    len(m)==0
	fmt.Println(m)       // map[]

	// create by make
	m1 := make(map[int]string) // m1 != nil	len(m1) == 0
	fmt.Println(m1)

	fmt.Println(m == nil, m1 == nil) //true	false
	fmt.Println(len(m), len(m1))     // 0 0

	// Most operations on map like lookup, delete, len, range are safe on nil map and are like empty map.
	//but store value to nil map causes panic:
	//m[1] = "name"	// panic. because m == nil. map must be allocated before use for storing data. panic: assignment to entry in nil map
	fmt.Println("lookup nil map: ", m1[25]) // "" -> nil maps are like empty maps (except when assign to them)

	m1[10] = "ali" //sunbscript notation. if key=10 is not exist, create it and if exist edit it
	m1[20] = "reza"

	fmt.Println(m1[20]) //in subscript notation (m1[20]) if key=20 exists, returns its value, if not exist return zero-value of value type (int)

	r1 := m1[1000]            // r1 == "" because key 1000 is not exist
	fmt.Println(r1, r1 == "") // "" true

	r2, ok := m1[1000] //if key=1000 exist: ok==true r2==value, if key = 1000 is not exist: ok==false r2 == ""( vlaue-type zero-value)
	fmt.Println(r2, ok)

	//create by map literal
	m2 := map[int]string{
		10: "ali",
		20: "reza",
		30: "javad",
		40: "mohamad",
		50: "shayan",
	}

	// built-in len function
	fmt.Println("map element count: ", len(m2)) //map element count:  5

	delete(m2, 20) //remove item with key = 20

	fmt.Println(m2) // map[10:ali 30:javad 40:mohamad 50:shayan]

	// a map element is not a variable and can not take its address
	//pm := &m1[20]   // ERROR

	// order of map elements is not specified and may change in each execution
	for key, value := range m2 {
		fmt.Printf("%d\t%s\n", key, value)
	}

	for key := range m2 {
		fmt.Println(key)
	}

	for _, value := range m2 {
		fmt.Println(value)
	}

	sortMap(m2)

	// maps are not comparable and can not be compared to each other using ==, ... .only comparison against nil is valid: m1 == nil

	ma1 := map[int]string{10: "a", 20: "b", 40: "c"}
	ma2 := map[int]string{10: "a", 20: "b"}
	fmt.Println("equal:", isEqual(ma1, ma2))

}

// map as set
func test2() {
	namesSet := make(map[string]bool)

	namesSet["ali"] = true
	namesSet["reza"] = true

	//exist
	if namesSet["ali"] {
		fmt.Println("ali exist")
	} else {
		fmt.Println("ali not exist")
	}

	if !namesSet["bbb"] {
		fmt.Println("bbb not exist")
	} else {
		fmt.Println("bbb exist")
	}
}

func sortMap(m map[int]string) {
	var keys []int = make([]int, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	fmt.Println(keys)
	sort.Ints(keys)
	fmt.Println(keys)

	for _, v := range keys {
		fmt.Printf("%d\t%s\n", v, m[v])
	}
}

func isEqual(a, b map[int]string) bool {
	if len(a) != len(b) {
		return false
	}

	for akey, avalue := range a {
		if bvalue, ok := b[akey]; !ok || avalue != bvalue {
			return false
		}
	}
	return true
}
