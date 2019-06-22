# Hugo-Algolia-Chinese-Builder

这是一个使用go语言jieba分词器编写的中文分词器


## 快速开始

### 最新更新

1. 使用go进行分词，摒弃node.js分词
2. 优化速度
3. 加入缓存机制，每次通过md5比对文件，只对有变化的文件分词

### 1 下载main
把`main` 和 `compile` 放在 hugo project根目录中.

compile为可执行脚本，放你hugo编译项目的命令


### 2 配置
创建config.yaml文件，并写入下面的配置
```yaml
---
algolia:
  index: "your algolia index"
  key: "your algolia admin key"
  appID: "your algolia appID"
#http:
 #  httpProxy: "127.0.0.1:1087"
---
```

### 4 改写你的compile脚本
修改compile中的命令为自己的命令

### 运行
运行下面的命令
```bash
main
```
