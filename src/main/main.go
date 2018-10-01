package main

import (
	"fmt"
	"../utils"
	"../po"
	"strings"
	"encoding/json"
	"time"
	"strconv"
	"gopkg.in/yaml.v2"
		"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

func main() {
	printAuthorInfo()
	existsConfigFile()
	fmt.Println("==================== update algolia index start ====================")
	startTime := time.Now().Unix()
	execCompile()
	existsAlgoliaCompileResultJson()
	fmt.Println("compile success: " + strconv.FormatInt(time.Now().Unix()-startTime, 10) + " sec");
	startTime = time.Now().Unix()
	mdList := getMarkDownList()
	algoliasList := getAlgoliasList()
	po.WaitGroup.Add(len(mdList))
	for _, path := range mdList {
		po.Queue.Push(path)
		go ParticiplesAsynchronous()
	}
	po.WaitGroup.Wait()
	fmt.Println("participles success: " + strconv.FormatInt(time.Now().Unix()-startTime, 10) + " sec");
	var objArray=[]algoliasearch.Object{}
	for _,algolias:=range algoliasList{
		participles, _ := po.ParticiplesMap.Load(algolias.Title)
		participlesArray := participles.([]string)
		mapObj := utils.Struct2Map(algolias)
		mapObj["objectID"]=mapObj["objectid"]
		mapObj["content"]=(strings.Join(participlesArray, " ")+" "+mapObj["content"].(string))
		objArray=append(objArray,mapObj)
	}
	utils.UpdateAlgolia(objArray)
	bytes, _ := json.Marshal(objArray)
	utils.WriteFile(po.ALGOLIA_COMPLIE_JSON_PATH,bytes)
	fmt.Println("update success: " + strconv.FormatInt(time.Now().Unix()-startTime, 10) + " sec");
	fmt.Println("==================== update algolia index end ====================");
}

func ParticiplesAsynchronous() error {
	path := po.Queue.Pop().(string)
	mdYaml, context := utils.ReadMdContext(path)
	mdConf := po.MdYaml{}
	err := yaml.Unmarshal([]byte(mdYaml), &mdConf)
	if err != nil {
		fmt.Println("generate error: " + path)
		return err
	}
	participles := utils.JieBaParticiples(mdConf.Title,context)
	po.ParticiplesMap.Store(mdConf.Title, participles)
	fmt.Println("generate success: " + path)
	po.WaitGroup.Done()
	return nil
}

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

func existsConfigFile() {
	fmt.Println("====================== check config file start =====================");
	result := true;
	var res, _ = utils.Exists(po.PARENT_DIR_PATH)
	result = result && res
	res, _ = utils.Exists(po.ALGOLIA_CONFIG_YAML_PATH)
	result = result && res
	res, _ = utils.Exists(po.COMPLIE_EXEC_PATH)
	result = result && res
	res, _ = utils.Exists(po.CONENT_DIR_PATH)
	result = result && res
	if result {
		fmt.Println("check success: all file found");
	} else {
		fmt.Println("check error: please check these files that are not found");
	}
	fmt.Println("====================== check config file end =======================\n");
	if !result {
		panic("error exit")
	}
}

func execCompile() {
	out, _ := utils.ExecShell(po.COMPLIE_EXEC_PATH)
	fmt.Print(out)
}

func existsAlgoliaCompileResultJson() {
	var res, _ = utils.Exists(po.ALGOLIA_COMPLIE_JSON_PATH)
	if !res {
		panic("error exit")
	}
}

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

func getAlgoliasList() []po.Algolia {
	jsonString := utils.ReadFileString(po.ALGOLIA_COMPLIE_JSON_PATH)

	var array []po.Algolia
	json.Unmarshal([]byte(jsonString), &array);
	for _,algolias:=range array{
		po.AlgoliasMap[algolias.Title]=algolias
	}
	return array

}
