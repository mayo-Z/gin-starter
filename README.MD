# 描述
这是一个结构清晰的gin脚手架，有dao,dto,controller层。

此项目以gin框架为基础，构建了用户注册/登录/获取信息/退出/注销五个接口。以会员等级为例，构建了增删改查四个接口。

此项目可通过redis扩展分布式session， 路由中以中间件的形式构建了翻译器，验证器，自定义返回方法。

使用了viper读取配置文件，使用了gorm与数据库进行连接。

使用了泛型提取了dao层的增删改查接口。

# 使用方法
此代码用到了mysql,redis,swagger，需自行配置环境


1、用git将项目保存到本地

```
git clone https://github.com/uptocorrupt/gin-scaffold.git
```

2、进config文件夹里修改配置文件dev.yml.并将其重命名为application.yml

3、进项目根目录初始化swag，初始化mod。
```
go install github.com/swaggo/swag/cmd/swag@latest
```
```
swag init
```
```
go mod tidy
```
4、运行项目

5、在网页上输入host/swagger/index.html，测试和查看接口
如（[http://127.0.0.1:8090/swagger/index.html](http://127.0.0.1:8090/swagger/index.html)）
