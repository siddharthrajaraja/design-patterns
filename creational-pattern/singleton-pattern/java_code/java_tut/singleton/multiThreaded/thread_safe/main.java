package singleton.multiThreaded.thread_safe;

import singleton.multiThreaded.thread_safe.threads.barThread;
import singleton.multiThreaded.thread_safe.threads.fooThread;

public class main {

    public static void main(String[] args) {
        Thread t1 = new Thread(new barThread());
        Thread t2 = new Thread(new fooThread());
        t1.start();
        t2.start();
    }
}
