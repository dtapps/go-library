package gotime

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	fmt.Println("当前的时间：", Current().Now())
	fmt.Println("当前的时间：", Current().Format())
	fmt.Println("当前的时间：", Current().Timestamp())
	fmt.Println("当前的时间：", Current().TimestampWithMillisecond())
	fmt.Println("7100秒前的时间：", Current().BeforeSeconds(7100).Format())
	fmt.Println("2小时前的时间：", Current().BeforeHour(2).Format())
	fmt.Println("7100秒后的时间：", Current().AfterSeconds(7100).Format())
	fmt.Println("2小时后的时间：", Current().AfterHour(2).Format())
}

func TestStartOfDay(t *testing.T) {
	fmt.Println(Current().Format())
	fmt.Println(Current().StartOfDay().Format())
	fmt.Println(Current().EndOfDay().Format())
	fmt.Println(Current().Timestamp())
	fmt.Println(Current().StartOfDay().Timestamp())
	fmt.Println(Current().EndOfDay().Timestamp())
	fmt.Println(Current().BeforeDay(1).Format())
	fmt.Println(Current().BeforeDay(1).StartOfDay().Format())
	fmt.Println(Current().BeforeDay(1).EndOfDay().Format())
	fmt.Println(Current().AfterDay(1).Format())
	fmt.Println(Current().AfterDay(1).StartOfDay().Format())
	fmt.Println(Current().AfterDay(1).EndOfDay().Format())
}

func TestDiff(t *testing.T) {
	fmt.Println(Current().DiffInHourWithAbs(SetCurrentParse("2021-11-26 14:50:00").Time))
	fmt.Println(Current().DiffInHour(SetCurrentParse("2021-11-26 14:50:00").Time))
	fmt.Println(Current().DiffInMinutesWithAbs(SetCurrentParse("2021-11-26 14:50:00").Time))
	fmt.Println(Current().DiffInMinutes(SetCurrentParse("2021-11-26 14:50:00").Time))
}

func TestUnix(t *testing.T) {
	fmt.Println(SetCurrentUnix(1640067240).Format())
}
