package main

import "fmt"
import "github.com/influxdata/influxdb/pkg/estimator/hll"

type Counter struct {
	hll        *hll.Plus
	hllCounter []*hll.Plus
	max        int
}

func main() {
	counter := NewCounter()
	b := counter.Add("User 4569 visit home")
	fmt.Println(b)
	b = counter.Add("User 4569 visit home")
	fmt.Println(b)
	b = counter.Add("User 4569 visit home")
	fmt.Println(b)
	b = counter.Add("User 888 visit home")
	fmt.Println(b)
	b = counter.HowMany("User 888 visit home")
	fmt.Println(b)
	b = counter.HowMany("User 4569 visit home")
	fmt.Println(b)
}

func (counter *Counter) Add(s string) int {
	el := []byte(s)
	var i int
	var cek bool
	i = 0
	for i < counter.max {
		cek = counter.cekCounterAdd(i, el)
		i = i + 1
		if cek == true {
			return i
		}
	}
	return i
}

func (counter *Counter) HowMany(s string) int {
	el := []byte(s)
	var cek bool
	var i int
	i = 0
	for i < counter.max {
		cek = counter.cekCounter(i, el)
		if cek == true {
			return i
		}
		i = i + 1
	}
	return i
}

func (counter *Counter) cekCounter(i int, element []byte) bool {
	countBefore := counter.hllCounter[i].Count()
	var clone = counter.hllCounter[i].Clone()
	clone.Add(element)
	countAfter := clone.Count()
	if countAfter != countBefore {
		return true
	} else {
		return false
	}
}

func (counter *Counter) cekCounterAdd(i int, element []byte) bool {
	countBefore := counter.hllCounter[i].Count()
	counter.hllCounter[i].Add(element)
	countAfter := counter.hllCounter[i].Count()
	if countAfter != countBefore {
		return true
	} else {
		return false
	}
}

func NewCounter() Counter {
	var i int
	i = 0
	data, _ := hll.NewPlus(18)
	counter := Counter{}
	counter.hll = data
	counter.max = 8
	counter.hllCounter = make([]*hll.Plus, counter.max)
	for i < counter.max {
		counter.hllCounter[i] = generateHll()
		i += 1
	}

	return counter
}

func generateHll() *hll.Plus {
	np, _ := hll.NewPlus(18)
	return np
}
