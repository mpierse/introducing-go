package main

import (
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"mpierse/introducing-go/chapter8/math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sort"
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
	fmt.Println("Files and Folders")
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.New("this is how you make an error"))
	}

	//Container/List Package
	//the sort.Interfact requires 3 methods: Len, Less, and Swap
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	fmt.Println("unsorted:")
	fmt.Println(kids)
	sort.Sort(ByName(kids))
	fmt.Println("sorted:")
	fmt.Println(kids)

	//Hashes and Cryptography
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)

	//Servers
	//Http
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(
			"Content-Type",
			"text/html",
		)

		_, err = io.WriteString(w, `<html>
	<head><title>Chat</title></head>
	<body>
	Let's chat!
	</body>
	</html>`)
		if err != nil {
			return
		}
	})

	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
	fmt.Println(math.Average([]float64{}))

}

type Person struct {
	name string
	age  int
}
type ByName []Person

// these methods are required by sort.Interface
func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].name < ps[j].name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

//Servers
//TCP

type Listener interface {
	//Accept waits for and returns the next connection to the listener.
	Accept() (c net.Conn, err error)
	//Close closes the listener.
	//Any blocked Accept operations will be unblocked and return errors.
	Close() error
	//Addr returns the listener's network address.
	Addr() net.Addr
}

/* Exercise Answers
1. We use packages to organize code and make it reusable in other programs
2. Identifiers that are lower cased can only be used in the package they are defined in.
Identifiers that are upper cased can be used by other packages.
3. A package alias is used to give a package a different name when importing it.
It is defined in the import statement before the package name. i.e. `import m "math"`
4. See chapter8/math/math.go for code
5. You could document these functions by adding a comment on the line above the function
declaration.
*/
