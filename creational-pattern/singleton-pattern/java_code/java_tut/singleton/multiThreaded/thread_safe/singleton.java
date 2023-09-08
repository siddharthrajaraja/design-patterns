package singleton.multiThreaded.thread_safe;

public class singleton {

    public String value;

    // The field must be declared volatile so that double check lock would work
    // correctly.
    public static volatile singleton instance;

    private singleton(String value) {
        this.value = value;
    }

    public static singleton getInstance(String value) {
        if (instance != null)
            return instance;

        // synchronised keyword is used in DCL (double checked locking)
        synchronized (singleton.class) {
            if (instance == null) {
                instance = new singleton(value);
            }
            return instance;
        }
    }

}
