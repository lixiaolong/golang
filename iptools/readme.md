# iptools

批量查询 IP 地址归宿地

```
# 参考来源
https://gitee.com/lionsoul/ip2region/

# golang lib
github.com/lionsoul2014/ip2region/binding/golang/ip2region

# db下载
https://gitee.com/lionsoul/ip2region/tree/master/data
```

使用方法

```
Usage:
  -i string
        input file name(all.txt | all.json | input/ | ip1;ip2) (default "input")
  -o string
        output file name (default "output.xlsx")

# txt，一行一个 IP 地址
iptools.exe -i all.txt

# json 格式
iptools.exe -i all.json

# json 格式
[
    {
        "key": "8.8.8.8",
        "doc_count": 20
    },
    {
        "key": "1.1.1.1",
        "doc_count": 9
    }
]

# 目录，里面放的是 json
iptools.exe -i input/

# 直接查询，结果直接输出
iptools.exe -i "8.8.8.8;1.1.1.1"
```