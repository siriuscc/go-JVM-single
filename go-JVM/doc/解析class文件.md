[TOC]

## 解析class文件

反编译输出附加信息
```
javap -verbose HelloWorld.class
```


### go 和java的基本类型对照表


|GO|Java| 说明
|:-:|:-:|:-:|
|int8|byte|
|uint||
|int16|short|
|uint16|char|
|int32|int|
|uint32||
|int64|long
|uint64||
|float32|float|
|float64|double|


## ClassFile 格式


```cpp
ClassFile {
    u4 magic;       // 标志这个文件是class文件，0xCAFEBABE.
    u2 minor_version;
    u2 major_version;
    u2 constant_pool_count;
    cp_info constant_pool[constant_pool_count-1];
    u2 access_flags;
    u2 this_class;
    u2 super_class;
    u2 interfaces_count;
    u2 interfaces[interfaces_count];
    u2 fields_count;
    field_info fields[fields_count];
    u2 methods_count;
    method_info methods[methods_count];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
```

### major_version.minor_version

版本号，必须class文件的版本号必须小于等于当前的jvm的版本号。


### constant_pool_count，constant_pool

常量池的长度和数组

常量池 是一个table，表示字符常量，类和接口名，属性名，和其他在ClassFIle结构中和子结构中的常量。每一个实体，首先是一个tag标记位。常量池的标记位为 1到 constant_pool_count-1


### access_flags：

表示访问标识

|flag|byte-code|desc|
|:-:|:-:|:-:|
|ACC_PUBLIC |0x0001 |Declared public; may be accessed from outside its package.
|ACC_FINAL |0x0010 |Declared final; no subclasses allowed.
|ACC_SUPER |0x0020 |Treat superclass methods specially when invoked by the invokespecial instruction.
|ACC_INTERFACE |0x0200 |Is an interface, not a class.
|ACC_ABSTRACT |0x0400 |Declared abstract; must not be instantiated.
|ACC_SYNTHETIC |0x1000 |Declared synthetic; not present in the source code.
|ACC_ANNOTATION |0x2000 |Declared as an annotation type.
|ACC_ENUM |0x4000 |Declared as an enum type.




+ 如果这是一个interface, ACC_INTERFACE位被设置，如果没有，这是一个类。

+ 如果ACC_INTERFACE，则ACC_ABSTRACT 也必须被设置。ACC_FINAL,ACC_SUPER, and ACC_ENUM 不能被设置。

+ 如果ACC_INTERFACE 没被设置，表中其他标志都可以设置，当然，不能同时存在ACC_FINAL 和 ACC_ABSTRACT。
+ ACC_SUPER标志指示两种（class or interface）替代语义中的哪一种，如果它出现在该类或接口中，则将由调用特殊指令(invokSpecial)来表示。Java虚拟机指令集的编译器应该设置ACC_SUPER标志。在Java SE 8及更高版本中，Java虚拟机考虑在每个类文件中设置ACC_SUPER标记，而不考虑类文件中标记的实际值和类文件的版本。
+ ACC_SYNTHETIC: 表示类或接口是由编译程序生成的，不会出现在源码
+ ACC_ANNOTATION：注解标记。如果有ACC_ANNOTATION，ACC_INTERFACE也必须设置
+ ACC_ENUM： 表示这个类或者他的父类是一个枚举类








### this_class

一个 constant_pool 的CONSTANT_Class_info索引，表示这个类（或接口）的信息。


### super_class

+ 0或者非零整数。0表示没有父类，表示Object类，或者是interface。
+ 非零指向常量池索引，对应一个CONSTANT_Class_info，表示一个父类信息。


### 



CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}


CONSTANT_Class_info


### cp_info :常量池 

```cpp
cp_info {
    u1 tag;
    u1 info[];
}
```
tag ：

|Constant Type |Value|
|:-|:-:|
|CONSTANT_Class| 7|
|CONSTANT_Fieldref| 9|
|CONSTANT_Methodref |10|
|CONSTANT_InterfaceMethodref |11|
|CONSTANT_String| 8|
|CONSTANT_Integer| 3|
|CONSTANT_Float| 4|
|CONSTANT_Long| 5|
|CONSTANT_Double| 6|
|CONSTANT_NameAndType |12|
|CONSTANT_Utf8| 1|
|CONSTANT_MethodHandle |15|
|CONSTANT_MethodType |16|
|CONSTANT_InvokeDynamic| 18


info 是一个常量池索引



CONSTANT_Utf8：


tips: 这里是按照modified UTF-8 编码的，不是UTF8。 



> String content is encoded in modified UTF-8. Modified UTF-8 strings are encoded so that code point sequences that contain only non-null ASCII characters can be represented using only 1 byte per code point, but all code points in the Unicode codespace can be represented. Modified UTF-8 strings are not null-terminated. 

> There are two differences between this format and the "standard" UTF-8 format. First, the null character (char)0 is encoded using the 2-byte format rather than the 1-byte format, so that modified UTF-8 strings never have embedded nulls. Second, only the 1-byte, 2-byte, and 3-byte formats of standard UTF-8 are used. The Java Virtual Machine does not recognize the four-byte format of standard UTF-8; it uses its own two-times-three-byte format instead.


标准UTF8和MUTF8主要有两个区别：

1. MUTF8 使用两个byte来编码null字符串
2. 只使用了标准UTF8中的1-byte,2-byte,3-byte. JVM 不 解析4-byte的字符串











### field_info

```cpp
field_info {
    u2 access_flags;
    u2 name_index;
    u2 descriptor_index;
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
```

### access_flags ：

参考jvms

### attribute_info


attribute_info 为了实现可扩展，并没有强制用tag，所以是依据attribute_name_index 来区别的，这导致很多虚拟机可以自己实现一些attribute。但是，至少都存在attribute_name_index，attribute_length。attribute_length表示的是后面的字节数目，不包括前面attribute_name_index和attribute_length 消耗的六个字节。


```go
attribute_info {
    u2 attribute_name_index; // 表示属性名
    u4 attribute_length;
    u1 info[attribute_length];
}
```

五个属性对于Java虚拟机正确解释类文件至关重要：


+ ConstantValue
+ Code
+ StackMapTable
+ Exceptions
+ BootstrapMethods


12个属性对于Java SE平台的类库正确解释类文件至关重要：
+ InnerClasses
+ EnclosingMethod
+ Synthetic
+ Signature
+ RuntimeVisibleAnnotations
+ RuntimeInvisibleAnnotations
+ RuntimeVisibleParameterAnnotations
+ RuntimeInvisibleParameterAnnotations
+ RuntimeVisibleTypeAnnotations
+ RuntimeInvisibleTypeAnnotations
+ AnnotationDefault
+ MethodParameters


六个属性对于Java虚拟机或Java SE平台的类库对类文件的正确解释并不重要，但对于工具很有用：
+ SourceFile
+ SourceDebugExtension
+ LineNumberTable
+ LocalVariableTable
+ LocalVariableTypeTable
+ Deprecated



本次实现前面的5+12个属性。后面的六个属性不实现。

#### ConstantValue_attribute

```go
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index; // 指向常量池
}
```
constantvalue_index必须指向如下表格的格式。
|Field Type |Entry Type
|:-|:-|
|long   |CONSTANT_Long
|float  |CONSTANT_Float
|double |CONSTANT_Double
|int, short, char, byte, boolean |CONSTANT_Integer
|String |CONSTANT_String



#### Code_attribute

```cpp

Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;   // 操作数栈的最大size
    u2 max_locals;  // 最大本地变量表,
    u4 code_length; // 指令的长度
    u1 code[code_length];   // 指令
    u2 exception_table_length;  // 异常表
    {   u2 start_pc;    // 异常的handle范围，[start_pc,end_pc)
        u2 end_pc;
        u2 handler_pc;  // 指向异常处理指令的开始位置
        u2 catch_type;  // catch类型，指向一个CONSTANT_Class_info
    } exception_table[exception_table_length];
    u2 attributes_count;    
    attribute_info attributes[attributes_count]; //  递归结构，存储一些 code的属性信息
}

```

对于本地变量表，记得long or double 是占用两个位置的。
如果catch_type为0，catch 所有异常，（常量池索引从 1 开始）

code属性包含了一个方法的JVM指令和辅助信息。包括\<init>和\<cinit>方法。


如方法是native or abstract,就不应该有code属性。否则，code属性必须存在。


#### StackMapTable

变长属性，是code属性的属性表里的属性。用于运行时类型检查


在50.0及以上版本，如果一个Code属性不存在StackMapTable，会有一个 implicit stack map（隐含的stack映射）。
number_of_entries=0 时，两者等价。

```cpp
StackMapTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_entries;
    stack_map_frame entries[number_of_entries];
}
```


> A stack map frame specifies (either explicitly or implicitly) the bytecode offset at which it applies, and the verification types of local variables and operand stack entries for that offset.

一个 stack map frame(帧) 指定了 对应的 字节码的位置，本地变量表和操作数栈的缩进。


> Each stack map frame described in the entries table relies on the previous frame for some of its semantics. The first stack map frame of a method is implicit, and computed from the method descriptor by the type checker (§4.10.1.6). The stack_map_frame structure at entries[0] therefore describes the second stack map frame of the method.


每一个frame 描述依赖于前一个frame的语义。方法的第一个frame 是隐含 由 type checker计算的。所以 entries[0] 描述的是第二个frame。




这里由于涉及到方法的校验，本次不实现，所以不搞。。。




#### Exceptions_attribute

记录方法抛出的异常
```c
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions]; // 每一个都指向一个CONSTANT_Class_info，表示具体抛出的异常类
}
```

仅当满足下面三个条件至少一个，方法才抛出异常：

1. exception 是 RuntimeException 或其子类 的一个实例
2. exception是 Error 或其子类的一个实例
3. exception 是 exception_index_table 里描述的类或其子类的实例。


#### LineNumberTable_attribute

辅助信息，主要用于调试用。记录源代码的行数和对应指令的映射。


```c
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;    // 开始的指令位置
        u2 line_number; // 源代码中对应的行数
    } line_number_table[line_number_table_length];
}
```


#### Synthetic_attribute

固定长度。用于 ClassFile, field_info, or method_info . A class
member that does not appear in the source code must be marked using a Synthetic attribute, or else it must have its ACC_SYNTHETIC flag set.


```
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
```



#### SourceFile

可选的，固定长度的属性。记录class文件对应的文件名。不会包含目录信息。


```c
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index; // 指向一个 CONSTANT_Utf8_info
}
```



```c
LocalVariableTable_attribute {
u2 attribute_name_index;
u4 attribute_length;
u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;  // 变量名，指向一个 CONSTANT_Utf8_info
        u2 descriptor_index;    // 描述符，CONSTANT_Utf8_info
        u2 index;   // 此变量必须在本地变量表的索引 为index
    } local_variable_table[local_variable_table_length];
}
```


在[start_pc, start_pc + length)，对应的变量必须存在值。
descriptor_index，是一个field descriptor：

+ FieldDescriptor:FieldType
+ FieldType:
    + BaseType
    + ObjectType
    + ArrayType
+ BaseType:(one of)
    + B C D F I J S Z
+ ObjectType:
    + L ClassName ;
+ ArrayType:[ ComponentType
+ ComponentType: FieldType

#### Deprecated_attribute


```c
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
```


可选，固定长度属性，用于 ClassFile, field_info, or method_info  结构。 表示

使用Deprecated_attribute 标记，

A class, interface, method, or field may be marked using a Deprecated attribute to indicate that the class, interface, method, or field has been superseded



#### Signatures_attribute

Signatures encode declarations written in the Java programming language that use types outside the type system of the Java Virtual Machine. They support reflection and debugging, as well as compilation when only class files are available.




#### RuntimeVisibleAnnotations

RuntimeVisibleAnnotations
```
RuntimeVisibleAnnotations_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 num_annotations;
    annotation annotations[num_annotations];
}


annotation {
    u2 type_index;
    u2 num_element_value_pairs;
    {   u2 element_name_index;
        element_value value;
    } element_value_pairs[num_element_value_pairs];
}
```




---

|class  |access_flags||
|:-|:-|:-|
|ACC_PUBLIC      |0x0001 |Declared public; may be accessed from outside its package.
|ACC_FINAL       |0x0010 |Declared final; no subclasses allowed.
|ACC_SUPER       |0x0020 |Treat superclass methods specially when invoked by the invokespecial instruction.
|ACC_INTERFACE   |0x0200 |Is an interface, not a class.
|ACC_ABSTRACT    |0x0400 |Declared abstract; must not be instantiated.
|ACC_SYNTHETIC   |0x1000 |Declared synthetic; not present in the source code.
|ACC_ANNOTATION  |0x2000 |Declared as an annotation type.
|ACC_ENUM        |0x4000 |Declared as an enum type.





### method_info

```cpp
method_info {
    u2 access_flags;
    u2 name_index;
    u2 descriptor_index;
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
```

|Flag Name |Value |Interpretation
|:-|:-:|:-|
|ACC_PUBLIC| 0x0001| Declared public; may be accessed from outside its package.
|ACC_PRIVATE| 0x0002| Declared private; accessible only within the defining class.
|ACC_PROTECTED |0x0004 |Declared protected; may be accessed within subclasses.
|ACC_STATIC |0x0008 |Declared static.
|ACC_FINAL  |0x0010 |Declared final; must not be overridden 
|ACC_SYNCHRONIZED |0x0020 |Declared synchronized; invocation is wrapped by a monitor use.
|ACC_BRIDGE  |0x0040 |A bridge method, generated by the compiler.
|ACC_VARARGS |0x0080 |参数数量可变
|ACC_NATIVE  |0x0100 |Declared native; implemented in a language other than Java.
|ACC_ABSTRACT|0x0400 |Declared abstract; no implementation is provided.
|ACC_STRICT  |0x0800 |Declared strictfp; floating-point mode is FPstrict.
|ACC_SYNTHETIC|0x1000|Declared synthetic; not present in the source code.




规则：

+ ACC_PUBLIC, ACC_PRIVATE,and ACC_PROTECTED 只能三选一
+ 接口的Methods可以有表中大部分属性，除了ACC_PROTECTED, ACC_FINAL, ACC_SYNCHRONIZED, and ACC_NATIVE 
+ 版本号小于52.0. 接口的Method 必须有ACC_PUBLIC&ACC_ABSTRACT；版本号>=52.0，接口的每个方法必须设置其ACC_PUBLIC和ACC_PRIVE标志中的一个。





+ 如果一个方法（对于类或接口），有ACC_ABSTRACT，必须不能有 ACC_PRIVATE, ACC_STATIC, ACC_FINAL, ACC_SYNCHRONIZED, ACC_NATIVE, or ACC_STRICT 


+ 所有构造函数 可以有一个ACC_PUBLIC, 或ACC_PRIVATE, 或ACC_PROTECTED。也可以设置 ACC_VARARGS, ACC_STRICT, and ACC_SYNTHETIC ，但不能有表中的其他flags




所有表中没有提到的位，应该为0，并且在后续的版本中可能有实现


All bits of the access_flags item not assigned in Table 4.6-A are reserved for
future use. They should be set to zero in generated class files and should be
ignored by Java Virtual Machine implementations.


descriptorIndex：
指向一个CONSTANT_Utf8_info ，描述符
```
例如：
    Object m(int i, double d, Thread t) {...}
    方法描述符 (IDLjava/lang/Thread;)Ljava/lang/Object;
```




- [ ] 为什么使用MUTF8 而不是标准UTF8