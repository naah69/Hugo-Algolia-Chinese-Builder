package po

type Article struct {
	Yaml        MdYaml
	Content     string
	Md5Value    string
	Path        string
	Participles *[]string
}
