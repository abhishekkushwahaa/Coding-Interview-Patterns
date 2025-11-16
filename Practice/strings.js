// TC - O(n) & SC - O(1)
function reverseWord(str) {
  // 1. Split string into string
  let words = str.trim().split(" ")

  // 2. Filters out empty strings (multiple space)
  words = words.filter(word => word.length > 0)

  // 3. Reverse string
  words.reverse()

  // 4. Return join words back into a string
  return words.join(" ")
}
let reverseWordString = reverseWord("I'm Balji!")
console.log(reverseWordString);

// TC - O(n) & SC - O(1)
function reverseWordWiseString(str) {
  let arr = str.split("");

  // Helper to reverse a portion in place
  function reverse(l, r) {
    while (l < r) {
      [arr[l], arr[r]] = [arr[r], arr[l]];
      l++;
      r--;
    }
    return;
  }

  // Step 1: Reverse entire char array
  let left = 0, right = arr.length - 1;
  reverse(left, right);

  // Step 2: Reverse each word individually
  let start = 0;
  for (let i = 0; i <= arr.length; i++) {
    if (i === arr.length || arr[i] === " ") {
      reverse(start, i - 1);
      start = i + 1;
    }
  }

  return arr.join("");
}
let reverseWordsString = reverseWordWiseString("I'm Balji!");
console.log(reverseWordsString);
