class Node {
  constructor(value) {
    this.value = value
    this.next = null
  }
}

function addTwoNumbers(l1, l2) {
  let dummy = new Node(0)
  let current = dummy
  let carry = 0

  while (l1 !== null || l2 !== null || carry > 0) {
    let x = (l1 !== null) ? l1.value : 0
    let y = (l2 !== null) ? l2.value : 0

    let sum = x + y + carry;

    carry = Math.floor(sum / 10)

    current.next = new Node(sum % 10)
    current = current.next

    if (l1 !== null) l1 = l1.next
    if (l2 !== null) l2 = l2.next
  }
  return dummy.next
}

function PrintList(head) {
  let current = head
  let value = []
  while (current !== null) {
    value.push(current.value)
    current = current.next
  }
  console.log(value.join(" -> "))
}

// Example
let l1 = new Node(2);
l1.next = new Node(4);
l1.next.next = new Node(3);

let l2 = new Node(5);
l2.next = new Node(6);
l2.next.next = new Node(4);

let result = addTwoNumbers(l1, l2);
PrintList(result);
