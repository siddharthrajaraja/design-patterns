package singleton.multiThreaded.non_thread_safe;

public final class singleton {

    private static singleton instance;

    public static Object multiThreaded;

    public String value;

    private singleton(String value) {

        try {
            Thread.sleep(1000);
        } catch (InterruptedException e) {

            e.printStackTrace();
        }
        this.value = value;
    }

    public static singleton getInstance(String value) {
        if (instance == null) {
            System.out.println("creating new instance of singleton class");
            instance = new singleton(value);
        }
        return instance;
    }
}
