// LeetCode: https://leetcode.com/problems/move-zeroes/
// Example: moveZeroes([0, 1, 0, 3, 12]) => [1, 3, 12, 0, 0]

function moveZeroes(nums) {
  let left = 0;
  let right = 0;

  while (right < nums.length) {
    if (nums[right] !== 0) {
      [nums[left], nums[right]] = [nums[right], nums[left]];
      left++;
    }
    right++;
  }
  return nums;
}
// const nums = [0, 1, 0, 3, 12];
// console.log(moveZeroes(nums));

// Time Complexity = O(n)
