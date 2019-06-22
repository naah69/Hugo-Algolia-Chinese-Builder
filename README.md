# Hugo-Algolia-Chinese-Builder

这是一个使用go语言jieba分词器编写的中文分词器


## 快速开始

### 最新更新

1. 完全使用go进行分词，摒弃node.js分词
2. 加入sego分词，使用双分词优化质量，同时优化分词速度
3. 加入缓存机制，每次通过md5比对文件，只对有变化的文件分词
4. 支持上传索引时使用http代理
5. 支持使用自定义分词字典，自定义停用词

### 1 下载
[release](https://github.com/naah69/Hugo-Algolia-Chinese-Builder/releases)页面内下载压缩包

解压压缩包到`hugo project`根目录中.


### 2 配置
修改`config.yaml`文件
```yaml
---
#配置你的api key
algolia:
  index: "your algolia index"
  key: "your algolia admin key"
  appID: "your algolia appID"

#字典和停用词位置
#使用data目录会让hugo命令报错
participles:
  dict:
    path: "sdata/dict.txt"
    stop-path: "sdata/stop.txt"

#配置http代理(可选)
#http:
 #  httpProxy: "127.0.0.1:1087"
---
```


### 4 改写你的compile脚本
修改compile中的命令为自己hugo的打包命令

### 运行
运行下面的命令
```bash
main
```
