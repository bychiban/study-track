using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;

namespace LeetCode
{
    public class Program
    {
        public static async Task Main(string[] args)
        {
            var foo = new PrintInOrder();
            var tasks = new List<Task> {
                Task.Run(() => foo.First(() => {
                    Thread.Sleep(1000);
                    Console.Write("first"); })),
                Task.Run(() => foo.Second(() => Console.Write("second"))),
                Task.Run(() => foo.Third(() => Console.Write("third")))
            };

            await Task.WhenAll(tasks);
            Console.WriteLine();

            var fooBar = new PrintFooBarAlternately(3);
            tasks = new List<Task> {
                Task.Run(() => fooBar.Foo(() => {
                    Thread.Sleep(1000);
                    Console.Write("foo");
                })),
                Task.Run(() => fooBar.Bar(() => {
                    Thread.Sleep(2000);
                    Console.Write("bar");
                }))
            };
            await Task.WhenAll(tasks);
        }
    }
}

