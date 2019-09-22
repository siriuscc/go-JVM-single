
## System 类的构造
[TOC]

本文以实现打印HelloWorld 为主线，基于 自己正在写JVM的实践，一步步探讨System类。

### System 类和 输入输出流

输入输出流是在 线程初始化的时候绑定到System 类的静态属性上的。

```java
public final class System {

    public final static InputStream in = null;
    public final static PrintStream out = null;
    public final static PrintStream err = null;    
    /* 
     * 通过静态初始化注册native方法
     * 
     * 虚拟机 会调用 System.initializeSystemClass() 来完成这个类初始化
     */
    private static native void registerNatives();
    static {
        registerNatives();
    }

    /**
     * Initialize the system class.  Called after thread initialization.
     */
    private static void initializeSystemClass() {
        props = new Properties();
        initProperties(props);  // initialized by the VM
        ...

        FileInputStream fdIn = new FileInputStream(FileDescriptor.in);
        FileOutputStream fdOut = new FileOutputStream(FileDescriptor.out);
        FileOutputStream fdErr = new FileOutputStream(FileDescriptor.err);
        setIn0(new BufferedInputStream(fdIn));
        setOut0(newPrintStream(fdOut, props.getProperty("sun.stdout.encoding")));
        setErr0(newPrintStream(fdErr, props.getProperty("sun.stderr.encoding")));
        ...
    }

```

线程初始化，会加载System类，执行类初始化方法`<clinit>`，`<clinit>` 是 代码中 静态变量赋初始值的语句 和静态代码块的语句集合。


一般的，如果一个类有`native`方法，会在静态代码块调用 `registerNatives()`来注册native函数。 registerNatives 是由JVM实现的。特殊的，在System 的静态代码块中，官方注释中说 registerNatives  会调用 System.initializeSystemClass()。在 openJDK 的目录下可以找到(jdk-8\src\share\native\java\lang\System.c )。

```cpp

JNIEXPORT void JNICALL
Java_java_lang_System_registerNatives(JNIEnv *env, jclass cls)
{
    (*env)->RegisterNatives(env, cls,
                            methods, sizeof(methods)/sizeof(methods[0]));
}
```

后面的内容不再深入，先打住。

在同一个文件下，再看一下对out，err,的初始化。 JNI方法 的命名是 包路径_方法名，以下划线分隔。


```cpp

JNIEXPORT void JNICALL
Java_java_lang_System_setIn0(JNIEnv *env, jclass cla, jobject stream)
{
    jfieldID fid =
        // 得到属性id
        (*env)->GetStaticFieldID(env,cla,"in","Ljava/io/InputStream;");
    if (fid == 0)
        return;
    // 根据属性id，设置实例 的属性
    (*env)->SetStaticObjectField(env,cla,fid,stream);
}

JNIEXPORT void JNICALL
Java_java_lang_System_setOut0(JNIEnv *env, jclass cla, jobject stream)
{
    jfieldID fid =
        (*env)->GetStaticFieldID(env,cla,"out","Ljava/io/PrintStream;");
    if (fid == 0)
        return;
    (*env)->SetStaticObjectField(env,cla,fid,stream);
}

JNIEXPORT void JNICALL
Java_java_lang_System_setErr0(JNIEnv *env, jclass cla, jobject stream)
{
    jfieldID fid =
        (*env)->GetStaticFieldID(env,cla,"err","Ljava/io/PrintStream;");
    if (fid == 0)
        return;
    (*env)->SetStaticObjectField(env,cla,fid,stream);
}

```


### FileInputStream 的相关实现


System类的启动还需要 FileInputStream 的一些native方法支持。
找到FileInputStream.initIDs。 JDoc没有说时干嘛的，被迫去看openJDK。(jdk-8\src\share\native\java\io\FileInputStream.c)


```cpp
jfieldID fis_fd; /* id for jobject 'fd' in java.io.FileInputStream */
/**************************************************************
 * static methods to store field ID's in initializers
 */
Java_java_io_FileInputStream_initIDs(JNIEnv *env, jclass fdClass) {
    fis_fd = (*env)->GetFieldID(env, fdClass, "fd", "Ljava/io/FileDescriptor;");
}
```

fis_id 应该是 fd对象在 FileInputStream 中的id，类似于实例槽。fd应该是类似于文件句柄。




### JVM 初始化参数

在System类初始化的时，会调用`initializeSystemClass`方法，


也就是在执行：
```java
    private static void initializeSystemClass() {
        ...
        FileInputStream fdIn = new FileInputStream(FileDescriptor.in);
        FileOutputStream fdOut = new FileOutputStream(FileDescriptor.out);
        FileOutputStream fdErr = new FileOutputStream(FileDescriptor.err);
        setIn0(new BufferedInputStream(fdIn));
        // 如果不设置sun.stdout.encoding，file.encoding, 这里会出问题
        setOut0(newPrintStream(fdOut, props.getProperty("sun.stdout.encoding")));
        setErr0(newPrintStream(fdErr, props.getProperty("sun.stderr.encoding")));
        ...
    }
```



间接的，会到这个：
D:\surroundings\jdk1.8.0_131\src.zip!\java\nio\charset\Charset.java


```java
    public static Charset defaultCharset() {
        if (defaultCharset == null) {
            synchronized (Charset.class) {
                // csn 为null
                String csn = AccessController.doPrivileged(
                    new GetPropertyAction("file.encoding"));
                    // lookup方法抛出异常
                Charset cs = lookup(csn);
                if (cs != null)
                    defaultCharset = cs;
                else
                    defaultCharset = forName("UTF-8");
            }
        }
        return defaultCharset;
    }

    private static Charset lookup(String charsetName) {
        if (charsetName == null)
            // 这里抛出
            throw new IllegalArgumentException("Null charset name");
        Object[] a;
        if ((a = cache1) != null && charsetName.equals(a[0]))
            return (Charset)a[1];
        // We expect most programs to use one Charset repeatedly.
        // We convey a hint to this effect to the VM by putting the
        // level 1 cache miss code in a separate method.
        return lookup2(charsetName);
    }
```


```java

// 获取参数的一个行动
public class GetPropertyAction implements PrivilegedAction<String> {
    private String theProp;
    private String defaultVal;// null

    public GetPropertyAction(String theProp) {
        this.theProp = theProp;
    }

    public GetPropertyAction(String var1, String var2) {
        this.theProp = var1;
        this.defaultVal = var2;
    }

    public String run() {

        String var1 = System.getProperty(this.theProp);
        return var1 == null ? this.defaultVal : var1;
    }
}
```

所以，__在虚拟机内部，需要初始化一些必要的属性。__


```go
var sysProps = map[string]string{
	"java.version":         "1.8.0",
	"java.vendor":          "jvm.go",
	"java.vendor.url":      "https://github.com/zxh0/jvm.go",
	"java.home":            "todo",
	"java.class.version":   "52.0",
	"java.class.path":      "todo",
	"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
	"os.name":              runtime.GOOS,   
	"os.arch":              runtime.GOARCH, 
	"os.version":           "",             
	"file.separator":       "/",            
	"path.separator":       ":",            
	"line.separator":       "\n",           
	"user.name":            "",             
	"user.home":            "",             
	"user.dir":             ".",            
	"user.country":         "CN",           
	"file.encoding":        "UTF-8",
	"sun.stdout.encoding":  "UTF-8",
    "sun.stderr.encoding":  "UTF-8",
}
```

JVM启动时，会调用System类的initProperties 方法，在这个方法写入默认属性。具体的思路就是调用props.setProperty(key,value)


```java
    private static native Properties initProperties(Properties props);
```

```java
public class FileInputStream extends InputStream{

    /* File Descriptor - handle to the open file */
    private final FileDescriptor fd;
```


### Unsafe 类

Unsafe类提供了操作指针`直接访问内存`，直接`操作`内存数据的能力。为了支持 System类的运行，Unsafe 要实现的方法如下。

```java

    // 获取 数组的第一个元素 相对于 这个类的偏移量
    public native int arrayBaseOffset(Class<?> klass);

    // 获取数组 的步长， offset 是arr[0], offset+scale 就能得到arr[1]
    public native int arrayIndexScale(Class<?> klass);

    // 指针的size, sizeof(p),应该是8
    public native int addressSize();

    /** 获取field 的 slot属性
     *  - 对于实例，field是以slot占位置的
     *  - 可以理解为slot是一个int32，long和double占两个slot
     */
    public native long objectFieldOffset(Field field);
    /** CAS 操作
     *  - 对于数组，就是对arr[i] 做CAS操作
     *  - 对于非数组，就是 对 offset 上的属性，做CAS操作
     */
    public final native boolean compareAndSwapObject(Object oop, long offset, Object expected, Object updateValue);
    // CAS ,操作的是int
    public final native boolean compareAndSwapInt(Object obj, long offset, int expected, int updateValue);
    // CAS， 操作的是long
    public final native boolean compareAndSwapLong(Object obj, long offset, long expected, long updateValue);

    public native int getIntVolatile(Object obj, long offset);
    public native Object getObject(Object obj, long offset);


    /*****************   内存操作  **************/

    // 分配 size个byte，返回地址
    public native long allocateMemory(long size);
    public native void freeMemory(long address);
    // 对 [address,address+size] 重新分配
    public native long reallocateMemory(long address, long size);
    
    public native byte getByte(Object obj, long var2);
    public native void putLong(long address, long value);


```

#### CAS原理

看一下 `Unsafe_CompareAndSwapInt`。(hotspot-8\src\share\vm\prims\unsafe.cpp)

```cpp
UNSAFE_ENTRY(jboolean, Unsafe_CompareAndSwapInt(JNIEnv *env, jobject unsafe, jobject obj, jlong offset, jint e, jint x))
  UnsafeWrapper("Unsafe_CompareAndSwapInt");
  oop p = JNIHandles::resolve(obj);
  jint* addr = (jint *) index_oop_from_field_offset_long(p, offset);
  return (jint)(Atomic::cmpxchg(x, addr, e)) == e;
UNSAFE_END
```

这里核心的调用时是`Atomic::cmpxchg(x, addr, e)`，从名字上看，这是一个原子操作。(hotspot-8\src\share\vm\runtime\atomic.cpp)

```cpp
unsigned Atomic::cmpxchg(unsigned int exchange_value,
                         volatile unsigned int* dest, unsigned int compare_value) {
  assert(sizeof(unsigned int) == sizeof(jint), "more work to do");
  return (unsigned int)Atomic::cmpxchg((jint)exchange_value, (volatile jint*)dest,
                                       (jint)compare_value);
}
```

最终实现取决于底层OS，比如linux x86，实现内联在hotspot部分代码：(hotspot-8\src\os_cpu\linux_x86\vm\atomic_linux_x86.inline.hpp)
```cpp
// Adding a lock prefix to an instruction on MP machine
#define LOCK_IF_MP(mp) "cmp $0, " #mp "; je 1f; lock; 1: "


inline jint     Atomic::cmpxchg    (jint     exchange_value, volatile jint*     dest, jint     compare_value) {
  int mp = os::is_MP();
  __asm__ volatile (LOCK_IF_MP(%4) "cmpxchgl %1,(%3)"
                    : "=a" (exchange_value)
                    : "r" (exchange_value), "a" (compare_value), "r" (dest), "r" (mp)
                    : "cc", "memory");
  return exchange_value;
}
```


从上面的代码中可以看到，如果是CPU是`多核`(multi processors)的话，会添加一个lock;前缀，这个lock;前缀也是`内存屏障`，它的作用是在执行后面指令的过程中`锁总线`(或者是锁cacheline)，保证一致性。后面的指令`cmpxchgl`就是x86的比较并交换指令了(汇编指令)。

这里能大致的理解为：从汇编层面 利用 `cmpxchgl` 做原子写入，而加锁，是给指令加的，汇编语言不了解，姑且跳过。打住了。


参考：[Jdk1.6 JUC源码解析(1)-atomic-AtomicXXX](https://brokendreams.iteye.com/blog/2250109)







### java 的 AccessController.doPrivileged使用


AccessController.doPrivileged 意思是对指定文件不用做权限检查. 




[java 的 AccessController.doPrivileged使用](https://huangyunbin.iteye.com/blog/1942509)






### Object.hashCode 的实现细节


实现Object.hashCode

在 jdk-8\src\share\native\java\lang\Object.c，可以看到前面有一段：

```cpp
static JNINativeMethod methods[] = {
    {"hashCode",    "()I",                    (void *)&JVM_IHashCode},
    {"wait",        "(J)V",                   (void *)&JVM_MonitorWait},
    {"notify",      "()V",                    (void *)&JVM_MonitorNotify},
    {"notifyAll",   "()V",                    (void *)&JVM_MonitorNotifyAll},
    {"clone",       "()Ljava/lang/Object;",   (void *)&JVM_Clone},
};

JNIEXPORT void JNICALL
Java_java_lang_Object_registerNatives(JNIEnv *env, jclass cls)
{
   
    (*env)->RegisterNatives(env, cls,
                            methods, sizeof(methods)/sizeof(methods[0]));
}
```

这里将 `hashCode()` 映射到 `(void *)&JVM_IHashCode` 这个函数。

关于`(void *)&JVM_IHashCode` 的定义，在openJDK中。(jdk-8\src\share\javavm\export\jvm.h)


```cpp
/*
 * java.lang.Object
 */
JNIEXPORT jint JNICALL
JVM_IHashCode(JNIEnv *env, jobject obj);
```

作为JVM 的implement，hotspot8实现了这个方法。 (hotspot-8\src\share\vm\prims\jvm.cpp)

```cpp
// java.lang.Object ///////////////////////////////////////////////

JVM_ENTRY(jint, JVM_IHashCode(JNIEnv* env, jobject handle))
  JVMWrapper("JVM_IHashCode");
  // as implemented in the classic virtual machine; return 0 if object is NULL
  return handle == NULL ? 0 : ObjectSynchronizer::FastHashCode (THREAD, JNIHandles::resolve_non_null(handle)) ;
JVM_END
```

如果 对象的handle 为null，返回0，否则调用 FastHashCode。继续跟进去，看到最后生成的逻辑。 细节不解释不深入，官方给的代码注释写得很清楚了。需要注意的是，如果对象为null，java中调用`null`对象的函数会直接抛出`NullPointerException`。


直接看到hash值的生成过程,在 get_next_hash
方法。(hotspot-8\src\share\vm\runtime\synchronizer.cpp)：

```cpp
// hashCode() generation :
//
// Possibilities:
// * MD5Digest of {obj,stwRandom}
// * CRC32 of {obj,stwRandom} or any linear-feedback shift register function.
// * A DES- or AES-style SBox[] mechanism
// * One of the Phi-based schemes, such as:
//   2654435761 = 2^32 * Phi (golden ratio)
//   HashCodeValue = ((uintptr_t(obj) >> 3) * 2654435761) ^ GVars.stwRandom ;
// * A variation of Marsaglia's shift-xor RNG scheme.
// * (obj ^ stwRandom) is appealing, but can result
//   in undesirable regularity in the hashCode values of adjacent objects
//   (objects allocated back-to-back, in particular).  This could potentially
//   result in hashtable collisions and reduced hashtable efficiency.
//   There are simple ways to "diffuse" the middle address bits over the
//   generated hashCode values:
static inline intptr_t get_next_hash(Thread * Self, oop obj) {
  intptr_t value = 0 ;
  if (hashCode == 0) {
     // This form uses an unguarded global Park-Miller RNG,
     // so it's possible for two threads to race and generate the same RNG.
     // On MP system we'll have lots of RW access to a global, so the
     // mechanism induces lots of coherency traffic.
     value = os::random() ;
  } else
  if (hashCode == 1) {
     // This variation has the property of being stable (idempotent)
     // between STW operations.  This can be useful in some of the 1-0
     // synchronization schemes.
     intptr_t addrBits = cast_from_oop<intptr_t>(obj) >> 3 ;
     value = addrBits ^ (addrBits >> 5) ^ GVars.stwRandom ;
  } else
  if (hashCode == 2) {
     value = 1 ;            // for sensitivity testing
  } else
  if (hashCode == 3) {
     value = ++GVars.hcSequence ;
  } else
  if (hashCode == 4) {
      // 直接用内存地址
     value = cast_from_oop<intptr_t>(obj) ;
  } else {
     // Marsaglia's xor-shift scheme with thread-specific state
     // This is probably the best overall implementation -- we'll
     // likely make this the default in future releases.
     unsigned t = Self->_hashStateX ;
     t ^= (t << 11) ;
     Self->_hashStateX = Self->_hashStateY ;
     Self->_hashStateY = Self->_hashStateZ ;
     Self->_hashStateZ = Self->_hashStateW ;
     unsigned v = Self->_hashStateW ;
     v = (v ^ (v >> 19)) ^ (t ^ (t >> 8)) ;
     Self->_hashStateW = v ;
     value = v ;
  }

  value &= markOopDesc::hash_mask;
  if (value == 0) value = 0xBAD ;
  assert (value != markOopDesc::no_hash, "invariant") ;
  TEVENT (hashCode: GENERATE) ;
  return value;
}
```


+ hashcode=1，使用最简单的随机算法。
+ hashcode=4，直接使用的内存地址。
+ 默认使用的是hashcode>=5的 `xor-shift scheme`算法。 可以用JVM parameter -XX:hashCode来调整。

`xor-shift scheme`是弗罗里达州立大学一位叫做George Marsaglia的老师发明的使用位移和异或运算生成随机数的方法, 所以在计算机上运算速度非常快(移位指令需要的机器周期更少).






代码中给出了四种策略，那么，我们的实现方案又如何呢？当然是随机hash。






为OopDesc添加一个hashCode 的属性。对于一个oop，一旦构造出来，如果没有hashcode，就给一个随机数。如果有，就不构造。注意，这里hash值非0；

另外，在Java中，hashCode 相等，可以不是同一个对象，具体在HashMap的注意事项中就有。

Java中hashCode 应该是基于实时计算的，这里做缓存。我们一次性计算。



