public class StringReflect {
    public static void main(String[] args) {
        String a = "abc";
        Class<? extends String> aClass = a.getClass();
        System.out.println(aClass);
    }
}