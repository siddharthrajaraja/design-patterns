package singleton.singleThreaded;

public class main {
    public static void main(String[] args) {

        singleton singleton1 = singleton.getInstance("FOO");
        singleton anothersingleton1 = singleton.getInstance("BAR");
        System.out.println(singleton1.value);
        System.out.println(anothersingleton1.value);
    }
}
