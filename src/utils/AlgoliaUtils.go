package utils

import (
	"../po"
	"encoding/json"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//更新分词
func UpdateAlgolia(objects []algoliasearch.Object) bool {
	data, _ := ioutil.ReadFile(po.ALGOLIA_CONFIG_YAML_PATH)
	conf := po.ConfigYaml{}
	yaml.Unmarshal(data, &conf)
	client := algoliasearch.NewClient(conf.Algolia.AppID, conf.Algolia.Key)
	index := client.InitIndex(conf.Algolia.Index)
	index.Clear()
	index.AddObjects(objects)
	return true
}

func main() {
	content, _ := ioutil.ReadFile("/Users/naah/Documents/Hugo/Naah-Blog/public/algolia.json")

	var objects []algoliasearch.Object
	if err := json.Unmarshal(content, &objects); err != nil {
		return
	}
	UpdateAlgolia(objects)
}
