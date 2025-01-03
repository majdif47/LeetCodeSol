package main 

func numberOfChild(n int, k int) int {
  x,y := k/(n-1), k%(n-1)
  if k <= n-1 {
    return k
  }
  if x%2 == 0 {
    return y
  }
  return n-1-y
}

