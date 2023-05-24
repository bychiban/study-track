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
            var solution = new Solution();
            Console.WriteLine(solution.MaxScore(new int[]{1,3,3,2},new int[]{2,1,3,4},3));

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

