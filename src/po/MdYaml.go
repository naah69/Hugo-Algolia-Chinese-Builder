package po

type MdYaml struct {
	Title       string   `yaml:"title"`
	Subtitle    string   `yaml:"subtitle"`
	Description string   `yaml:"description"`
	Keyword     string   `yaml:"keyword"`
	Date        string   `yaml:"date"`
	Author      string   `yaml:"author"`
	Tags        []string `yaml:"tags"`
	Image       string   `yaml:"image"`
	Categories  []string `yaml:"categories"`
	Draft       bool     `yaml:"draft"`
}
