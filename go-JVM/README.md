[TOC]




## 前言


此处记录开发的线索，主要的困难和idea


## 目录

* [搜索class文件](./doc/搜索class文件.md)
* [解析class文件](./doc/解析class文件.md)
* [运行时数据区](./doc/运行时数据区.md)
* [解析器和二进制指令](./doc/解析器和二进制指令.md)
* [类和对象模型](./doc/类和对象模型.md)
* [方法调用和返回](./doc/方法调用和返回.md)
* [数组和字符串](./doc/数组和字符串.md)
* [本地方法调用](./doc/本地方法调用.md)
* [异常处理](./doc/异常处理.md)
* [System类的构造](./doc/System类的构造.md)


// 多个catch 和finally 的实现 没做




occurrence:n, 发生，出现
concrete: adj,具体的



本月要做完 native方法的开发，完善异常机制。
下个月开始写论文。

今天主要的任务就是整理代码。完善一些没完成的功能








## golang 的一些细节

+ go switch自带break
+ go 没有继承，只有组合，对应`嵌入类型`
+ golang 不允许`循环导包`，开发者要思考项目的`领域模型`
+ 切片是对数组的部分引用，是左闭右开的
+ init 函数所在的包被导入，init函数自动调用一次
+ golang 的receiver 指针为空，也能调用函数
+ 小心结构体之间赋值的`深拷贝`


```go
	for _, field := range klass.fields {
		if !field.accessFlags.IsStatic() {
			field.slotId = slotId
			slotId++

			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
```
如果klass.fields 的组件类型不是指针，boom



## tool


参考:https://github.com/zxh0/classpy



### tips:


慎用copy code，容易导致一些细节的错误。

感觉需要一个hack 关机，然后提醒git提交



### 规范


+ golang的Getter 和Setter 推荐使用 大写字母开头的属性名，如name属性对应Name()
+ golang 返回布尔值的方法，建议用HasX(),IsX()
+ golang 推荐使用驼峰命名，方法名称一般是动词+名词(),如LoadClass(...)


本次开发，很多的命名沿用Java的规范，是GetSome格式的。


### go资料

+ [the-way-to-go 中文版](https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)