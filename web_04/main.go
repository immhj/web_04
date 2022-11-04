package main

import (
	"fmt"
	"sync"
	"time"
)

var str1 string
var T1 time.Time
var s1 []string
var s []string
var str string
var T time.Time
var wg sync.WaitGroup

func tips() {
	fmt.Println("1.单次日程提醒功能")
	fmt.Println("2.重复日程提醒功能")
	fmt.Println("输入-1结束")
}
func My_Function() (err error) {
	var now time.Time
	for {
		time.Sleep(500 * time.Millisecond)
		now = time.Now()
		now.Format("2006-01-02 15:04:05")
		if T.Sub(now) < time.Second {
			fmt.Println(s[0])
			s = s[1:]
			break
		}
	}
	wg.Done()
	return
}
func TW_Function() (err error) {
	var now time.Time
	for {
		time.Sleep(500 * time.Millisecond)
		now = time.Now()
		if T1.Sub(now) < time.Second {
			fmt.Println(s1[0])
			T1.Add(24 * time.Hour)
		}
	}
	return
}
func main() {
	var op int
	for {
		tips()
		fmt.Println("请输入序号")
		fmt.Scan(&op)
		if op == 1 {
			loc, err := time.LoadLocation("Asia/Shanghai")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("请设置日程提醒的时间（格式为：2006/01/02/15:04:05）")
			fmt.Scan(&str)
			T, err = time.ParseInLocation("2006/01/02/15:04:05", str, loc)
			if err != nil {
				panic(err)
			}
			fmt.Println("请设置日程提醒的内容")
			fmt.Scan(&str)
			s = append(s, str)
			wg.Add(1)
			go My_Function()
		} else if op == 2 {
			loc, err := time.LoadLocation("Asia/Shanghai")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("请设置日程提醒的时间（格式为：2006/01/02/15:04:05）")
			fmt.Scan(&str1)
			T1, err = time.ParseInLocation("2006/01/02/15:04:05", str1, loc)
			if err != nil {
				panic(err)
			}
			fmt.Println("请设置日程提醒的内容")
			fmt.Scan(&str1)
			s1 = append(s1, str)
			fmt.Println("日程将在每天这个时刻提醒")
			go TW_Function()
		} else if op == -1 {
			wg.Wait()
			return
		}
	}
}
