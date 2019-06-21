package po

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

//修改回来
//var PARENT_DIR_PATH string = getCurrentDirectory()
var PARENT_DIR_PATH string = "/Users/naah/Documents/Hugo/Naah-Blog"
var ALGOLIA_CONFIG_YAML_PATH string = PARENT_DIR_PATH + "/config.yaml"
var COMPLIE_EXEC_PATH string = PARENT_DIR_PATH + "/compile"
var CONENT_DIR_PATH string = PARENT_DIR_PATH + "/content"
var ALGOLIA_COMPLIE_JSON_PATH string = PARENT_DIR_PATH + "/public/algolia.json"
var CACHE_ALGOLIA_JSON_PATH string = PARENT_DIR_PATH + "/cache_algolia.json"
var MD5_ALGOLIA_JSON_PATH string = PARENT_DIR_PATH + "/md5_algolia.json"

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
