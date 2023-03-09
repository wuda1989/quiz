
# quiz 測驗

<img src="https://img.shields.io/badge/golang-1.19.3-blue" alt=""/>
<img src="https://img.shields.io/badge/gin-1.9.0-lightBlue" alt=""/>
<img src="https://img.shields.io/badge/gorm-1.24.6-red" alt=""/>

## 1. Run docker-compose

```
# 進入 quiz 資料夾
cd quiz

# 執行
docker-compose up -d server mysql
```

## 2. Comment API
### create a new comment.
```
curl --location '127.0.0.1:8888/quiz/v1/comment' \
--header 'Content-Type: application/json' \
--data '{
"uuid": "",
"parentid": "a1205dab-824a-4e3a-bcd2-ed6102e60ae9",
"comment": "根據中央氣象局地震測報中心地震報告這起規模....",
"author": "氣象局網站",
"update": "2022-06-01T01:12:33Z",
"favorite": false

}'
```

### get comment by uuid
```
curl --location '127.0.0.1:8888/quiz/v1/comment/2852872f-815d-41b7-91c6-317662756c48'
```

### modify comment by uuid.
```
curl --location --request PUT '127.0.0.1:8888/quiz/v1/comment/e5f107b7-dedb-45b1-be01-36e5c9b39d51' \
--header 'Content-Type: application/json' \
--data '{
  "uuid": "e5f107b7-dedb-45b1-be01-36e5c9b39d51",
  "parentid": "a1205dab-824a-4e3a-bcd2-ed6102e60ae9",
  "comment": "根據中央氣象局地震測報中心地震報告，這起規模...",
  "author": "氣象局網站",
  "update": "2022-06-01T01:12:33Z",
  "favorite": false
  
}'
```

### remove comment by uuid.
```
curl --location --request DELETE '127.0.0.1:8888/quiz/v1/comment/29f0c41d-86a5-47d4-9735-44b3f3c01ea1'
```