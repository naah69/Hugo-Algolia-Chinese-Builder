package utils

import (
	"io/ioutil"
	"strings"
	"os"
	"fmt"
	"os/exec"
	"bytes"
)
func ReadFileString(path string)string {
	bytes, _ := ioutil.ReadFile(path)
	return string(bytes)
}

func WriteFile(path string,bytesArray []byte){
	ioutil.WriteFile(path, bytesArray, 0666)
}

func ReadMdContext(path string) (string,string) {
	str := ReadFileString(path)
	str = str[4 : len(str)-1]
	yaml:=str[0 : strings.Index(str, "---")]
	var context string
	if strings.Index(str, "---")+3 >= len(str)-1 {
		context=""
	}else{
		context = str[strings.Index(str, "---")+4 : len(str)-1]
	}

	return yaml,context
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		fmt.Println("exists error: " + path + " not found")
		return false, nil
	}
	return true, err
}

func ExecShell(s string) (string, error){
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()

	return out.String(), err
}

func GetAllFiles(path string,array *[]string)[]string{
	files, _ := ioutil.ReadDir(path)
	for _, fileInfo := range files {
		absoluteFilePath:=path +"/"+ fileInfo.Name()
		info, _ := os.Stat(absoluteFilePath)
		if info.IsDir() {
			GetAllFiles(absoluteFilePath,array)
		}
		*array=append(*array, absoluteFilePath )
	}
	return *array
}