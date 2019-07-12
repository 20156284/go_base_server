# goCRM
golang 语言开发的CRM 系统

1、安装  iris 框架 由Will 开发的所有 go web 项目都是基于iris 框架进行开发
     安装命令如下 (安装的前提是你已经安装了 go sdk)
   **go get githu.com/kataras/iris** 

2、安装xrom go 框架
   **go get github.com/go-xorm/xorm**

3、安装数据库驱动
   目前xrom 框架支持一下驱动
   mysql、mymysql、postgres、tidb、sqlite、mssql、oracle(测试）
   **go get github.com/go-sql-driver/mysql**

4、如下是项目框架搭建后的说明：
* config：项目配置文件及读取配置文件的相关功能
  config.json：项目配置文件。
* controller：控制器目录、项目各个模块的控制器及业务逻辑处理的所在目录
* datasource：实现mysql连接和操作、封装操作mysql数据库的目录。
* model：数据实体目录，主要是项目中各业务模块的实体对象的定义
* service：服务层目录。用于各个模块的基础功能接口定义及实现，是各个模块的数据层。
* static：配置项目的静态资源目录。
* util：提供通用的方法封装。
* main.go：项目程序主入口
