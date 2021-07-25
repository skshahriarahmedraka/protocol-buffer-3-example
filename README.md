
## python3
```
protoc -I=proto --python_out=python  proto/simple.proto
```
here -I = input directory , --python_out = output directory , proto file without argument 

## java
```
protoc -I=proto --java_out=java  proto/simple.proto
```
here -I = input directory , --java_out = output directory , proto file without argument 

## js es5
```
protoc --js_out=import_style=commonjs,binary:. employees.proto
```
```
npm install google-protobuf
```
### install protoc compiler

https://github.com/google/protobuf/releases
 
### Make sure you grab the latest version
```
curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
```
### Unzip
```
unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
```
### Move protoc to /usr/local/bin/
```
sudo mv protoc3/bin/* /usr/local/bin/
```
### Move protoc3/include to /usr/local/include/
```
sudo mv protoc3/include/* /usr/local/include/
```
### Optional: change owner
```
sudo chown [user] /usr/local/bin/protoc
sudo chown -R [user] /usr/local/include/google
```

## for python3 support
```
sudo pip3 install protobuf
```

## for Golang support 
```
 sudo apt install protobuf-compiler
 ```
 ```
 sudo apt install golang-goprotobuf-dev
```
```
 go get github.com/golang/protobuf
 ```
 ```
 go get github.com/golang/protobuf/proto
```
```
 export PATH=$PATH:$GOPATH/bin
 ```
 example :
 ```
 protoc --go_out=. *.proto
 ```
 ```
 protoc -I=src/ --go_out=plugins=grpc:src/ src/simple/simple.proto
 ```
 ```
 protoc -I=src/ --go_out=src/ src/simple/simple.proto
```

