# alter-webhook
prometheus -> altermanager -> alter-webhook -> webhook

# 测试
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{     "receiver": "web.hook",     "status": "firing",     "alerts": [{         "status": "firing",         "labels": {             "alertname": "node",             "instance": "192.168.0.10",             "job": "test1111",             "severity": "critical"         },         "annotations": {             "description": "test1 192.168.0.10 节点断联已超过1分钟！",             "summary": "192.168.0.10 down "         },         "startsAt": "2022-04-28T08:44:23.05Z",         "endsAt": "0001-01-01T00:00:00Z",         "generatorURL": "http://localhost.localdomain:19090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=1",         "fingerprint": "726681bf4674e8a5"     }, {         "status": "firing",         "labels": {             "alertname": "node",             "instance": "192.168.1.10",             "job": "test2222",             "severity": "critical"         },         "annotations": {             "description": "test2 192.168.1.10 节点断联已超过1分钟！",             "summary": "192.168.1.10 down "         },         "startsAt": "2022-04-28T08:44:23.05Z",         "endsAt": "0001-01-01T00:00:00Z",         "generatorURL": "http://localhost.localdomain:19090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=1",         "fingerprint": "726681bf4674e8a5"     } ],     "groupLabels": {         "alertname": "node"     },     "commonLabels": {         "alertname": "node",         "instance": "192.168.0.10",         "job": "rh7",         "severity": "critical"     },     "commonAnnotations": {         "description": "rh7 192.168.0.10 节点断联已超过1分钟！",         "summary": "192.168.0.10 down "     },     "externalURL": "http://192.168.0.10:19092",     "version": "4",     "groupKey": "{}:{alertname=\"node\"}",     "truncatedAlerts": 0 }' \
 http://localhost:8080/v1/webhook
```

```
https://www.commands.dev/workflows/post_json_data_with_c_url
https://www.kandaoni.com/news/14421.html
https://yunlzheng.gitbook.io/prometheus-book/parti-prometheus-ji-chu/alert/alert-manager-use-receiver/alert-manager-extension-with-webhook
https://www.json.cn/
```