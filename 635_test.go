package LeetCode

import (
	"strings"
	"fmt"
	"testing"
	"strconv"
)

type LogSystem struct {
	Logs []LogEntry
}

type LogEntry struct {
	Id int
	Time int64
}

func Constructor() LogSystem {
	return LogSystem{}
}


func (this *LogSystem) Put(id int, timestamp string)  {
	logentry := LogEntry{}
	logentry.Id = id
	// time string parese
	times,_ := strconv.Atoi(strings.Join(strings.Split(timestamp,":"),""))
	logentry.Time = int64(times)
	this.Logs = append(this.Logs,logentry)
}

func (this *LogSystem) Retrieve(s string, e string, gra string) []int {
	start,_ := strconv.Atoi(strings.Join(strings.Split(s,":"),""))
	end,_ := strconv.Atoi(strings.Join(strings.Split(e,":"),""))

	startt := int64(start)
	endd := int64(end)
	var num int64
	var res []int
	switch gra {
	case "Year":
		num  = 10000000000
	case "Month":
		num  = 100000000
	case "Day":
		num  = 1000000
	case "Hour":
		num  = 10000
	case "Minute":
		num  = 100
	case "Second":
		num  = 1
	}
	for _,v := range this.Logs {
		if (startt/num <= v.Time/num) && (v.Time/num <= endd/num) {
			res = append(res,v.Id)
		}
	}
	return res
}

func TestPut(t *testing.T) {
	obj := Constructor()
	obj.Put(1, "2017:01:01:23:59:59");
	obj.Put(2, "2017:01:01:22:59:59");
	obj.Put(3, "2016:01:01:00:00:00");
	fmt.Println(obj.Retrieve("2016:01:01:01:01:01","2017:01:01:23:00:00","Year"))
	fmt.Println(obj.Retrieve("2016:01:01:01:01:01","2017:01:01:23:00:00","Hour"))
}

