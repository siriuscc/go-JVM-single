public class StringTest {
    public static void main(String[] args) {

        // int x = 1;
        // String s = "abc" + x;
        // System.out.println(s);

        String s1 = "abc1";
        String s2 = "abc2";
        System.out.println(s1 == s2); // false

        int x3 = 1;
        String s4 = "abc" + x3;
        System.out.println(s1 == s4); // false
        System.out.println(s4);

        s4 = s4.intern();
        System.out.println(s1 == s4); // true
    }
}