# 第四节课作业lv1

## 前排提示：

1、每次关闭postman之类软件页面之前请退出（“/quit”）“账号”。

2、初次使用是没有账号的，需要注册再登录才可使用。

3、使用与留言相关的功能需要注意请求中“comment”的单复数

## 功能介绍：

### 1、注册

在请求处输入“/register”

key需要username、password、check question（忘记密码设置的问题）、check answer（问题的答案）

### 2、登录

在请求处输入“/login”

key需要username、password

无法多个账户同时登录，换账号前需要退出已登录的账号

登录后才可以使用除了“初始化”、“注册”、“忘记密码”以外的功能

### 3、用令牌获取用户信息

在请求处输入“/user/get”

Header处需要Authorization

### 4、更改密码

在请求处输入“/change password”

key需要new password

更改后会自动退出账号，需要重新登录

### 5、忘记密码

#### 1.输入忘记密码的请求

在请求处输入“/forget password”

key需要username

会获得问题，在下一步回答

#### 2.回答问题获得改密码资格

在请求处输入“/answer”

key需要answer

#### 3.更改密码

在请求处输入“/update password”

key需要new password

密码就更改成功，可以登录了

### 6、留言

在请求处输入“/add comment”

key需要comment

### 7、浏览留言板

在请求处输入“/scan comments”

不需要任何key

### 8、删除留言

在请求处输入“/delete comment”

key需要num

num是用户想删除的留言的序号

### 9、清空留言板

在请求处输入“/clear comments”

不需要任何key

### 10、退出账号

在请求处输入“/quit”

不需要任何key

每次关闭请求之前一定要退出！！！！！！！

### 11、注销账号

在请求处输入“/unsubscribe”

不需要任何key

### 12、恢复初始化

在请求处输入“/clear all”

不需要任何key

使用该功能时不可以有账户已登录

## 该作业的数据库的表：
![T$8PWPO((773{USVH96@@AU](https://user-images.githubusercontent.com/116962163/205443266-ef366d40-242a-43a7-ac32-0b2faa8e5911.png)
![J5WHWT0OIXVZRWB UWCOWRF](https://user-images.githubusercontent.com/116962163/205443274-ab3b4e0f-fdaf-463c-b4be-b82ad42d72a2.png)
![RG K9S9YC$_5I%~O7$WXZMB](https://user-images.githubusercontent.com/116962163/205443280-2d884571-d00a-47b7-b3c1-012a7576780e.png)

## lv0的作业截图：
![JNNG233H8BDOK7X7ORON_LV](https://user-images.githubusercontent.com/116962163/205443301-b637e142-72de-46ae-8e7b-6fbc2187809d.png)

