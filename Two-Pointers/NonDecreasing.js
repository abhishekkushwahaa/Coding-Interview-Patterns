function checkPossibility(nums) {
  let count = 0

  for (let i = 0; i < nums.length - 1; i++) {
    if (nums[i] > nums[i + 1]) {
      count++
      if (count > 1) return false

      if (i === 0 || nums[i - 1] <= nums[i + 1]) {
        nums[i] = nums[i + 1]
      } else {
        nums[i + 1] = nums[i]
      }
    }
  }
  return true
};
