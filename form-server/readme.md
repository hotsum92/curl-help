# form-server

form検証用

## Usage

実行

```
./run.sh
```

リクエスト

```
curl --http1.0 http://localhost:18888/greeting
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

[form-method-post.html](form-method-post.html)

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
<form method="POST" action="url" enctype="multipart/form-data">
    <input type="text" name="message" value="Hello"/><br>
    <input type="file" name="file"/><br>
    <input type="submit" value="SUBMIT"/>
</form>
```

ファイルは絶対パスで指定する必要がある

```
curl --http1.0 -F title="The Art of Community" -F author="Jono Bacon" -F attachment-file=@/path/test.txt http://localhost:18888
curl --http1.0 -F "attachment-file=@test.txt;type=text/html" http://localhost:18888
curl --http1.0 -F "attachment-file=@test.txt;filename=sample.txt" http://localhost:18888
```

## フォームを利用したリダイレクト

```
<!DOCTYPE html>
<html>
<body onload="document.forms[0].submit()">
<form action=" リダイレクトしたい先 " method="post">
<input type="hidden" name="data"
value=" 送りたいメッセージ />
<input type="submit" value="Continue"/>
</form>
</body>
```

(redirect.html)[redirect.html]

## foundを返してリダイレクト

foundを返して、リダイレクトするほうが、送信用と表示用でブックマークを分けたほうがいい

[found-form.html](found-form.html)
