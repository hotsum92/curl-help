# certification-server

certification-server is a simple server that validate certificates.


## baseic auth

```
curl --http1.0 --basic -u user:pass http://localhost:18888
```

## digest auth

```
curl --http1.0 --digest -u user:pass http://localhost:18888/digest
```
