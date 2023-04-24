using System;
using System.Threading;

namespace LeetCode
{
    public class PrintFooBarAlternately {
        private int n;
        private SemaphoreSlim _semaphoreFoo;
        private SemaphoreSlim _semaphoreBar;

        public PrintFooBarAlternately(int n) {
            this.n = n;
            _semaphoreFoo = new SemaphoreSlim(0);
            _semaphoreBar = new SemaphoreSlim(0);
            _semaphoreFoo.Release();
        }

        public void Foo(Action printFoo) {
            
            for (int i = 0; i < n; i++) {
                _semaphoreFoo.Wait();
                // printFoo() outputs "foo". Do not change or remove this line.
                printFoo();
                _semaphoreBar.Release();
            }
        }

        public void Bar(Action printBar) {
            
            for (int i = 0; i < n; i++) {
                _semaphoreBar.Wait();
                // printBar() outputs "bar". Do not change or remove this line.
                printBar();
                _semaphoreFoo.Release();
            }
        }
    }
}

