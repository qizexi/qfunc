//常用函数工具合集
//作者：齐泽西
//日期：2016-12-02
package qfunc

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

//字符串切片是否存在某个值
//slice 切片信息
//val 查询的值
//return 存在返回true 否则返回false
func InSliceＳtring(slice []string, val string) bool {
	for _, value := range slice {
		if value == val {
			return true
		}
	}
	return false
}

//json数据解析
//jsonstr 待解析的ｊｓｏｎ字符串
//return map[string]string, error
func DecodeJson(jsonstr string) (map[string]string, error) {
	//解析
	var v interface{}
	err := json.Unmarshal([]byte(jsonstr), &v)
	if err != nil {
		return nil, err
	}
	m := v.(map[string]interface{})

	//以map[string]string存储返回
	mapjson := make(map[string]string)
	for k, val := range m {
		mapjson[k] = fmt.Sprintf("%v", val)
	}

	if len(mapjson) == 0 {
		return nil, errors.New("解析ｊｓｏｎ数据失败！")
	}

	return mapjson, nil
}

//时间戳转日期格式函数
//time 待转换的时间戳
//return 返回转换后的日期格式yyyy-mm-dd hh:ii:ss
func Time2Date(ctime int64) string {
	tm := time.Unix(ctime, 0)
	return fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d",
		tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
}

//日期格式转时间戳
//datestr 日期，格式yyyy-mm-dd HH:ii:ss
//n 加多少个时区，默认０时区
//return 时间戳，转换错误返回0
func Date2Time(datestr string, n int) int64 {
	//分割时间：日期，时间
	arr := strings.Split(datestr, " ")
	//格式不对，返回0
	if len(arr) != 2 {
		return 0
	}

	//分解出年月日
	darr := strings.Split(arr[0], "-")
	//格式不对，返回0
	if len(darr) != 3 {
		return 0
	}

	//分解出时分秒
	tarr := strings.Split(arr[1], ":")
	//格式不对，返回0
	if len(tarr) != 3 {
		return 0
	}

	//字符串转整形
	y, _ := strconv.Atoi(darr[0])
	tmpm, _ := strconv.Atoi(darr[1])
	m := time.Month(tmpm)
	d, _ := strconv.Atoi(darr[2])
	h, _ := strconv.Atoi(tarr[0])
	i, _ := strconv.Atoi(tarr[1])
	s, _ := strconv.Atoi(tarr[2])

	//得到转换后的Time对象
	t := time.Date(y, m, d, h, i, s, 0, time.UTC)
	//进行时区的转换
	tsamp := t.Unix()
	tsamp += int64(n * 3600)

	return tsamp
}

//获取前多少天的时间
//n 前多少天
//return 时间戳
func GetNpreTime(n int) int64 {
	return time.Now().AddDate(0, 0, -n).Unix()
}

//获取前多少天的日期
//n 前多少天
//return 日期：yyyy-mm-dd HH:ii:ss
func GetNpreDate(n int) string {
	nytime := time.Now().AddDate(0, 0, -n).Unix()
	return Time2Date(nytime)
}

//获取前多少个小时的日期
//n 前多少小时
//return 日期：yyyy-mm-dd HH:ii:ss
func GetNpreHourDate(n int) string {
	t := time.Now()

	y := t.Year()
	m := t.Month()
	d := t.Day()

	h := t.Hour()
	mm := t.Minute()
	s := t.Second()

	//昨天的日期
	h = h - n

	tsamp := Date2Time(fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d",
		y, m, d, h, mm, s), -8)

	return Time2Date(tsamp)
}

//map类型转query
func Map2Query(data map[string]string) string {
	v := &url.Values{}
	for key, val := range data {
		v.Add(key, val)
	}
	return v.Encode()
}

//map类型转json
func Map2Json(data map[string]string) string {
	jrs, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jrs)
}

//获取随机字符串
//moreStr 更多的字符信息
//return 返回随机字符串
func GetRandStr(moreStr string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := fmt.Sprintf("%s_%d_%d", moreStr, int(time.Now().Unix()), r.Intn(10000))
	bstr, _ := Md5(str)
	return bstr
}

//map转xml
//data map数据
//return xml 字符串
func Map2Xml(data map[string]string) string {
	xml1 := ""
	for key, val := range data {
		xml1 += "<" + key + ">" + val + "</" + key + ">"
	}

	return xml1
}

//sha1加密
//str 待加密的字符串
//return (密文，错误信息)
func Sha1(str string) (string, error) {
	sha1er := sha1.New()
	n, err := sha1er.Write([]byte(str))
	if err != nil || n != len(str) {
		return "", err
	}
	signSlice := sha1er.Sum(nil)
	return fmt.Sprintf("%x", signSlice), nil
}

//sha2加密
func Sha2(str string) (string, error) {
	sha1er := sha256.New()
	n, err := sha1er.Write([]byte(str))
	if err != nil || n != len(str) {
		return "", err
	}
	signSlice := sha1er.Sum(nil)
	return fmt.Sprintf("%x", signSlice), nil
}

//md5加密
//str 待加密的字符串
//return (密文，错误信息)
func Md5(str string) (string, error) {
	md5er := md5.New()
	n, err := md5er.Write([]byte(str))
	if err != nil || n != len(str) {
		return "", err
	}
	signSlice := md5er.Sum(nil)
	return fmt.Sprintf("%x", signSlice), nil
}

//base64加密
//str　待加密的字符串
//return (密文,错误信息)
func Base64Encode(str string) (string, error) {
	B64Table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	coder := base64.NewEncoding(B64Table)
	estr := coder.EncodeToString([]byte(str))

	return estr, nil
}

//base64加密
//str　待加密的字符串
//return (密文,错误信息)
func Base64Decode(str string) (string, error) {
	B64Table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	coder := base64.NewEncoding(B64Table)
	dstr, err := coder.DecodeString(str)

	return string(dstr), err
}

//过滤 emoji 表情
//content 转入的字符串信息
func FilterEmoji(content string) string {
	new_content := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			new_content += string(value)
		}
	}
	return new_content
}

//读取csv文件
//fpath csv文件路径
func ReadCsv(fpath string) ([][]string, error) {
	//读取数据
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ','

	r_arr, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return r_arr, nil
}

//mysql安全过滤
func SafeSql(content string) string {
	//将'替换为'',\替换为\\,
	content = strings.Replace(content, `'`, `''`, -1)
	content = strings.Replace(content, `\`, `\\`, -1)

	return content
}

//字符串截取
//content 待截取的字符串
//slen 截取的长度
func SubStr(content string, slen int) string {
	if slen <= 0 {
		return ""
	}
	r := []rune(content)
	if len(r) > slen {
		return string(r[0:slen])
	} else {
		return content
	}
}

//文件上传
//r Request
//formname 表单ｆｉｌｅ类型的ｉｎｐｕｔ名称
//sfilepath 保存文件的路径
//return (保存的文件名，错误信息）
func Upload(r *http.Request, formname string, sfilepath string) (string, error) {
	//上传文件信息
	f, h, err := r.FormFile(formname)
	if err != nil {
		return "", errors.New("上传文件未选择！")
	}
	//获取大小的接口
	type Sizer interface {
		Size() int64
	}
	//获取文件信息的接口
	type Stat interface {
		Stat() (os.FileInfo, error)
	}
	//文件大小是否为0
	var usize int64 = 0
	if fs, ok := f.(Sizer); ok {
		usize = fs.Size()
	} else {
		if stat1, ok := f.(Stat); ok {
			ufinfo, err := stat1.Stat()
			if err != nil {
				return "", err
			}
			usize = ufinfo.Size()
		}
	}
	if usize <= 0 {
		return "", errors.New("文件大小不正确！")
	}

	//移动上传文件
	picname, _ := Md5(h.Filename + fmt.Sprint("%d", time.Now().Unix()))
	picext := filepath.Ext(h.Filename)
	picpath := picname + picext
	pf, err := os.Create(sfilepath + picpath)
	if err != nil {
		return "", errors.New("创建保存目标失败！")
	}
	nsize, err := io.Copy(pf, f)
	if err != nil {
		return "", errors.New("文件保存失败！")
	}
	if nsize < usize {
		return "", errors.New("文件未能上传成功！")
	}

	return picpath, nil
}

//文件下载
//w *http.ResponseWriter
//content 文件内容信息
//filename 下载的文件名称
func DownLoad(w *http.ResponseWriter, content string, filename string) {
	if filename == "" {
		filename = Time2Date(time.Now().Unix())
	}
	(*w).Header().Set("Content-type", "application/octet-stream")
	(*w).Header().Set("Accept-Ranges", "bytes")
	(*w).Header().Set("Accept-Length", fmt.Sprintf("%d", len(content)))
	(*w).Header().Set("Content-Disposition", "attachment;filename="+filename)
	fmt.Fprint(*w, content)
}
