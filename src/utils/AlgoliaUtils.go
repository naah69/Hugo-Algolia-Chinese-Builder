package utils

import (
	"../constant"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"net/http"
	"net/url"
)

//更新分词
func UpdateAlgolia(objects []algoliasearch.Object) bool {

	client := algoliasearch.NewClient(constant.AlgoliaCongig.Algolia.AppID, constant.AlgoliaCongig.Algolia.Key)
	if constant.AlgoliaCongig.Http.Proxy != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse("http://127.0.0.1:1087")
			//return url.Parse("ss://rc4-md5:123456@ss.server.com:1080")
		}
		tr := &http.Transport{Proxy: proxy}
		httpclient := &http.Client{
			Transport: tr,
		}
		client.SetHTTPClient(httpclient)

	}

	index := client.InitIndex(constant.AlgoliaCongig.Algolia.Index)
	index.Clear()
	index.AddObjects(objects)
	return true
}
