
public class InvokeDemo implements Runnable {

    public InvokeDemo() {
        System.out.println("invoke instance Special");
    }

    public static void staticMethod() {
        System.out.println("invole static Method");
    }

    public void instanceMethod() {
        System.out.println("invoke instance Method");
    }

    public void test() {
        // 调用静态方法 invokestatic
        InvokeDemo.staticMethod();
        InvokeDemo demo = new InvokeDemo();
        // 调用实例方法 invokevirtual
        demo.instanceMethod();
        // 调用 父类的方法,invokespecial
        super.equals(null);
        // 调用接口方法，对应invokeinterface指令
        this.run();
    }

    public static void main(String[] args) {
        // 调用 invokespecial 指令
        InvokeDemo invoke = new InvokeDemo();
        // 调用实例方法
        invoke.test();
    }

    public void run() {
        System.out.println("invoke instance Interface");
    }
}