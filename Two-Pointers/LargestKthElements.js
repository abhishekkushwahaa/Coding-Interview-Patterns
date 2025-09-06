function LargestKthElement(nums, k) {
  nums.sort((a, b) => {
    if (a.length !== b.length) {
      return b.length - a.length
    } else {
      return b.localeCompare(a)
    }
  })
  return nums[k - 1]
}
