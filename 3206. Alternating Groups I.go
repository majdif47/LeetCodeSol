package main
//runtime 0ms

func numberOfAlternatingGroups(colors []int) int {
  start := 0
  end := 0
  count := 0
  n := len(colors)
  for {
    if start >= n {
      break
    }
    end = start+2
    if end >= n {
      end = (start+2) - (n)
    }
    if start+1 >= n {
      if colors[end-1] != colors[end] && colors[end-1] != colors[start] {
        count++
      }
      }else {
        if colors[start+1] != colors[end] && colors[start+1] != colors[start] {
          count++
        }
    }
    start++
  }
  return count
}



