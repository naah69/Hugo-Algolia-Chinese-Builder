package po

type ConfigYaml struct {
	Algolia struct {
		Index string `yaml:"index"`
		Key   string `yaml:"key"`
		AppID string `yaml:"appID"`
	}
}
