class Node {
  constructor(value) {
    this.value = value
    this.next = null
  }
}

function midPoint(head) {

  if (!head) return null;

  let slow = head;
  let fast = head;

  while (fast !== null && fast.next !== null) {
    slow = slow.next
    fast = fast.next.next
  }

  return slow;
}

function printList(head) {
  let current = head
  let result = []

  while (current !== null) {
    result.push(current.value)
    current = current.next
  }
  console.log(result.join(" -> "))
}

let list = new Node(2);
list.next = new Node(3);
list.next.next = new Node(5);
list.next.next.next = new Node(7);

let midPointofList = midPoint(list)
// printList(midPointofList.result);
console.log(midPointofList.value)
