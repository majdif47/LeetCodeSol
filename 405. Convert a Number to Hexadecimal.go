func toHex(num int) string {
  if num == 0 {
    return "0"
  }
  n, mask := uint32(num), "0123456789abcdef"

  str := ""

  for n != 0 {
    str = string(mask [n& 0xF]) + str
    n >>= 4
  }
  return str
}
