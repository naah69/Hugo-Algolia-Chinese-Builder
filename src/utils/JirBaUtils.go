package utils

import (
		"strings"
	"github.com/deckarep/golang-set"
			"github.com/yanyiwu/gojieba"
)

func JieBaParticiples(title string,context string) []string {
	jiebaParticiplesArray := jieBaParticiples(context)
	jiebaParticiplesSet := mapset.NewSet()
	for _, obj := range jiebaParticiplesArray {
		jiebaParticiplesSet.Add(obj)
	}
	jiebaParticiplesSet = removeWord(jiebaParticiplesSet)
	slice := jiebaParticiplesSet.ToSlice()
	array := InterfaceArray2StringArray(slice)
	return array
}

func jieBaParticiples(context string) []string {
	x := gojieba.NewJieba()
	defer x.Free()
	return x.CutForSearch(context, true)

}

//func Sego()  {
//	// 载入词典
//	var segmenter sego.Segmenter
//	segmenter.LoadDictionary("/Users/naah/software/GOPATH/src/github.com/huichen/sego/data/dictionary.txt")
//
//	// 分词
//	text := []byte("小明硕士毕业于中国科学院计算所，后在日本京都大学深造")
//	segments := segmenter.Segment(text)
//
//	// 处理分词结果
//	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
//	fmt.Println(sego.SegmentsToSlice(segments, false))
//
//}
func removeWord(wordSet mapset.Set) mapset.Set {
	str := "一,、,。,七,☆,〈,∈,〉,三,昉,《,》,「,」,『,』,‐,【,】,В,—,〔,―,∕,〕,‖,〖,〗,‘,’,“,”,〝,〞,!,\",•,#,$,%,&,…,',㈧,∧,(,),*,∪,+,,,-,.,/,︰,′,︳,″,︴,︵,︶,︷,︸,‹,︹,:,›,︺,;,︻,<,︼,=,︽,>,︾,?,︿,@,﹀,﹁,﹂,﹃,﹄,≈,义,﹉,﹊,﹋,﹌,﹍,﹎,﹏,﹐,﹑,﹔,﹕,﹖,[,\\,],九,﹝,^,﹞,_,﹟,也,`,﹠,①,﹡,②,﹢,③,④,﹤,⑤,⑥,﹦,⑦,⑧,﹨,⑨,﹩,⑩,﹪,﹫,|,白,~,二,五,¦,«,¯,±,´,·,¸,»,¿,ˇ,ˉ,ˊ,ˋ,×,四,˜,零,÷,─,！,＂,＃,℃,＄,％,＆,＇,（,）,＊,＋,，,－,．,／,０,１,２,３,４,５,６,７,８,９,：,会,；,＜,＝,＞,？,＠,Ａ,Ｂ,Ｃ,Ｄ,Ｅ,Ｆ,Ｇ,Ｉ,Ｌ,Ｒ,Ｔ,Ｘ,Ｚ,［,］,＿,ａ,ｂ,ｃ,ｄ,ｅ,ｆ,ｇ,ｈ,ｉ,ｊ,ｎ,ｏ,｛,｜,｝,～,Ⅲ,↑,→,Δ,■,Ψ,▲,β,γ,λ,μ,ξ,φ,￣,￥,\\n,},{, "
	split := strings.Split(str, ",")
	splitSet := mapset.NewSet()
	for index := range split {
		splitSet.Add(split[index])
	}
	return wordSet.Difference(splitSet)
}
