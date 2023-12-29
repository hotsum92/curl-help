# echo-server

helloを返すだけのサーバー

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

urlencodeされる
`-X GET -d ...`で強制的にボディにデータを入れることができる

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

## curlでメソッドを指定

```
curl --http1.0 -X post http://localhost:18888/greeting
```

## curlでjson送信

--data-ascii, -data, -dオプションは、データをエスケープせず、すでにエスケープされているとみなす
エスケープしたい場合は,--data-urlencodeオプションを使う

```
curl -d "{\"hello\": \"world\"}" -H "Content-Type: application/json" http://localhost:18888
```

## base64

```
str='%7B%22hello%22%3A%20%22world%22%7D'
printf '%b\n' "${str//%/\\x}"

{"hello": "world"}
```

```
printf '%b\n' "${test//%/\\x}" | base64
eyJoZWxsbyI6ICJ3b3JsZCJ9Cg==
```

## form送信

以下の方法では、&やスペースはそのまま送信されていまう
```
curl --http1.0 -d title="The Art of Community" -d author="Jono Bacon" http://localhost:18888
```

htmlでのform

```
<form method="POST">
    <input name="title">
    <input name="author">
    <input type="submit">
</form>
```

似たような変換では、--data-urlencodeがあるが、こちらはスペースを+ではなく、%20に変換する(RFC3986)

```
curl --http1.0 --data-urlencode title="Head First PHP & MySQL" --data-urlencode author="Lynn Beighley, Michael Morrison" http://localhost:18888
```

以下２つを、URLエンコードと呼ばれている

* RFC3986(パーセントエンコード)

```
title=Head%20First%20PHP%20%26%20MySQL&author=Lynn%20Beighley%2C%20Michael%20Morrison
```

* RFC1866

```
title=Head+First+PHP+%26+MySQL&author=Lynn+Beighley%2C+Michael+Morrison
```

ただし、上記は同じアルゴリズムで変換できるので問題になることはない

## multipart/form-data

マルチーパートフォーム形式でファイルを送信できる

```
<form action="POST" enctype="multipart/form-data">
</form>
```


## プロキシ

プロキシとの違いは、GETのあとにスキーマなどの情報が入っている

```
curl --http1.0 -x http://localhost:18888 -U user:pass http://example.com/helloworld
```

```
GET http://example.com/helloworld HTTP/1.0
Accept: */*
Proxy-Authorization: Basic dXNlcjpwYXNz
Proxy-Connection: Keep-Alive
User-Agent: curl/7.68.0
```

```
curl --http1.0 -U user:pass http://localhost:18888/hellowworld
```

```
GET /hellowworld HTTP/1.0
Host: localhost:18888
Accept: */*
User-Agent: curl/7.68.0
```

## cache

[example.com](http://example.com/)にアクセスしたときのキャッシュがある場合のエスポンス

etagと更新日時より
```
If-Modified-Since: Thu, 17 Oct 2019 07:18:26 GMT
If-None-Match: "3147526947+gzip"
```

304 Not Modifiedが返ってくる


```
HTTP/1.1 304 Not Modified
Age: 31161
Cache-Control: max-age=604800
Date: Fri, 29 Dec 2023 19:05:36 GMT
Etag: "3147526947+gzip"
Expires: Fri, 05 Jan 2024 19:05:36 GMT
Last-Modified: Thu, 17 Oct 2019 07:18:26 GMT
Server: ECS (sac/2552)
Vary: Accept-Encoding
X-Cache: HIT
```

```
GET / HTTP/1.1
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
Accept-Encoding: gzip, deflate
Accept-Language: ja-JP,ja;q=0.9,en-US;q=0.8,en;q=0.7
Cache-Control: max-age=0
Connection: keep-alive
Host: example.com
If-Modified-Since: Thu, 17 Oct 2019 07:18:26 GMT
If-None-Match: "3147526947+gzip"
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36
```
