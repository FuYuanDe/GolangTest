// main.go
package main

import (
	"fmt"
	"strconv"
	"time"
)

func GetSeconds() string {
	now := time.Now()
	secs := now.Unix()
	return strconv.FormatInt(secs, 10)

}

func main() {
	fmt.Println(GetSeconds())
}
