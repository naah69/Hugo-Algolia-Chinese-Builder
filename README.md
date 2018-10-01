# Hugo-Algolia-Chinese-Builder

这是一个使用go语言jieba分词器编写的中文分词器

它依赖hugo-algolia编译出的json文件格式

Detail to look [Art White](https://github.com/naah69/hugo-theme-artwhite)

## 快速开始

### 1 下载main
把`main` 和 `compile` 放在 hugo project中.

### 2 安装hugo-algolia
运行下面的命令:
```bash
$ npm install hugo-algolia -g
```

### 3 配置
创建config.yaml文件，并写入下面的配置
```yaml
---
baseurl: "your baseurl"
DefaultContentLanguage: "zh-cn"
hasCJKLanguage: true
languageCode: "zh-cn"
title: "your site title"
theme: "hugo-theme-artwhite"
metaDataFormat: "yaml"
algolia:
  index: "your algolia index"
  key: "your algolia admin key"
  appID: "your algolia appID"
---
```

### 4 改写你的compile脚本
修改compile中的主题

### 运行
运行下面的命令
```bash
main
```
