package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.UTC().Format(time.RFC3339))
	fmt.Println(now.Format(time.RFC3339))
}
