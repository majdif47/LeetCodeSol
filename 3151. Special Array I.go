package main
import "fmt"



func isArraySpecial(nums [4]int) bool {
  for i := 0 ; i <= len(nums) - 2 ; i++ {
    if len(nums) == 1 {return true}
    if nums[i]%2 == nums[i+1] {
      fmt.Println(nums[i],nums[i+1])
      return false
    }
  }
  return true
}


func main() {
  arr := [4]int{4,3,1,6}
  isArraySpecial(arr)
}
