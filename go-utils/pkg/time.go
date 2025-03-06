package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/golang/glog"
	"github.com/jinzhu/now"
)

func TimeTest() {
	loc, err := time.LoadLocation("Europe/Istanbul")
	if err != nil {
		return
	}

	ed := time.Now().In(loc).AddDate(0, 0, 1)
	ed = now.New(ed).EndOfDay()
	sd := ed.AddDate(0, 0, -7)
	fmt.Println("sd: ", sd)
	sd = now.New(sd).BeginningOfDay()

	fmt.Println("sd: ", sd)
	fmt.Println("ed: ", ed)

	glog.Infof("checkRefunds [%d] range %s - %s", 1, sd, ed)

	sdStr := sd.Format("20060102")
	edStr := ed.Format("20060102")

	fmt.Println("sdStr: ", sdStr)
	fmt.Println("edStr: ", edStr)

	fmt.Println(now.BeginningOfDay().Add(24 * time.Hour).In(loc))

	year, month, day := time.Now().Date()

	fmt.Printf("%d-%d-%d\n", year, int(month), day)
	fmt.Printf("%d-%d-%d 23:59:59\n", year, int(month), day)

	fmt.Println(now.BeginningOfDay().UTC())
	fmt.Println(now.EndOfDay().UTC())

	sd = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	ed = time.Date(year, month, day+1, 0, 0, 0, 0, time.UTC)

	fmt.Println(sd)
	fmt.Println(ed)

	currentTime := time.Now().In(loc)

	today := now.New(currentTime).BeginningOfDay().In(loc)
	tomorrow := today.Add(24 * time.Hour).In(loc)

	fmt.Println(currentTime)
	fmt.Println(today)
	fmt.Println(tomorrow)
}

func test1() {
	loc, _ := time.LoadLocation("Europe/Istanbul")
	value := "2022-12-17T11:15:20.767+03:00"
	layouts := []string{
		"2006-01-02T15:04:05",
		"20060102",
		time.RFC3339,
		"2006-01-02",
		"2006-01-02-15.04.05.000000",
		"02.01.2006",
		"2006-01-02 15:04",
	}

	for a, layout := range layouts {
		prosTime, err := time.ParseInLocation(layout, value, loc)
		if err == nil {
			log.Println("Test case: ", a+1)
			log.Println("Test error: ", err)
			log.Println("Test result: ", prosTime)
		}
		fmt.Println()
	}

	b := time.Now().Format(time.RFC3339)
	log.Println(b)
}

func test2() {
	Scopes := "a:b c:d"

	params := url.Values{}
	params.Add("a", Scopes)
	urlEncode := params.Encode()

	fmt.Println("urlEncode= ", urlEncode)
}

func test3() {
	diff := 53996
	expiresIn := 3600

	fmt.Println(time.Duration(expiresIn-60) * time.Millisecond)
	fmt.Println(time.Duration(diff) * time.Millisecond)
	fmt.Println(time.Duration(expiresIn-60)*time.Millisecond > time.Duration(diff)*time.Millisecond)
	fmt.Println(expiresIn - 60)
	fmt.Println(diff)
	fmt.Println(expiresIn-60 > diff)
}

func test4() {
	a := "2023-06-06T00:00:00Z"
	format := "2006-01-02T15:04:05Z"
	parsedDate, err := time.Parse(format, a)
	if err != nil {

		fmt.Println(err)
	}

	fmt.Println(parsedDate)
}

const RFC3339WithoutTimeZone = "2006-01-02T15:04:05"

func TimeTr() {
	f := "2023-01-09T00:00:00"
	s := time.Now()
	a := s.Format(f)
	b := s.Format(time.RFC3339)
	c := s.Format(RFC3339WithoutTimeZone)
	fmt.Println("s: ", s)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
}
