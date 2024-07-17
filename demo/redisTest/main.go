package main

import (
	"fmt"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/cache/redisCache"
	"strconv"
)

func cacheBase(cache Interface.ICache) {
	cache.Prefix("base")

	for i := 0; i < 1000; i++ {
		cache.Set(strconv.Itoa(i), "value"+strconv.Itoa(i), 0)
	}

	for i := 0; i < 1000; i++ {
		data, err := cache.Get(strconv.Itoa(i))
		if err != nil {
			return
		}
		fmt.Println(data)
	}
}

func cacheList(cache Interface.ICache) {
	c, ok := cache.(Interface.CacheSupportList)
	if !ok {
		fmt.Println("cache does not support list")
	}

	for i := 0; i < 100; i++ {
		c.ListPush("list", "value"+strconv.Itoa(i))
	}

	data, err := c.ListRange("list", 0, 50)
	if err != nil {
		return
	}
	fmt.Println(data)

	/*	for i := 0; i < 100; i++ {
		data, err := c.ListPop("list")
		if err != nil {
			return
		}
		fmt.Println(data)
	}*/

}

func cacheSet(cache Interface.ICache) {
	c, ok := cache.(Interface.CacheSupportSet)
	if !ok {
		fmt.Println("cache does not support set")
	}

	for i := 0; i < 100; i++ {
		c.SetAdd("demo", "value"+strconv.Itoa(i))
	}

	for i := 0; i < 100; i++ {
		data, err := c.SetMembers(strconv.Itoa(i))
		if err != nil {
			return
		}
		fmt.Println(data)
	}

	fmt.Println(c.SetIsMember("demo", "value1"))

}

func cacheObject(cache Interface.ICache) {
	c, ok := cache.(Interface.CacheSupportHash)
	if !ok {
		fmt.Println("cache does not support object")
	}
	type Object struct {
		Name  string `redis:"name"`
		Age   int    `redis:"age"`
		Value []byte `redis:"value"`
	}
	old := Object{
		Name: "name",
		Age:  10,
		Value: []byte{
			1, 2, 3, 4, 5,
		},
	}

	err := c.HashSetObject("object2", &old, 0)
	if err != nil {
		fmt.Println("1", err)
		return
	}

	var obj Object
	err = c.HashGetObject("object2", &obj)
	if err != nil {
		fmt.Println("2", err)
		return
	}
	fmt.Println(obj)
}

func main() {
	cache := redisCache.NewCache()

	//cacheBase(cache)
	//cacheList(cache)
	cacheObject(cache)
}
