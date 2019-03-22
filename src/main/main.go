package main

import (
	"../po"
	"../utils"
	"encoding/json"
	"fmt"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"gopkg.in/yaml.v2"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	//打印作者信息
	printAuthorInfo()

	//判断配置文件是否存在
	existsConfigFile()

	fmt.Println("==================== update algolia index start ====================")
	startTime := time.Now().Unix()

	//运行编译
	execCompile()

	//判断是否存在algolia的json
	existsAlgoliaCompileResultJson()
	fmt.Println("compile success: " + strconv.FormatInt(time.Now().Unix()-startTime, 10) + " sec")
	startTime = time.Now().Unix()

	//获取md列表
	mdList := getMarkDownList()

	//获取分词列表
	algoliasList := getAlgoliasList()

	//创建WaitGroup（java中的countdown）
	po.WaitGroup.Add(len(mdList))

	//设置cpu并行数
	runtime.GOMAXPROCS(runtime.NumCPU())

	//创建线程池
	pool := new(utils.ThreadPool)
	pool.Init(runtime.NumCPU(), len(mdList))

	//循环添加任务
	for _, path := range mdList {
		po.Queue.Push(path)
		pool.AddTask(func() error {
			return ParticiplesAsynchronous()
		})
	}
	pool.Start()

	//主线程阻塞
	po.WaitGroup.Wait()
	pool.Stop()
	fmt.Println("participles success: " + strconv.FormatInt(time.Now().Unix()-startTime, 10) + " sec")

	//创建分词
	var objArray = []algoliasearch.Object{}
	for _, algolias := range algoliasList {
		participles, _ := po.ParticiplesMap.Load(algolias.Title)
		participlesArray := participles.([]string)
		mapObj := utils.Struct2Map(algolias)
		mapObj["objectID"] = mapObj["objectid"]
		mapObj["content"] = (strings.Join(participlesArray, " ") + " " + mapObj["content"].(string))
		objArray = append(objArray, mapObj)
	}

	//更新分词
	utils.UpdateAlgolia(objArray)
	bytes, _ := json.Marshal(objArray)
	utils.WriteFile(po.ALGOLIA_COMPLIE_JSON_PATH, bytes)

	fmt.Println("update success: " + strconv.FormatInt(time.Now().Unix()-startTime, 10) + " sec")
	fmt.Println("==================== update algolia index end ====================")
}

//多线程分词
func ParticiplesAsynchronous() error {
	path := po.Queue.Pop().(string)
	mdYaml, context := utils.ReadMdContext(path)
	mdConf := po.MdYaml{}
	err := yaml.Unmarshal([]byte(mdYaml), &mdConf)
	if err != nil {
		fmt.Println("generate error: " + path)
		return err
	}
	participles := utils.JieBaParticiples(mdConf.Title, context)
	po.ParticiplesMap.Store(mdConf.Title, participles)
	fmt.Println("generate success: " + path)
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

//判断algolia的json是否存在
func existsAlgoliaCompileResultJson() {
	var res, _ = utils.Exists(po.ALGOLIA_COMPLIE_JSON_PATH)
	if !res {
		panic("error exit")
	}
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

//获取algolia列表
func getAlgoliasList() []po.Algolia {
	jsonString := utils.ReadFileString(po.ALGOLIA_COMPLIE_JSON_PATH)

	var array []po.Algolia
	json.Unmarshal([]byte(jsonString), &array)
	for _, algolias := range array {
		po.AlgoliasMap[algolias.Title] = algolias
	}
	return array

}
