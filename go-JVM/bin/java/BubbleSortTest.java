import java.util.Arrays;

/**
 * 冒泡排序，测试数组操纵
 */
public class BubbleSortTest {
    public static void main(String[] args) {
        int[] arr = { 32, 63, 22, 11, 44, 234, 12, 34, 85, 1 };
        bubbleSort(arr);

        System.out.println(Arrays.toString(arr));

    }

    private static void bubbleSort(int[] arr) {
        for (int i = 0; i < arr.length; ++i) {
            for (int j = i + 1; j < arr.length; ++j) {
                if (arr[i] < arr[j]) {
                    swap(arr, i, j);
                }
            }
        }
    }

    private static void swap(int[] arr, int i, int j) {
        int tmp = arr[i];
        arr[i] = arr[j];
        arr[j] = tmp;
    }
}
