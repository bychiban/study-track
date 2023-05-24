using System;
using System.Collections.Generic;
using System.Linq;

namespace LeetCode{
    public class Solution {
        public long MaxScore(int[] nums1, int[] nums2, int k) {
            var n = nums1.Length;
            var pairs = new List<Tuple<int, int>>(n);
            for (int i = 0; i < n; ++i) {
                pairs.Add(Tuple.Create(nums1[i], nums2[i]));
            }
            pairs = pairs.OrderByDescending(t => t.Item2).ToList();
            
            var topKHeap = new PriorityQueue<int,int>(k);
            long topKSum = 0;
            for (int i = 0; i < k; ++i) {
                topKSum += pairs[i].Item1;
                topKHeap.Enqueue(pairs[i].Item1, pairs[i].Item1);
            }
            
            long answer = topKSum * pairs[k - 1].Item2;
            
            for (int i = k; i < n; ++i) {
                topKSum += pairs[i].Item1 - topKHeap.Dequeue();
                topKHeap.Enqueue(pairs[i].Item1, pairs[i].Item1);
                
                answer = Math.Max(answer, topKSum * pairs[i].Item2);
            }
            
            return answer;
        }
    }
}