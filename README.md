## Hanaya - 图片压缩服务器

目前仅提供基于libwebp的图片质量压缩，支持根据url传参缩放、压缩图片质量。

### 依赖安装配置
本项目依赖[libwebp](https://developers.google.com/speed/webp)，下文所述安装包可前往[官网仓库](https://storage.googleapis.com/downloads.webmproject.org/releases/webp/index.html)下载，MacOS与Linux可以通过包管理工具直接安装

##### MacOS:
    brew install webp

##### Linux:
    sudo apt-get update
    sudo apt-get install libwebp-dev

##### Windows
解压压缩包到任意路径，如D://libwebp-1.2.1-windows-x64-no-wic

配置如下环境变量：

    CGO_CFLAGS     -ID://libwebp-1.2.1-windows-x64-no-wic/include
    CGO_LDFLAGS    -LD://libwebp-1.2.1-windows-x64-no-wic/lib

- 注1：Windows需下载后缀为no-wic的版本，如libwebp-1.2.1-windows-x64-no-wic.zip  ，默认的包中为lib文件，在windows中cgo无法正常编译。
- 注2：M1版MacOS，若是使用homebrew安装，注意包路径有可能不是默认的 **/usr/local/include/** ，需要配置如下CGO环境变量

M1 MacOS环境变量：

    CGO_CFLAGS    -I/opt/homebrew/include
    CGO_LDFLAGS  -L/opt/homebrew/lib

### 压缩图片

项目启动后，默认端口为8889

#### 接口地址：POST /webp + 参数组合 + file

|  参数 |  用途 | 范围		|
| ------------- | -------------|--------------|
| l  |  level，无损压缩时的压缩等级，越大压缩比越高，压缩所需的时间也越长，有损压缩时此参数无效 | 0-9|
| q  |  quality，有损压缩时的图片质量，无损压缩时此参数无效 | 0-100|
| s | size,长宽范围，使用\*隔开，会以超过此范围的最长边为参照，等比压缩图片大小 | -|
| lossless | 标识压缩为无损压缩，默认为有损压缩 | -|

表单参数：

| 参数名  | 描述 |
| ------------ |------|
|  file | 需要压缩的图片

eg:
想要将图片限制在1200\*1200的范围内，并进行一个百分之75的有损压缩，你应该请求的地址为：/webp/s1200\*1200/q75

#### 返回值：
压缩成功：
> Content-Type: image/webp

>Body: 压缩后的图片内容

##### 压缩失败：
> Content-Type：application/json

Body:
```json
{
	"isOk":false,
	"msg":"reason",
}
```