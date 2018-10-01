# Hugo-Algolia-Chinese-Builder

This is Chinese index builder with Algolia For Hugo.

It dependen hugo-algolia.

Detail to look [Art White](https://github.com/naah69/hugo-theme-artwhite)

## Quick Start

### 1 download main
put the `main` and `compile` in the hugo project.

### 2 install hugo-algolia
run the following command:
```bash
$ npm install hugo-algolia -g
```

### 3 configuration
create new file named config.yaml, and add the following:
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

### 4 modify your 'compile'
modify the theme in 'compile'

### run
run the following command:
```bash
main
```
