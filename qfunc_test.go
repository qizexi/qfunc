package qfunc

import (
	"fmt"
	"testing"
	"time"
)

func TestInSliceＳtring(t *testing.T) {
	tmpSlice := []string{"goods", "hi", "look"}
	if !InSliceＳtring(tmpSlice, "goods") {
		t.Log("测试失败，ｇｏｏｄｓ应存在！")
		t.Fail()
	}
	if InSliceＳtring(tmpSlice, "2") {
		t.Log("测试失败，2不应存在！")
		t.Fail()
	}
}

func TestDecodeJson(t *testing.T) {
	jsontext := `{"name":"qizexi", "sex":"male"}`
	jsontext_err := `{"name":"qizexi",}`

	rsmap, err := DecodeJson(jsontext)
	if err != nil {
		t.Log("测试失败，jsontext应是正确的ｊｓｏｎ格式！")
		t.Fail()
	}
	_ = rsmap

	rsmap1, err := DecodeJson(jsontext_err)
	if err == nil {
		t.Log("测试失败，jsontext_err应是非正确的ｊｓｏｎ格式！")
		t.Fail()
	}
	_ = rsmap1
}

func TestTime2Date(t *testing.T) {
	t1 := time.Now().Unix()

	st := Time2Date(t1)
	if len(st) != 19 {
		t.Log("测试失败，时间转换结果长度不对，应该为19个长度，实际为：" + fmt.Sprintf("%d", len(st)))
		t.Fail()
	}
}

func TestDate2Time(t *testing.T) {
	date1 := "2016-11-12 22:59:59"
	date2 := "2016-11-12 22.59.59"
	date3 := "2016-11-12"
	date5 := "20161112 22:59:59"

	t1 := Date2Time(date1, -8)
	if t1 == 0 {
		t.Log("转换时间失败，ｄａｔｅ1时间格式应为正确的！")
		t.Fail()
	}

	t2 := Date2Time(date2, -8)
	if t2 != 0 {
		t.Log("转换时间失败，ｄａｔｅ2时间格式不应为正确的！")
		t.Fail()
	}

	t3 := Date2Time(date3, -8)
	if t3 != 0 {
		t.Log("转换时间失败，ｄａｔｅ3时间格式不应为正确的,date3格式不完整！")
		t.Fail()
	}

	t5 := Date2Time(date5, -8)
	if t5 != 0 {
		t.Log("转换时间失败，ｄａｔｅ5时间格式不应为正确的！")
		t.Fail()
	}
}

func TestMap2Query(t *testing.T) {
	m := make(map[string]string)
	m["id"] = "1001"
	m["title"] = "怎么才能成为最好的程序员！"

	rs := Map2Query(m)
	if rs == "" {
		t.Log("测试失败，ｍ应为正确格式！")
		t.Fail()
	}

	v := make(map[string]string)

	rs1 := Map2Query(v)
	if rs1 != "" {
		t.Log("测试失败，ｖ为空！")
		t.Fail()
	}
}

func TestGetRandStr(t *testing.T) {
	rs := GetRandStr("test_")
	if rs == "" {
		t.Log("测试失败，ＧｅｔRandStr函数不正确！")
		t.Fail()
	}

	rs1 := GetRandStr("")
	if rs1 == "" {
		t.Log("测试失败，ＧｅｔRandStr函数不正确！")
		t.Fail()
	}
}
