package po

type ConfigYaml struct {
	Algolia struct {
		Index string `yaml:"index"`
		Key   string `yaml:"key"`
		AppID string `yaml:"appID"`
	}

	Http struct {
		Proxy string `yaml:"httpProxy"`
	}

	Participles struct {
		Dict struct {
			Path     string `yaml:"path"`
			StopPath string `yaml:"stop-path"`
		}
	}
}
