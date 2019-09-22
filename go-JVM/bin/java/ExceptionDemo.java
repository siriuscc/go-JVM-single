public class ExceptionDemo {
    void cantBeZero(int i) throws Exception {
        if (i == 0) {
            throw new Exception();
        }
    }
}