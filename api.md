# login
## `POST` `/login`
### `application/x-www-form-urlencoded`  
|请求参数 | 类型 | 备注 |
|--------|-----|------|
username|必选|用户名|
password|必选|密码|

|返回参数|说明|
|------|---|
|date|返回消息|
|token|用户token|

|date|说明|
|---|---|
|"info":服务器错误|服务器错误|
|“密码错误”|`password`与`username`不匹配|
|”用户不存在“|`username`不存在|
|用户token|登录成功|

# Register
## `POST` `/register`
### `application/x-www-form-urlencoded`
|请求参数 | 类型 | 备注 |
|--------|-----|------|
username|必选|用户名|
password|必选|密码|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|"用户名格式有误(min=4,max-10)"|`username`格式有误|
|"密码格式有误(min=6,max=16)"|`password`格式有误|
|”服务器错误“|服务器错误|
|”用户名已经存在“|`username`已存在|
|“注册成功”|注册成功|


#User
##`POST` `/user/password`
### `application/x-www-form-urlencoded` `headers`
#### 修改密码
|请求参数|类型|备注|
|---|---|---|
|oldPassword|必选|旧密码|
|newPassword|必选|新密码|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|"服务器错误"|服务器错误|
|”旧密码错误“|`oldPassword`与数据库中用户的`password`不匹配|
|”修改失败“|修改失败|
|“修改成功”|修改成功|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|

##`POST` `/user/avatar`
### `multipart/form-data` `headers`
#### 用户上传头像
|请求参数|类型|备注|
|---|---|---|
|avatar|必选|头像文件|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|
|”参数解析失败“|上传文件出错|
|”文件过大“|`avatar`过大|
|”文件格式错误“|`avatar`类型有误|
|”保存头像失败“|`avatar`保存到本地失败|
|”上传失败“|上传失败|
|“上传成功”|上传成功|

## `POST` `/user/:username/introduction`
### `application/x-www-form-urlencoded` `headers`
#### 上传个人介绍
|请求参数|类型|备注|
|---|---|---|
|introduction|必选|自我介绍|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|
|“服务器错误”|服务器出错|
|“成功”|上传成功|

## `GET` `/user/:username/avatar`
### 获取用户头像
|请求参数|类型|备注|
|---|---|---|
|:username|必选|用户名|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错|
|""loadString":user.Url,"Address":user.Address,"|成功返回|
## `GET` `/user/:username/introduction`
### 获取用户个人介绍
|请求参数|类型|备注|
|---|---|---|
|:username|必选|用户名|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错|
|user.Introduction|成功返回|

## `GET` `/user/:username/wantSee`
### 获取用户想看
|请求参数|类型|备注|
|---|---|---|
|:username|必选|用户名|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错|
|wants|成功返回|

## `GET` `/user/:username/seen`
### 获取用户看过
|请求参数|类型|备注|
|---|---|---|
|:username|必选|用户名|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错|
|seens|成功返回|

## `GET` `/user/:username/comment`
### 获取用户短评
|请求参数|类型|备注|
|---|---|---|
|:username|必选|用户名|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错|
|comments|成功返回|

## `GET` `/user/:username/longComment`
### 获取用户影评
|请求参数|类型|备注|
|---|---|---|
|:username|必选|用户名|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错|
|comments|成功返回|

#Movie
## `POST` `/movie/:movieId/wantSee`
###`application/x-www-form-urlencoded` `headers`
#### 添加想看
|请求参数|类型|备注|
|---|---|---|
|:movieId|必选|电影id|
|label|必选|电影标签(用户自定义)|
|comment|必选|对电影的评价|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错,删除成功|
|“数据库中查询不到该电影”|数据库信息不足|
|“成功”|添加成功|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|

## `DELETE` `/movie/:movieId/wantSee` 
`headers`
### 删除想看
|请求参数|类型|备注|
|---|---|---|
|:movieId|必选|电影id|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错,删除失败|
|“数据库中查询不到该电影”|数据库信息不足|
|“成功”|删除成功|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|

## `POST` `/movie/:movieId/seen`
###`application/x-www-form-urlencoded` `headers`
#### 添加看过
|请求参数|类型|备注|
|---|---|---|
|:movieId|必选|电影id|
|label|必选|电影标签(用户自定义)|
|comment|必选|对电影的评价|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错,添加失败|
|“数据库中查询不到该电影”|数据库信息不足|
|“成功”|添加成功|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|

## `DELETE` `/movie/:movieId/seen`
`headers`
### 删除看过
|请求参数|类型|备注|
|---|---|---|
|:movieId|必选|电影id|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错,删除失败|
|“数据库中查询不到该电影”|数据库信息不足|
|“成功”|删除成功|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|

## `POST` `/movie/:movieId/comment`
###`application/x-www-form-urlencoded` `headers`
#### 添加短评
|请求参数|类型|备注|
|---|---|---|
|:movieId|必选|电影id|
|content|必选|短评内容|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错,评论失败|
|“数据库中查询不到该电影”|数据库信息不足|
|“上传成功”|评论成功|
|“上传失败”|服务器出错，上传失败|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|

## `DELETE` `/movie/:movieId/comment` `headers`
### 删除短评
|请求参数|类型|备注|
|---|---|---|
|:movieId|必选|电影id|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错|
|“数据库中查询不到该电影”|数据库信息不足|
|“成功”|删除成功|
|“删除失败”|删除失败|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|
## `POST` `/movie/:movieId/longComment`
###`application/x-www-form-urlencoded` `headers`
#### 添加影评
|请求参数|类型|备注|
|---|---|---|
|:movieId|必选|电影id|
|content|必选|影评内容|
|title|必选|影评标题|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错,评论失败|
|“数据库中查询不到该电影”|数据库信息不足|
|“上传成功”|评论成功|
|“上传失败”|服务器出错，上传失败|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|

## `DELETE` `/movie/:movieId/longComment` `headers`
### 删除短评
|请求参数|类型|备注|
|---|---|---|
|:movieId|必选|电影id|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错|
|“数据库中查询不到该电影”|数据库信息不足|
|“成功”|删除成功|
|“删除失败”|删除失败|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|

# MovieDis
## `POST` `/discussion/:movieId/`
####`application/x-www-form-urlencoded` `headers`
#### 发布话题
|请求参数|类型|备注|
|---|---|---|
|title|必选|话题标题|
|content|必选|话题内容|
|:movieId|必选|电影id|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错,发布失败|
|“数据库中查询不到该电影”|数据库信息不足|
|“上传成功”|发布成功|
|“上传失败”|服务器出错，发布失败|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|

## `DELETE` `/discussion/:movieId/` `headers`
### 删除话题
|请求参数|类型|备注|
|---|---|---|
|title|必选|话题标题|
|:movieId|必选|电影id|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错,删除失败|
|“数据库中查询不到该电影”|数据库信息不足|
|“删除成功”|删除成功|
|“删除失败”|服务器出错，删除失败|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|

## `POST` `/discussion/:movieId/dis_comment`
### 发布话题的评论
###`application/x-www-form-urlencoded` `headers`
|请求参数|类型|备注|
|---|---|---|
|title|必选|评论的话题的标题|
|comment|必选|评论内容|
|:movieId|必选|电影id|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错,评论失败|
|“数据库中查询不到该电影”|数据库信息不足|
|“评论成功”|评论成功|
|“评论失败”|服务器出错，评论失败|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|

## `DELETE` `/discussion/:movieId/dis_comment`
###`application/x-www-form-urlencoded` `headers`
#### 删除话题评论
|请求参数|类型|备注|
|---|---|---|
|title|必选|评论的话题的标题|
|:movieId|必选|电影id|
|Authorization|必选|token|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|“服务器错误”|服务器出错,删除失败|
|“数据库中查询不到该电影”|数据库信息不足|
|“上传成功”|删除成功|
|“上传失败”|服务器出错，删除失败|
|”请求头中auth为空“|用户token为空|
|”请求头中auth格式有误“|用户token验证失败|

#MovieGet
## `GET` `/movieGet/search`
###`application/x-www-form-urlencoded`
#### 搜索
|请求参数|类型|说明|
|---|---|---|
|keyword|必选|关键字|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|"msg":"没有找到相关电影，换个搜索词试试吧。"|数据库中查询不到相关信息|
|movies|查询到的电影信息的切片|

## `GET` `/movieGet/:movieId`
### 获取单个电影的详情
|请求参数|类型|说明|
|---|---|---|
|:movieId|必选|电影id|

|返回参数|说明|
|---|---|
|date|返回消息|

|date|说明|
|---|---|
|"无法获取电影相关信息"|服务器错误或数据库中没有该电影信息|
|movie|电影详情的结构体|
























