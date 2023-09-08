package singleton.singleThreaded;

public final class singleton {

    private static singleton instance;

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
            instance = new singleton(value);
        }
        return instance;
    }
}
