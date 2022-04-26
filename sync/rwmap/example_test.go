/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package rwmap

import "fmt"

func ExampleRWMap_Del() {
	m := NewRWMap(3)
	m.Set("name", "jack")
	m.Set("age", 26)
	m.Set("sex", "man")
	m.Del("name")
	fmt.Println(m.Map())
	// Output:
	// map[age:26 sex:man]
}

func ExampleRWMap_Map() {
	m := NewRWMap(2)
	m.Set("name", "jack")
	m.Set("age", 26)
	newMap := m.Map()
	newMap["name"] = "tom"
	fmt.Println(m.Map())
	// Output:
	// map[age:26 name:jack]
}

func ExampleRWMap_Each() {
	m := NewRWMap(2)
	m.Set("name", "jack")
	m.Set("age", 26)

	m.Each(func(key string, val interface{}) bool {
		fmt.Printf("key:%s --- val:%v \n", key, val)

		return true
	})
	// Output:
	// key:age --- val:26
	// key:name --- val:jack
}

func ExampleRWMap_Get() {
	m := NewRWMap(2)
	m.Set("name", "jack")
	fmt.Println(m.Get("name"))
	fmt.Println(m.Get("age"))
	//	Output:
	//	jack true
	//	<nil> false
}

func ExampleRWMap_Len() {
	m := NewRWMap(3)
	m.Set("name", "jack")
	m.Set("age", 26)
	fmt.Println(m.Len())
	//	Output:
	//	2
}
