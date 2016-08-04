# picpicgo
go语言图片爬虫

![](./logo.jpg)

# Usage
```
$ ./picpic -h
Usage of picpic:
  -dir string
        自定义存放的路径 (default "./Downloads/")
  -fpic string
        自定义过滤图片链接的关键字
  -furl string
        自定义过滤网页链接的关键字
  -img string
        自定义图片属性名称，如data-original (default "src")
  -no int
        需要爬取的有效图片数量 (default 20)
  -re
        是否需要递归当前页面链接 (default true)
  -size int
        最小图片大小 单位kB (default 150)
  -url string
        起始网址
```

# TODO
- 完善反爬虫
- JS动态渲染
- 更好的日志输出
- web页面
- 国际化
