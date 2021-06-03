## 1. 创建psotgresql
`
CREATE USER abc PASSWORD 'abc';
create database abc with owner abc;
`
## 2.1 server端
### 2.1 启动参数(缺省情况下程序会获取config.yaml里的数据库配置)

```
go run main.go --ip "127.0.0.1" --port 5432 --db "abc" --username "abc" --password "abc"
```

## 2 web端

```bash
# clone the project
git clone https://github.com/piexlmax/gin-vue-admin.git

# enter the project directory
cd web

# install dependency
npm install

# develop
npm run serve
```
