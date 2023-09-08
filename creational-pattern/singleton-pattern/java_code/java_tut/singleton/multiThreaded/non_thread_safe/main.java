package singleton.multiThreaded.non_thread_safe;

import singleton.multiThreaded.non_thread_safe.threads.barThread;
import singleton.multiThreaded.non_thread_safe.threads.fooThread;

public class main {

    public static void main(String[] args) {
        Thread t1 = new Thread(new fooThread());
        Thread t2 = new Thread(new barThread());
        t1.start();
        t2.start();
    }
}
