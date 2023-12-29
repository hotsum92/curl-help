docker run --rm -it -v $PWD:/go/src -p 18888:18888 golang:1.21 bash -c "cd ./src && go run ./main.go"
