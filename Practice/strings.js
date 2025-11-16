function reverseString(str) {
  let arr = str.split("")
  let left = 0; let right = arr.length - 1;

  while (left < right) {
    // Swap Char
    [arr[left], arr[right]] = [arr[right], arr[left]]
    left++
    right--
  }

  return arr.join("")
}

let reverseStr = reverseString("balji")
console.log(reverseStr);
