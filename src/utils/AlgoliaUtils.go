package utils

import (
		"../po"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"encoding/json"
)

func UpdateAlgolia(objects []algoliasearch.Object) bool {
	data, _ := ioutil.ReadFile(po.ALGOLIA_CONFIG_YAML_PATH)
	conf := po.ConfigYaml{}
	yaml.Unmarshal(data, &conf)
	client := algoliasearch.NewClient(conf.Algolia.AppID, conf.Algolia.Key)
	index := client.InitIndex(conf.Algolia.Index)
	index.Clear()
	index.AddObjects(objects)
	return true;
}



func main() {
	content, _ := ioutil.ReadFile("/Users/naah/Documents/Hugo/Naah-Blog/public/algolia.json")

	var objects []algoliasearch.Object
	if err := json.Unmarshal(content, &objects); err != nil {
		return
	}
	UpdateAlgolia(objects);
}
