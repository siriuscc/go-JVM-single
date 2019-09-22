/**
 * 类的属性访问
 */
public class ClassDemo {

    private int id;
    private static int times = 1;

    private ClassDemo() {
        times++;
    }

    public ClassDemo(int id) {
        this();
        this.id = id;
    }

    public static void main(String[] args) {

        ClassDemo zlass = new ClassDemo(9);

        System.out.println(zlass.id); // 9
        System.out.println(times); // 2
        // System.out.println("msg:" + zlass.msg);

    }
}