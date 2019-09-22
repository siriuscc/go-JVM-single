import java.util.ArrayList;
import java.util.List;

/**
 * 自动装箱测试
 */
public class BoxTest {
    public static void main(String[] args) {
        typeErasure();
    }

    public static void boxTest() {
        List<Integer> list = new ArrayList<>();
        list.add(1);
        list.add(2);
        System.out.println(list.toString());

        for (int x : list) {
            System.out.println(x);
        }
    }

    public static void typeErasure() {
        List<Integer> list = new ArrayList<Integer>();
        list.add(1);
        System.out.println(list.get(0) + 999);
    }
}