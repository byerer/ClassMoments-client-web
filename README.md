# ClassMoments-client-web
班级圈速冲后端

## run
### wire
```shell
wire ./cmd
```

```shell
go run ./cmd/ClassMoments/main.go
```

```bash
docker run -d -p 3360:3306 --name ClassMoments -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=ClassMoments mysql:latest
```


## todo 
- token
- gorm自增主键 question