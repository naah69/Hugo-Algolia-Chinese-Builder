package utils
//
//import (
//	"builder/constant1"
//	"fmt"
//	"github.com/deckarep/golang-set"
//	"strings"
//)
//
//func main() {
//	//mergeDict()
//	mergeStop()
//}
//
//func mergeDict() {
//	path := "/Users/naah/software/GOPATH/src/github.com/go-ego/gse/data/dict/dictionary.txt,/Users/naah/software/GOPATH/src/github.com/go-ego/gse/data/dict/zh/dict.txt,/Users/naah/software/GOPATH/src/github.com/yanyiwu/gojieba/dict/jieba.dict.utf8"
//	split := strings.Split(path, ",")
//	set := mapset.NewSet()
//	for index1 := range split {
//		p := split[index1]
//		s := ReadFileString(p)
//		sa := strings.Split(s, "\r\n")
//		for index2 := range sa {
//			set.Add(sa[index2])
//			constant1.Num++
//		}
//
//	}
//	slice := set.ToSlice()
//	fmt.Print(len(slice))
//	array := InterfaceArray2StringArray(slice)
//	str := strings.Join(array, "\r\n")
//	WriteFile("1.dict", []byte(str))
//}
//func mergeStop() {
//	path := "/Users/naah/software/GOPATH/src/github.com/yanyiwu/gojieba/dict/stop_words.utf8,/Users/naah/Documents/Hugo/Naah-Blog/stop.txt"
//	split := strings.Split(path, ",")
//	set := mapset.NewSet()
//	for index1 := range split {
//		p := split[index1]
//		s := ReadFileString(p)
//		sa := strings.Split(s, "\r\n")
//		for index2 := range sa {
//			set.Add(sa[index2])
//			constant1.Num++
//		}
//
//	}
//	slice := set.ToSlice()
//	fmt.Print(len(slice))
//	array := InterfaceArray2StringArray(slice)
//	str := strings.Join(array, "\n")
//	WriteFile("stop1.txt", []byte(str))
//}
