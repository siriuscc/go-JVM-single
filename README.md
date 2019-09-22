
本项目主要参考 《自己动手写Java虚拟机》，还有JVMS-1.8，基于golang 实现了一个单线程的Java虚拟机，不可用于生产环境，只用于学习原理。分享出来，希望有助于需要的同学.

### 项目综述


本项目 使用GO语言实现JVM，主要研究方向为 
+ Java的异常机制。
+ Java 的内存结构
+ Java的对象和类体系
+ Java的字节码


本次不研究，忽略的一些关键技术：

+ 不研究GC
+ 不研究just-in-time
+ 不研究大部分的native方法，只实现少量方法
+ 不研究多线程安全
+ 不研究网络通信
+ 不研究swing


### 本项目适合哪些人看？


+ 对JVM有一定了解，但希望自己实现一把的人
+ 对虚拟机有一定了解的人


推荐学Java的初学者，可以先看几篇周志明的《深入理解Java虚拟机》，然后对着照Java虚拟机规范来看本项目。


### 项目目录

+ go-JVM: 项目目录
	+ bin: go直接编译后的可运行文件目录
		+ classes：存放java 文件编译后的claess
		+ java ：存放java
	+ doc: 文档
+ refs： 主要参考文献
+ target：提取出来的最终文件


### 运行


1. 确保安装了golang环境
2. 运行go-JVM/bin/make-go.bat 进行编译
3. 在 go-JVM/bin 下会产生 .exe文件，就是最终的文件
4. 运行go-JVM/bin/show.bat 进行测试


运行实例：
```
go-JVM -cp=classespath BubbleSortTest

```
具体的运行规则可以参考show.bat


### 补充


#### OpenJDK

OpenJDK(HotSpot JVM、Javac)源代码学习研究(包括代码注释、文档、用于代码分析的测试用例)

+ [OpenJDK-Research 下载](https://github.com/codefollower/OpenJDK-Research)
+ [JDK8 下载](https://download.java.net/openjdk/jdk8)


##### 下载hotSpot

1. 在 [openJDK官方站点](http://hg.openjdk.java.net/)找到对应JDK版本，如jdk8，点击进入，
2. 再选择hotSpot，点击左边的download，选择下载的打包方式，[下载](http://hg.openjdk.java.net/jdk8/jdk8/hotspot/file/87ee5ee27509)



##### hotSpot源码结构说明

```
├─agent                            Serviceability Agent的客户端实现
├─make                             用来build出HotSpot的各种配置文件
├─src                              HotSpot VM的源代码
│  ├─cpu                            CPU相关代码（汇编器、模板解释器、ad文件、部分runtime函数在这里实现）
│  ├─os                             操作系相关代码
│  ├─os_cpu                         操作系统+CPU的组合相关的代码
│  └─share                          平台无关的共通代码
│      ├─tools                        工具
│      │  ├─hsdis                      反汇编插件
│      │  ├─IdealGraphVisualizer       将server编译器的中间代码可视化的工具
│      │  ├─launcher                   启动程序“java”
│      │  ├─LogCompilation             将-XX:+LogCompilation输出的日志（hotspot.log）整理成更容易阅读的格式的工具
│      │  └─ProjectCreator             生成Visual Studio的project文件的工具
│      └─vm                           HotSpot VM的核心代码
│          ├─adlc                       平台描述文件（上面的cpu或os_cpu里的*.ad文件）的编译器
│          ├─asm                        汇编器接口
│          ├─c1                         client编译器（又称“C1”）
│          ├─ci                         动态编译器的公共服务/从动态编译器到VM的接口
│          ├─classfile                  类文件的处理（包括类加载和系统符号表等）
│          ├─code                       动态生成的代码的管理
│          ├─compiler                   从VM调用动态编译器的接口
│          ├─gc_implementation          GC的实现
│          │  ├─concurrentMarkSweep      Concurrent Mark Sweep GC的实现
│          │  ├─g1                       Garbage-First GC的实现（不使用老的分代式GC框架）
│          │  ├─parallelScavenge         ParallelScavenge GC的实现（server VM默认，不使用老的分代式GC框架）
│          │  ├─parNew                   ParNew GC的实现
│          │  └─shared                   GC的共通实现
│          ├─gc_interface               GC的接口
│          ├─interpreter                解释器，包括“模板解释器”（官方版在用）和“C++解释器”（官方版不在用）
│          ├─libadt                     一些抽象数据结构
│          ├─memory                     内存管理相关（老的分代式GC框架也在这里）
│          ├─oops                       HotSpot VM的对象系统的实现
│          ├─opto                       server编译器（又称“C2”或“Opto”）
│          ├─prims                      HotSpot VM的对外接口，包括部分标准库的native部分和JVMTI实现
│          ├─runtime                    运行时支持库（包括线程管理、编译器调度、锁、反射等）
│          ├─services                   主要是用来支持JMX之类的管理功能的接口
│          ├─shark                      基于LLVM的JIT编译器（官方版里没有使用）
│          └─utilities                  一些基本的工具类
└─test                             单元测试
--------------------- 
```


### 主要参考

1. 封亚飞.解密Java虚拟机：JVM设计原理与实现[M].北京：电子工业出版社，2017.
2. 周志明.深入理解Java虚拟机[M].北京：机械工业出版社，2013.
3. 张秀宏.自己动手写Java虚拟机[M].北京：机械工业出版社，2016.
    +  Github:[自己动手写Java虚拟机](https://github.com/zxh0/jvmgo-book.git)
4. Tim Lindholm, Frank Yellin, Gilad Bracha, Alex Buckley.The Java®  Virtual Machine Specification Java SE 8 Edition[S].

