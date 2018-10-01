package po

import (
	"path/filepath"
	"os"
	"log"
	"strings"
)

var PARENT_DIR_PATH string = getCurrentDirectory()
var ALGOLIA_CONFIG_YAML_PATH string = PARENT_DIR_PATH + "/config.yaml"
var COMPLIE_EXEC_PATH string = PARENT_DIR_PATH + "/compile"
var CONENT_DIR_PATH string = PARENT_DIR_PATH + "/content"
var ALGOLIA_COMPLIE_JSON_PATH string = PARENT_DIR_PATH + "/public/algolia.json"

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
