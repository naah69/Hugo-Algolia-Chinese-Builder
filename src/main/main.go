package main

import (
	"../po"
	"../utils"
	"bytes"
	"fmt"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"gopkg.in/yaml.v2"
	"runtime"
	"src/github.com/json-iterator/go"
	"strconv"
	"strings"
	"time"
)

func main() {
	//打印作者信息
	printAuthorInfo()

	//判断配置文件是否存在
	existsConfigFile()

	startTime := time.Now().UnixNano() / 1e6

	//运行编译
	execCompile()

	participlesStartTime := time.Now().UnixNano() / 1e6

	var mdList = getMarkDownList()
	var articleList = getArticleList(mdList)

	//获取分词列表
	cacheAlgoliasList := getCacheAlgoliasList()
	var taskNum = 0
	var flag = true
	//有缓存时
	if len(cacheAlgoliasList) != 0 {
		exists, _ := utils.Exists(po.MD5_ALGOLIA_JSON_PATH)
		if exists {
			flag = false
			//有md5map
			utils.ExecShell(po.MD5_ALGOLIA_JSON_PATH)
			po.Md5Map = po.NewConcurrentMap(getMd5Map())

			for _, article := range articleList {
				sss := article
				title := sss.Yaml.Title
				value := po.Md5Map.GetValue(title)
				oldMd5 := ""
				if value != nil {
					oldMd5 = value.(string)
				}

				if strings.Compare(oldMd5, sss.Md5Value) == -1 && po.CacheAlgoliasMap[title].Content != "" {
					po.Queue.Push(sss)
					po.NeedArticleList = append(po.NeedArticleList, sss)

					taskNum++
				}
			}
		}
	}

	//没缓存时
	if flag {
		for _, article := range articleList {
			po.Queue.Push(article)
			po.NeedArticleList = append(po.NeedArticleList, article)
			taskNum++
		}

	}

	//创建WaitGroup（java中的countdown）
	po.WaitGroup.Add(taskNum)

	//设置cpu并行数
	runtime.GOMAXPROCS(runtime.NumCPU())

	//创建线程池
	pool := new(utils.ThreadPool)
	pool.Init(runtime.NumCPU(), taskNum)

	//循环添加任务
	for i := 0; i < taskNum; i++ {
		pool.AddTask(func() error {
			return ParticiplesAsynchronous()
		})
	}
	pool.Start()

	//主线程阻塞
	po.WaitGroup.Wait()
	pool.Stop()
	fmt.Println("participles success: " + strconv.FormatInt((time.Now().UnixNano()/1e6)-participlesStartTime, 10) + " ms")

	//创建分词

	algoliaStartTime := time.Now().UnixNano() / 1e6
	for _, value := range po.NeedArticleList {
		cacheAlgoliasList = append(cacheAlgoliasList, po.Algolia{Title: value.Yaml.Title})
	}

	var objArray = []algoliasearch.Object{}
	old := po.CONENT_DIR_PATH + "/"
	for _, algolias := range cacheAlgoliasList {
		title := algolias.Title

		value := po.ArticleMap.GetValue(title)
		var article *po.Article = nil
		if value != nil {
			article = value.(*po.Article)
		} else {
			fmt.Println(title)
			continue
		}
		po.Md5Map.AddData(title, article.Md5Value)

		mapObj := utils.Struct2Map(article.Yaml)
		if article.Participles != nil {
			participlesArray := *article.Participles
			var buffer bytes.Buffer
			for _, str := range participlesArray {
				buffer.WriteString(str)
				buffer.WriteString(" ")
			}
			join := buffer.String()
			mapObj["content"] = join
		} else {
			mapObj["content"] = algolias.Content

		}
		uri := strings.Replace(article.Path, old, "", 1)
		mapObj["objectID"] = uri
		mapObj["uri"] = uri

		objArray = append(objArray, mapObj)
	}
	fmt.Println("generate algolia index success: " + strconv.FormatInt((time.Now().UnixNano()/1e6)-algoliaStartTime, 10) + " ms")
	uploadStartTime := time.Now().UnixNano() / 1e6
	//更新分词
	utils.UpdateAlgolia(objArray)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	algoliaBytes, _ := json.Marshal(objArray)
	md5Bytes, _ := json.Marshal(po.Md5Map.GetData())
	utils.WriteFile(po.ALGOLIA_COMPLIE_JSON_PATH, algoliaBytes)
	utils.WriteFile(po.CACHE_ALGOLIA_JSON_PATH, algoliaBytes)
	utils.WriteFile(po.MD5_ALGOLIA_JSON_PATH, md5Bytes)

	fmt.Println("update success: " + strconv.FormatInt((time.Now().UnixNano()/1e6)-uploadStartTime, 10) + " ms")
	fmt.Println("total : " + strconv.FormatInt((time.Now().UnixNano()/1e6)-startTime, 10) + " ms")
}

func getArticleList(mdList []string) []*po.Article {
	var articleList []*po.Article
	taskNum := len(mdList)
	//创建WaitGroup（java中的countdown）
	po.WaitGroup.Add(taskNum)

	//设置cpu并行数
	runtime.GOMAXPROCS(runtime.NumCPU())

	//创建线程池
	pool := new(utils.ThreadPool)
	pool.Init(runtime.NumCPU(), taskNum)

	for _, path := range mdList {
		path1 := path
		pool.AddTask(func() error {
			mdYaml, context := utils.ReadMdContext(path1)
			mdConf := po.MdYaml{}
			err := yaml.Unmarshal([]byte(mdYaml), &mdConf)
			if err != nil {
				fmt.Println("generate error: " + path1)
				return nil
			}
			article := po.Article{Yaml: mdConf, Content: context, Md5Value: utils.Md5V(context), Path: path1}
			articleList = append(articleList, &article)
			po.ArticleMap.AddData(mdConf.Title, &article)
			po.WaitGroup.Done()
			return nil
		})

	}

	pool.Start()
	//主线程阻塞
	po.WaitGroup.Wait()
	pool.Stop()
	return articleList
}

//多线程分词
func ParticiplesAsynchronous() error {
	article := po.Queue.Pop().(*po.Article)
	context := article.Content
	mdConf := article.Yaml

	participles := utils.Participles(mdConf.Title, context)
	article.Participles = &participles
	fmt.Println("generate success: " + article.Path)
	po.WaitGroup.Done()
	return nil
}

//打印作者信息
func printAuthorInfo() {
	fmt.Println("     ___           ___           ___           ___")
	fmt.Println("    /\\__\\         /\\  \\         /\\  \\         /\\__\\")
	fmt.Println("   /::|  |       /::\\  \\       /::\\  \\       /:/  /")
	fmt.Println("  /:|:|  |      /:/\\:\\  \\     /:/\\:\\  \\     /:/__/")
	fmt.Println(" /:/|:|  |__   /::\\~\\:\\  \\   /::\\~\\:\\  \\   /::\\  \\ ___")
	fmt.Println("/:/ |:| /\\__\\ /:/\\:\\ \\:\\__\\ /:/\\:\\ \\:\\__\\ /:/\\:\\  /\\__\\")
	fmt.Println("\\/__|:|/:/  / \\/__\\:\\/:/  / \\/__\\:\\/:/  / \\/__\\:\\/:/  /")
	fmt.Println("    |:/:/  /       \\::/  /       \\::/  /       \\::/  /")
	fmt.Println("    |::/  /        /:/  /        /:/  /        /:/  /")
	fmt.Println("    /:/  /        /:/  /        /:/  /        /:/  /")
	fmt.Println("    \\/__/         \\/__/         \\/__/         \\/__/")

	fmt.Println("================ Welcome to use naah-algolia-builder ===============")
	fmt.Println()
	fmt.Println("==================== Blog: http://www.naah69.com ===================")
	fmt.Println()
}

//判断配置文件是否存在
func existsConfigFile() {
	fmt.Println("====================== check config file start =====================")
	result := true

	var res, _ = utils.Exists(po.PARENT_DIR_PATH)
	result = result && res
	res, _ = utils.Exists(po.ALGOLIA_CONFIG_YAML_PATH)
	result = result && res
	res, _ = utils.Exists(po.COMPLIE_EXEC_PATH)
	result = result && res
	res, _ = utils.Exists(po.CONENT_DIR_PATH)
	result = result && res
	if result {
		fmt.Println("check success: all file found")
	} else {
		fmt.Println("check error: please check these files that are not found")
	}
	fmt.Println("====================== check config file end =======================\n")
	if !result {
		panic("error exit")
	}
}

//执行编译
func execCompile() {
	out, _ := utils.ExecShell(po.COMPLIE_EXEC_PATH)
	fmt.Print(out)
}

//获取md列表
func getMarkDownList() []string {
	var mdPathArray []string
	var filePathArray []string

	filePathArray = utils.GetAllFiles(po.CONENT_DIR_PATH, &filePathArray)
	for _, path := range filePathArray {
		if strings.HasSuffix(path, ".md") {
			mdPathArray = append(mdPathArray, path)
		}
	}
	return mdPathArray
}

func getCacheAlgoliasList() []po.Algolia {
	var res, _ = utils.Exists(po.CACHE_ALGOLIA_JSON_PATH)
	cacheAlgiliasArray := []po.Algolia{}
	if res {
		jsonString := utils.ReadFileString(po.CACHE_ALGOLIA_JSON_PATH)
		cacheAlgiliasArray = getAlgiliasJsonArray(jsonString)
		for _, algolias := range cacheAlgiliasArray {
			po.CacheAlgoliasMap[algolias.Title] = algolias
		}
	}
	return cacheAlgiliasArray

}

func getAlgiliasJsonArray(jsonString string) []po.Algolia {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var array []po.Algolia
	json.Unmarshal([]byte(jsonString), &array)

	return array
}

func getMd5Map() map[string]interface{} {
	md5Json := utils.ReadFileString(po.MD5_ALGOLIA_JSON_PATH)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var md5Map map[string]interface{}
	json.Unmarshal([]byte(md5Json), &md5Map)
	return md5Map
}
