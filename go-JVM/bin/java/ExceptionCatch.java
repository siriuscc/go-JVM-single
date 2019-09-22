
public class ExceptionCatch {

    public static void main(String[] args) {
        multiCatch();
        // try {
        // divZero();
        // } catch (Exception e) {
        // System.out.println("catch");
        // } finally {
        // System.out.println("Finally");

        // }
    }

    public static int finallyChange() {
        int n = 0;
        try {
            n = divZero();
            return n;
        } catch (Exception e) {
            System.out.println("catch");
        } finally {
            n = 999;
            System.out.println("Finally");
        }
        return n;
    }

    public static int returnFinally() {
        try {
            return divZero();
        } catch (Exception e) {
            System.out.println("catch");
            return 1;
        } finally {
            System.out.println("Finally");
            return -1;
        }

    }

    public static int divZero() {
        int n = 1000;
        return n / 0;
    }

    public static void multiCatch() {
        try {
            divZero();
        } catch (ArithmeticException e) {
            System.out.println("ArithmeticException catch");
        } catch (Exception e) {
            System.out.println("Exception catch");
        } finally {
            System.out.println("Finally");
        }
    }
}