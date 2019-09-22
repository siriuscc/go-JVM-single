[TOC]

## 搜索class文件


### 封装路径实体

-classpath/-cp支持语法：


```bash {.line-numbers}
java -cp path\to\classes ...
java -cp path\to\lib1.jar ...
java -cp path\to\lib2.zip ...

java -cp path \to\classes lib\a.jar;lib\b.jar;lib\c.zip ...

java -cp classes;lib\* ...

```

分别定义：


+ CompositeEntry: 按照；分割的多路径
+ DirEntry  ： 普通目录
+ ZipEntry  ： jar文件，zip文件
+ WildcardEntry ： path/* 带通配的


这是组合模式。


### 获得Jre目录

1. Xjre，指定jre的路径，
2. 如果不存在Xjre选项，则在当前目录下查找jre目录
3. 如果当前目录查找不到，可以在JAVA_HOME 中获得 


bootClasspath 指向 jre/lib
extClasspath 指向 jre/lib/ext


### 获得 appClasspath
1. 按照-cp 给出的路径
2. 如果没有，按照当前目录




### 读入二进制bytes




go 语言的访问控制：


大写开头：public
小写：包内访问



go中，

```go
// 这个函数只能给指针用
func (self *A)funtion(){

}

// 这个可以给指针用，也可以给对象用
func (self A)function(){

}

```

go语言返回指针不会发生复制，返回结构体会发生复制

