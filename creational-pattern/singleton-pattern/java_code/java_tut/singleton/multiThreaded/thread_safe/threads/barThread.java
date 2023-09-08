package singleton.multiThreaded.thread_safe.threads;

import singleton.multiThreaded.thread_safe.singleton;

public class barThread implements Runnable {

    @Override
    public void run() {
        singleton obj = singleton.getInstance("BAR");
        System.out.println(obj.value);
    }
}