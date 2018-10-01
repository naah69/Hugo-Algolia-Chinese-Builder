package po
type ConfigYaml struct {
	Baseurl                string `yaml:"baseurl"`
	DefaultContentLanguage string `yaml:"DefaultContentLanguage"`
	HasCJKLanguage         string `yaml:"hasCJKLanguage"`
	LanguageCode           string `yaml:"languageCode"`
	Title                  string `yaml:"title"`
	Theme                  string `yaml:"theme"`
	MetaDataFormat         string `yaml:"metaDataFormat"`
	Algolia                struct {
		Index string `yaml:"index"`
		Key   string `yaml:"key"`
		AppID string `yaml:"appID"`
	}
}
