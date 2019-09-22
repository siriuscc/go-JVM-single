

public class MyObject {

    public static int staticVar;
    public int instanceVar;

    public static void main(String[] args) {
        int x = 32768; // idc
        MyObject myObj = new MyObject();

        myObj.staticVar = x;
        x = myObj.staticVar;
        myObj.instanceVar = x;
        x = myObj.instanceVar;

        Object obj = myObj;
        if (obj instanceof MyObject) {
            myObj = (MyObject) obj;
            System.out.println(myObj.instanceVar);
        }
    }
}