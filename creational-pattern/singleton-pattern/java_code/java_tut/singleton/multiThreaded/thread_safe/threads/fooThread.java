package singleton.multiThreaded.thread_safe.threads;

import singleton.multiThreaded.thread_safe.singleton;

public class fooThread implements Runnable {

    @Override
    public void run() {
        singleton obj = singleton.getInstance("FOO");
        System.out.println(obj.value);
    }
}
