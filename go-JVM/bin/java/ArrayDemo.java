
/**
 * 数组测试
 */
public class ArrayDemo {

    public void test() {
        int[][][] x = new int[3][4][5];
    }

    public static void main(String[] args) {

        int[] a1 = new int[10]; // newarray
        String[] a2 = new String[10]; // anewarray
        int[][] a3 = new int[10][10]; // multianewarray
        int x = a1.length; // arraylength
        a1[0] = 100; // istore
        int y = a1[0]; // iaload
        a2[0] = "abc"; // aastore
        String a = a2[0]; // aaload

        System.out.println(a1[0]);
        System.out.println(a2[0]);
        System.out.println(a3[0][0]);
        System.out.println(x);
    }
}