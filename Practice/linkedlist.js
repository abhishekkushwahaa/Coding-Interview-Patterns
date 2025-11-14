class Node {
  constructor(value) {
    this.value = value
    this.next = null
  }
}

function deleteKthNode(head, k) {
  let dummy = { value: -1, next: head }
  let slow = dummy
  let fast = dummy

  for (let i = 0; i < k; i++) {
    fast = fast.next
  }

  while (fast.next) {
    slow = slow.next
    fast = fast.next
  }

  // Delete node
  slow.next = slow.next.next;

  return dummy.next;
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

let list = new Node(2)
list.next = new Node(4)
list.next.next = new Node(6)
list.next.next.next = new Node(8)

let DeleteNode = deleteKthNode(list, 2)
printList(DeleteNode)
