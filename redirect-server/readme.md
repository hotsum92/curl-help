# curl-help

Real World HTTPを参考に作成

## Usage

実行

```
./run.sh
```

## curlリダイレクト検証

```
curl -v -L http://localhost:18888/old-place

*   Trying 127.0.0.1:18888...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 18888 (#0)
> GET /old-place HTTP/1.1
> Host: localhost:18888
> User-Agent: curl/7.68.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 301 Moved Permanently
< Location: /new-place
< Date: Sat, 16 Dec 2023 14:41:04 GMT
< Content-Length: 36
< Content-Type: text/html; charset=utf-8
< 
* Ignoring the response-body
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:18888/new-place'
* Found bundle for host localhost: 0x55dbf3339050 [serially]
* Can not multiplex, even if we wanted to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (127.0.0.1) port 18888 (#0)
> GET /new-place HTTP/1.1
> Host: localhost:18888
> User-Agent: curl/7.68.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Sat, 16 Dec 2023 14:41:04 GMT
< Content-Length: 36
< Content-Type: text/html; charset=utf-8
< 
<html><body>new place</body></html>
* Connection #0 to host localhost left intact
```


