# Singleton Design Pattern

* Having a unique instance of a type in the entire program.

## Objectives

---

1. A single,shared value of some particular type
2. Restrict object creation of some type to a single unit along the entire program.

> Later we have Singleton implementation for Multithreading.

## Where is this used?

--- 

1. Database connection to make every query
2. Open SSH connection to a server to do a few tasks, and don't want to reopen the connection for each task
3. If you need to limit the access to some variable or space, you use Singleton as door to this variable
4. If you need to limit the number of calls to some places, you need a singleton instance to make the calls in the
   accepted window