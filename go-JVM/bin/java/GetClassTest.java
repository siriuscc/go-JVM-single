public class GetClassTest {

    public static void main(String[] args) {
        System.out.println(void.class.getName());// void
        System.out.println(short.class.getName());// short
        System.out.println(int.class.getName());// int
        System.out.println(long.class.getName());// long
        System.out.println(float.class.getName());// float
        System.out.println(Object.class.getName());// java.lang.Object
        System.out.println(int[].class.getName());// [I
        System.out.println(int[][].class.getName());// [[I
        System.out.println(Object[].class.getName());// [Ljava.lang.Object;
        System.out.println(Object[][].class.getName());// [[Ljava.lang.Object;
        System.out.println(Runnable.class.getName());// java.lang.Runnable
        System.out.println(new double[0].getClass().getName());// [D
        System.out.println(new String[0].getClass().getName());// [Ljava.lang.String;
    }
}
