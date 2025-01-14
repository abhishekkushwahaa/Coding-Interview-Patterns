// LeetCode: https://leetcode.com/problems/two-sum/
// Example: 2Sum([2, 7, 11, 15], 9) => [0, 1]
let twoSum = function(nums, target) {
  for (let i = 0; i < nums.length; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      if (nums[i] + nums[j] === target) {
        return [i, j]
      }
    }
  }
};
// Time Complexity -  O(n^2)
