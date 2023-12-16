# echo-server

Real World HTTPを参考に作成

## Usage

実行

```
./run.sh
```

リクエスト

```
curl --http1.0 http://localhost:18888/greeting
```

ログ

```
2023/12/14 16:22:42 start http listening :18888
GET /greeting HTTP/1.0
Host: localhost:18888
Accept: */*
User-Agent: curl/7.68.0
```

## chromeでアクセスした場合

```
GET /greeting HTTP/1.1
Host: localhost:18888
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
Accept-Encoding: gzip, deflate, br
Accept-Language: ja-JP,ja;q=0.9,en-US;q=0.8,en;q=0.7
Connection: keep-alive
Sec-Ch-Ua: "Chromium";v="118", "Google Chrome";v="118", "Not=A?Brand";v="99"
Sec-Ch-Ua-Mobile: ?0
Sec-Ch-Ua-Platform: "Linux"
Sec-Fetch-Dest: document
Sec-Fetch-Mode: navigate
Sec-Fetch-Site: none
Sec-Fetch-User: ?1
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36


GET /favicon.ico HTTP/1.1
Host: localhost:18888
Accept: image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8
Accept-Encoding: gzip, deflate, br
Accept-Language: ja-JP,ja;q=0.9,en-US;q=0.8,en;q=0.7
Connection: keep-alive
Referer: http://localhost:18888/greeting
Sec-Ch-Ua: "Chromium";v="118", "Google Chrome";v="118", "Not=A?Brand";v="99"
Sec-Ch-Ua-Mobile: ?0
Sec-Ch-Ua-Platform: "Linux"
Sec-Fetch-Dest: image
Sec-Fetch-Mode: no-cors
Sec-Fetch-Site: same-origin
User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36
```

## curlでのurl encode送信

```
curl --http1.0 --get --data-urlencode "search word" http://localhost:18888
```

## curlで通信内容確認

```
curl -v --http1.0 http://localhost:18888/greeting
```

## curlでヘッダーを送信

```
curl -H "X-Test: Hello" http://localhost:18888
```
