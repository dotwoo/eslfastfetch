# eslfastgo
eslfast 音频文章下载

![](./logo.jpg)

# Usage
```
$ ./eslfastgo -h
Usage of ./eslfastgo:
  -dir string
    	自定义存放的路径 (default "./Downloads/")
  -filter string
    	需要过滤掉的 url 关键词,使用/分隔 (default "dict/sent/cloze/w1/w2/w3/kewords")
  -l string
    	set the log level (default "info")
  -mp3 string
    	自定义mp3引用属性名称，如data-original (default "src")
  -no int
    	需要爬取的有效mp3数量 (default 20)
  -re
    	是否需要递归当前页面链接 (default true)
  -size int
    	最小Mp3大小 单位kB (default 150)
  -smp3 string
    	自定义Mp3链接的关键字 (default "ke")
  -sparent string
    	自定义过滤Mp3父页面链接须包含的关键字 (default "kidsenglish")
  -surl string
    	自定义网页链接的关键字 (default "kidsenglish")
  -url string
    	起始网址 (default "https://www.eslfast.com/kidsenglish/")
  -v	show the version
```

# TODO
- 完善反爬虫
- JS动态渲染
- 更好的日志输出
- web页面
- 国际化
