# esxls

从 es 中查询结果，并输出到 excel 中，es 查询语句是特殊的。
需要将查询的结果的 json 文件放到 input 目录中。

```json
[{
    "key": "1.1.1.1",
    "doc_count": 20,
    "timestamp1": {
        "hits": {
            "total": 20,
            "max_score": null,
            "hits": [
                {
                    "_index": "conn_index_2020.05.05-000007",
                    "_type": "conn_doc",
                    "_id": "ABCDEFABCDEF-EzN",
                    "_score": null,
                    "_source": {
                        "ip_dst_addr": "8.8.8.8",
                        "enrichments:geo:ip_src_addr:province": "内网"
                    },
                    "sort": [
                        1588188888888
                    ]
                }
            ]
        }
    }
}]
```


```shell
# 查询语句
curl -XGET localhost:9200/conn_index_2020.05.05*/_search?pretty -d '
{"query":{"bool":{"filter":{"bool":{"must":[{"terms":{"ip_dst_addr":["8.8.8.8","8.8.8.9"]}}]}}}},"size":0,"aggs":{"srcip":{"terms":{"field":"ip_src_addr","size":500},"aggs":{"timestamp1":{"top_hits":{"sort":[{"timestamp":{"order":"asc"}}],"size":1,"_source":["enrichments:geo:ip_src_addr:province", "ip_dst_addr"]}}}}}}'
```