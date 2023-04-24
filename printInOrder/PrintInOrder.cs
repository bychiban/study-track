using System;
using System.Threading;

namespace LeetCode
{
    public class PrintInOrder {
        private SemaphoreSlim _secondTask;
        private SemaphoreSlim _thirdTask;

        public PrintInOrder() {
            _secondTask = new SemaphoreSlim(0);
            _thirdTask = new SemaphoreSlim(0);
        }

        public void First(Action printFirst) {
            // printFirst() outputs "first". Do not change or remove this line.
            printFirst();
            _secondTask.Release();
        }

        public void Second(Action printSecond) {
            _secondTask.Wait();
            // printSecond() outputs "second". Do not change or remove this line.
            printSecond();
            _thirdTask.Release();
        }

        public void Third(Action printThird) {
            _thirdTask.Wait();
            // printThird() outputs "third". Do not change or remove this line.
            printThird();
        }
    }
}

