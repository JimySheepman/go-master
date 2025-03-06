package main

/*
#include <stdio.h>
#include <stdlib.h>

void out(char* in) {
  printf("%s\n", in);
}
*/
import "C"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"sync"
	"time"
	"unicode/utf8"
	"unsafe"
)

type Shades func() error

var shades = []Shades{
	shade01, shade02, shade03, shade04, shade05,
	shade06, shade07, shade08, shade09, shade10,
	shade11, shade12, shade13, shade14, shade15,
	shade16, shade17, shade18, shade19, shade20,
	shade21, shade22, shade23, shade24, shade25,
	shade26, shade27, shade28, shade29, shade30,
}

// Accidental Variable Shadowing
func shade01() error {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)

	return nil
}

// Array Function Arguments
func shade02() error {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(x)

	fmt.Println(x)

	return nil
}

// Accessing Non-Existing Map Keys
func shade03() error {
	x := map[string]string{"one": "a", "two": "", "three": "c"}

	// Bad
	if v := x["two"]; v == "" {
		fmt.Println("no entry")
	}

	// Good
	if _, ok := x["two"]; !ok {
		fmt.Println("no entry")
	}

	return nil
}

// Strings Are Not Always UTF8 Text
func shade04() error {
	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1))

	data2 := "A\xfeC"
	fmt.Println(utf8.ValidString(data2))

	return nil
}

// String Length
func shade05() error {
	data := "♥"
	fmt.Println(len(data))

	fmt.Println(utf8.RuneCountInString(data))

	data = "é"
	fmt.Println(len(data))
	fmt.Println(utf8.RuneCountInString(data))

	return nil
}

// Iteration Values For Strings in "range" Clauses
func shade06() error {
	data := "A\xfe\x02\xff\x04"
	for _, v := range data {
		fmt.Printf("%#x ", v)
	}

	fmt.Println()
	for _, v := range []byte(data) {
		fmt.Printf("%#x ", v)
	}
	fmt.Println()

	return nil
}

// Iterating Through a Map Using a "for range" Clause
func shade07() error {
	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range m {
		fmt.Println(k, v)
	}

	return nil
}

// Fallthrough Behavior in "switch" Statements
func shade08() error {
	isSpace := func(ch byte) bool {
		switch ch {
		case ' ':
		case '\t':
			return true
		}
		return false
	}

	fmt.Println(isSpace('\t'))
	fmt.Println(isSpace(' '))

	isSpace2 := func(ch byte) bool {
		switch ch {
		case ' ', '\t':
			return true
		}
		return false
	}

	fmt.Println(isSpace2('\t')) //prints true (ok)
	fmt.Println(isSpace2(' '))

	return nil
}

// Operator Precedence Differences
func shade09() error {
	fmt.Printf("0x2 & 0x2 + 0x4 -> %#x\n", 0x2&0x2+0x4)
	//prints: 0x2 & 0x2 + 0x4 -> 0x6
	//Go:    (0x2 & 0x2) + 0x4
	//C++:    0x2 & (0x2 + 0x4) -> 0x2

	fmt.Printf("0x2 + 0x2 << 0x1 -> %#x\n", 0x2+0x2<<0x1)
	//prints: 0x2 + 0x2 << 0x1 -> 0x6
	//Go:     0x2 + (0x2 << 0x1)
	//C++:   (0x2 + 0x2) << 0x1 -> 0x8

	fmt.Printf("0xf | 0x2 ^ 0x2 -> %#x\n", 0xf|0x2^0x2)
	//prints: 0xf | 0x2 ^ 0x2 -> 0xd
	//Go:    (0xf | 0x2) ^ 0x2
	//C++:    0xf | (0x2 ^ 0x2) -> 0xf

	return nil
}

// Unexported Structure Fields Are Not Encoded
func shade10() error {

	type MyData struct {
		One int
		two string
	}

	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in)

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded))

	var out MyData
	_ = json.Unmarshal(encoded, &out)

	fmt.Printf("%#v\n", out)

	return nil
}

// App Exits With Active Goroutines
func shade11() error {
	doIt := func(workerId int) {
		fmt.Printf("[%v] is running\n", workerId)
		time.Sleep(3 * time.Second)
		fmt.Printf("[%v] is done\n", workerId)
	}

	workerCount := 2

	for i := 0; i < workerCount; i++ {
		go doIt(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("doIt all done!")

	var wg sync.WaitGroup
	doIt2 := func(workerId int, wq <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
		fmt.Printf("[%v] is running\n", workerId)
		defer wg.Done()
		for {
			select {
			case m := <-wq:
				fmt.Printf("[%v] m => %v\n", workerId, m)
			case <-done:
				fmt.Printf("[%v] is done\n", workerId)
				return
			}
		}
	}

	done := make(chan struct{})
	wq := make(chan interface{})
	workerCount2 := 2

	for i := 0; i < workerCount2; i++ {
		wg.Add(1)
		go doIt2(i, wq, done, &wg)
	}

	for i := 0; i < workerCount2; i++ {
		wq <- i
	}

	close(done)
	wg.Wait()
	fmt.Println("doIt2 all done!")

	return nil
}

// Sending to an Unbuffered Channel Returns As Soon As the Target Receiver Is Ready
func shade12() error {
	ch := make(chan string)

	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2" //won't be processed

	return nil
}

type data struct {
	num   int
	key   *string
	items map[string]bool
}

func (d *data) pmethod() {
	d.num = 7
}

func (d data) vmethod() {
	d.num = 8
	*d.key = "v.key"
	d.items["vmethod"] = true
}

// Methods with Value Receivers Can't Change the Original Value
func shade13() error {
	key := "key.1"
	d := data{1, &key, make(map[string]bool)}

	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)

	d.pmethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)

	d.vmethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)

	return nil
}

// Closing HTTP Connections
func shade14() error {
	req, err := http.NewRequest("GET", "http://golang.org", nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Close = true
	//or do this:
	//req.Header.Add("Connection", "close")
	//or do this:
	//t := &http.Transport{DisableKeepAlives: true}
	//client := &http.Client{Transport: t}

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(len(string(body)))
	return nil
}

// JSON Encoder Adds a Newline Character
func shade15() error {
	data := map[string]int{"key": 1}

	var b bytes.Buffer
	json.NewEncoder(&b).Encode(data)

	raw, _ := json.Marshal(data)

	if b.String() == string(raw) {
		fmt.Println("same encoded data")
	} else {
		fmt.Printf("'%s' != '%s'\n", raw, b.String())
	}

	return nil
}

// JSON Package Escapes Special HTML Characters in Keys and String Values
func shade16() error {
	data := "x < y"

	raw, _ := json.Marshal(data)
	fmt.Println(string(raw))

	var b1 bytes.Buffer
	json.NewEncoder(&b1).Encode(data)
	fmt.Println(b1.String())

	var b2 bytes.Buffer
	enc := json.NewEncoder(&b2)
	enc.SetEscapeHTML(false)
	enc.Encode(data)
	fmt.Println(b2.String())

	return nil
}

// Unmarshalling JSON Numbers into Interface Values
func shade17() error {
	records := [][]byte{
		[]byte(`{"status": 200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		if err := json.NewDecoder(bytes.NewReader(record)).Decode(&result); err != nil {
			fmt.Println("error:", err)
			return err
		}

		var sstatus string
		if err := json.Unmarshal(result.Status, &sstatus); err == nil {
			result.StatusName = sstatus
		}

		var nstatus uint64
		if err := json.Unmarshal(result.Status, &nstatus); err == nil {
			result.StatusCode = nstatus
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}

	return nil
}

// JSON String Values Will Not Be Ok with Hex or Other non-UTF8 Escape Sequences
func shade18() error {
	type config struct {
		Data string `json:"data"`
	}

	raw := []byte(`{"data":"\xc2"}`)
	var decoded config

	if err := json.Unmarshal(raw, &decoded); err != nil {
		fmt.Println(err)
	}

	type config2 struct {
		Data []byte `json:"data"`
	}

	raw2 := []byte(`{"data":"wg=="}`)
	var decoded2 config2

	if err := json.Unmarshal(raw2, &decoded2); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v", decoded2)

	return nil
}

// "Hidden" Data in Slices
func shade19() error {
	get := func() []byte {
		raw := make([]byte, 10000)
		fmt.Println(len(raw), cap(raw), &raw[0])
		return raw[:3]
	}

	d := get()
	fmt.Println(len(d), cap(d), &d[0])

	get2 := func() []byte {
		raw := make([]byte, 10000)
		fmt.Println(len(raw), cap(raw), &raw[0])
		res := make([]byte, 3)
		copy(res, raw[:3])
		return res
	}

	d2 := get2()
	fmt.Println(len(d2), cap(d2), &d2[0])

	return nil
}

// Slice Data "Corruption"
func shade20() error {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir2 =>", string(dir2))

	dir1 = append(dir1, "suffix"...)
	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})

	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir2 =>", string(dir2))

	fmt.Println("new path =>", string(path))

	path = []byte("AAAA/BBBBBBBBB")
	sepIndex = bytes.IndexByte(path, '/')
	dir1 = path[:sepIndex:sepIndex]
	dir2 = path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir2 =>", string(dir2))

	dir1 = append(dir1, "suffix"...)
	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})

	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir2 =>", string(dir2))
	fmt.Println("new path =>", string(path))

	return nil
}

// "Stale" Slices
func shade21() error {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) //prints 3 3 [1 2 3]

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) //prints 2 2 [2 3]

	for i := range s2 {
		s2[i] += 20
	}

	//still referencing the same array
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [22 23]

	s2 = append(s2, 4)

	for i := range s2 {
		s2[i] += 10
	}

	//s1 is now "stale"
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [32 33 14]

	return nil
}

// Breaking Out of "for switch" and "for select" Code Blocks
func shade22() error {
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			break loop
		}
	}

	fmt.Println("out!")
	return nil
}

// Deferred Function Call Argument Evaluation
func shade23() error {
	var i int = 1

	defer fmt.Println("result =>", func() int { return i * 2 }())
	i++

	j := 1
	defer func(in *int) { fmt.Println("result =>", *in) }(&j)

	j = 2

	return nil
}

// Failed Type Assertions
func shade24() error {
	var data1 interface{} = "great"

	if d, ok := data1.(int); ok {
		fmt.Println("[is an int] value =>", d)
	} else {
		fmt.Println("[not an int] value =>", d)
		//prints: [not an int] value => 0 (not "great")
	}

	var data2 interface{} = "great"

	if res, ok := data2.(int); ok {
		fmt.Println("[is an int] value =>", res)
	} else {
		fmt.Println("[not an int] value =>", data2)
		//prints: [not an int] value => great (as expected)
	}

	return nil
}

// Same Address for Different Zero-sized Variables
func shade25() error {
	type data struct{}

	a := &data{}
	b := &data{}

	if a == b {
		fmt.Printf("same address - a=%p b=%p\n", a, b)
		//prints: same address - a=0x1953e4 b=0x1953e4
	}

	return nil
}

const (
	azero = iota
	aone  = iota
)

const (
	info  = "processing"
	bzero = iota
	bone  = iota
)

// The First Use of iota Doesn't Always Start with Zero
func shade26() error {
	fmt.Println(azero, aone) //prints: 0 1
	fmt.Println(bzero, bone) //prints: 1 2

	return nil
}

// "nil" Interfaces and "nil" Interfaces Values
func shade27() error {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil) //prints: <nil> true
	fmt.Println(in, in == nil)     //prints: <nil> true

	in = data
	fmt.Println(in, in == nil) //prints: <nil> false
	//'data' is 'nil', but 'in' is not 'nil'

	return nil
}

// Read and Write Operation Reordering
func shade28() error {
	var _ = runtime.GOMAXPROCS(3)
	var a, b int

	u1 := func() {
		a = 1
		b = 2
	}

	u2 := func() {
		a = 3
		b = 4
	}

	p := func() {
		println(a)
		println(b)
	}

	go u1()
	go u2()
	go p()
	time.Sleep(1 * time.Second)

	return nil
}

// Preemptive Scheduling
func shade29() error {
	done := false

	go func() {
		done = true
	}()

	for !done {
		runtime.Gosched()
	}
	fmt.Println("done!")

	return nil
}

// Can't Call C Functions with Variable Arguments
func shade30() error {
	cstr := C.CString("go")
	C.out(cstr) //ok
	C.free(unsafe.Pointer(cstr))

	return nil
}
