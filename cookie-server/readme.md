# cookie-server

This is a server that verifies cookies.

## Usage

-c option is to save cookies to a file.
-b option is to read cookies from a file and add cookies to the request.

```
curl --http1.0 -c cookie.txt -b cookie.txt -b "name1=value1" http://localhost:18888/cookie
```

## javascript

you can get cookies with document.cookie.

```
> console.log(document.cookie);
"_ga=GA1.2....; c_user=100002291...; csm=2; p = 02; act=147--2358...;..."
```
