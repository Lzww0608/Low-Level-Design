package main

import "fmt"

type Fifo struct {
	
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by FIFO strategy")
}