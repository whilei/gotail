package main 

import (
	"testing"
	"os"
	"time"
	"strings"
)

func TestLineCount(t *testing.T){
	name := "字节数组行统计"
	lines := []byte("hello\nworld\n")
	if i := LineCount(lines); i != 2 {
		t.Error(name)
	}else{
		t.Log(name)
	}
}

func TestFileByteRead(t *testing.T){
	name := "获取文件大小"
	filename := "/tmp/test.log"
	content := []byte("abcdefghij")
	fo, _ := os.Create(filename)
	_, _ = fo.Write(content)
	fo.Close()
	fileSize := GetFileSize(filename)
	if fileSize != len(content){
		t.Error(name)
	}else{
		t.Log(name)
	}
}

func TestByteArrayToMultiLines(t *testing.T){
	name := "多个字节数组到多行字符串"
	a := []byte("happy ")
	b := []byte("spring\n")
	c := []byte("festival\n!")
	b = append(a, b...)
	d := append(b, c...)
	lines := ByteArrayToMultiLines(d)
	if len(lines) != 3{
		t.Error(name)
	} else {
		t.Log(name)
	}
}

func TestReadLastNLines(t *testing.T){
	name := "测试最后n行字符串读取"
	filename := "/tmp/test.log"
	TestString := "a\nb\nc\nd\ne\nf\ng\nh\ni\nj"
	LineCount := len(strings.Split(TestString, "\n"))
	content := []byte(TestString)
	fo, _ := os.Create(filename)
	_, _ = fo.Write(content)
	fo.Close()
	lines, err := ReadLastNLines(filename, LineCount)
	if err != nil || len(lines) != LineCount {
		t.Error(name)
		t.Error(lines)
		t.Error(len(lines), "!=", LineCount)
	} else {
		t.Log(name)
	}
}

func TestFileMonitor(t *testing.T){
	fh, err := os.Create("test.log")
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()
	ticker := time.Tick(1 * time.Second)
	for now := range ticker{
		fh.WriteString(now.String())
	}
}