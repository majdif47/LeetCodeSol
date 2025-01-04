func isArraySpecial(nums []int) bool {
  for i := 0 ; i <= len(nums) - 2 ; i++ {
    if len(nums) == 1 {return true}
    if nums[i]%2 == nums[i+1] {
      return false
    }
  }
  return true
}


