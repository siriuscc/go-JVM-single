public class HashCodeDemo {
    public static void main(String[] args) {

        Object o1 = new Object();
        Object o2 = new Object();

        System.out.println(o1.hashCode());
        System.out.println(o1.hashCode());
        System.out.println(o2.hashCode());
    }
}