function ListNode(val, next) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
}

/**
 * @param {ListNode} head
 * @return {number}
 */
var pairSum = function(head) {
    let max = 0;
    let current = head;
    let mid = head;
    let count = 0;
    let stack = [];
    while (current !== null) {
        count++
        current = current.next
        if (count % 2 === 0){
            stack.push(mid.val)
            mid = mid.next
        }
    }
    while (mid !== null){
        max = Math.max(max, mid.val+stack.pop())
        mid = mid.next
    }

    return max
};
let head = new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4, null))))

console.log(pairSum(head));

