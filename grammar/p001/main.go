/**
 * @Author: Noaghzil
 * @Date:   2024-11-29 19:39:04
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2024-11-29 20:03:13
 */
package main

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	s := new(Show)
	s.Param["RMB"] = 10000
}
