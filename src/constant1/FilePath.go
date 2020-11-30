package constant1

import (
	"builder/po"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//修改回来
var PARENT_DIR_PATH string = GetCurrentPath()

//var PARENT_DIR_PATH string = "/Users/naah/Documents/Hugo/Naah-Blog"
var ALGOLIA_CONFIG_YAML_PATH string = PARENT_DIR_PATH + "/config.yaml"
var COMPLIE_EXEC_PATH string = PARENT_DIR_PATH + "/compile"
var CONENT_DIR_PATH string = PARENT_DIR_PATH + "/content"
var ALGOLIA_COMPLIE_JSON_PATH string = PARENT_DIR_PATH + "/public/algolia.json"
var CACHE_ALGOLIA_JSON_PATH string = PARENT_DIR_PATH + "/cache_algolia.json"
var MD5_ALGOLIA_JSON_PATH string = PARENT_DIR_PATH + "/md5_algolia.json"
var AlgoliaCongig = po.ConfigYaml{}
var Num int32 = 0

func init() {
	fmt.Println("current path:" + GetCurrentPath())
	data, _ := ioutil.ReadFile(ALGOLIA_CONFIG_YAML_PATH)
	yaml.Unmarshal(data, &AlgoliaCongig)
}

//func GetCurrentFilePath() string {
////	_, filePath, _, _ := runtime.Caller(1)
////	return filePath
////}

func GetCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
