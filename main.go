/**
 * Auth :   liubo
 * Date :   2018/10/24 15:12
 * Comment: 
 */

package main

import (
	"fmt"
	"github.com/irfansharif/cfilter"
)

func main() {
	cf := cfilter.New()

	// inserts 'buongiorno' to the filter
	cf.Insert([]byte("buongiorno"))

	test(cf, "buo")
	test(cf, "buongiorno")

	cf.Insert([]byte("hello world"))
	test(cf, "hello")
	test(cf, "hello world")

	// looks up 'hola' in the filter, may return false positive
	cf.Lookup([]byte("hola"))

	// returns 1 (given only 'buongiorno' was added)
	cf.Count()

	// tries deleting 'bonjour' from filter, may delete another element
	// this could occur when another byte slice with the same fingerprint
	// as another is 'deleted'
	cf.Delete([]byte("bonjour"))
}

func test(cf *cfilter.CFilter, word string) {
	fmt.Println("是否存在：" + word, cf.Lookup([]byte(word)))
}