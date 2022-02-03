# go-blog
Golang实现的blog博客项目

框架:gin+gorm

1. 数据库的设计

(1) 数据库配置:使用viper组件读取yml文件配置，gorm连接MySQL


(2) 数据表 数据模型

User: name, password, role

Article: category, title, cid, description, content  (cid是外键)

Category:name, ID


2. 统一封装处理模块

errormessage:  code ---> string用于错误时返回对应代码

response:  封装好的json返回格式



3. 各个功能接口 (RESTFUL API)


(1)用户接口JSON绑定CRUD

用户注册:密码加密bcrypt 用户名不能重复

用户登录: token


AddUser: POST 增加用户密码使用bcrypt加密存入数据库

GetUsers: GET查询用户列表考虑到分页功能需要query中增加pagesize和pagenum参数

EditUser: PUT 编辑用户信息(不包括密码)密码修改需要单独api

DeleteUser: DELETE 删除用户根据id删除读取path中的参数(软删除 并非完全删除数据库信息)




参考：

https://www.bilibili.com/video/BV1xL411V7Gw

https://www.bilibili.com/video/BV1ER4y1g7SR

https://blog.csdn.net/qq_44108469/category_11512073.html