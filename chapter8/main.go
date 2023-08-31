package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//Strings package
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Index("test", "e"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	fmt.Println(strings.Repeat("a", 5))
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	fmt.Println(strings.ToLower("TeST"))
	fmt.Println(strings.ToUpper("test"))
	//convert string to a slice of bytes
	arr := []byte("test")
	//convert slice of bytes to string
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(arr)
	fmt.Println(str)

	//Readers and Writers
	var buf bytes.Buffer
	buf.Write([]byte("test"))
	fmt.Println(buf.String())

	//Files and Folders
	// os.Mkdir("testdir", 0777)
}
