package main

import (
  "fmt"
)


func combine_mas(first_mas []float64, second_mas []float64) []float64 {
  var left_pos = 0
  var right_pos = 0
  var result_mas = []float64{}

  for true {
    if left_pos >= len(first_mas) {
      result_mas = append(result_mas, second_mas[right_pos:len(second_mas)]...)
      break
     } else if right_pos >= len(second_mas) {
      result_mas = append(result_mas, first_mas[left_pos:len(first_mas)]...)
      break
     } else {
          if first_mas[left_pos] < second_mas[right_pos] {
            result_mas = append(result_mas, first_mas[left_pos])
            left_pos++
          } else {
            result_mas = append(result_mas, second_mas[right_pos])
            right_pos++
          }
     }
  }
  return result_mas
}

func merge_sort(slice []float64) []float64  {
  var result_mas = []float64{}
  if len(slice) == 1 {
    fmt.Println(slice)
    return slice
  } else {
    var left_slice = slice[0:len(slice)/2]
    var right_slice = slice[len(slice)/2:len(slice)]
    left_slice = merge_sort(left_slice)
    right_slice = merge_sort(right_slice)
    result_mas = combine_mas(left_slice, right_slice)
  }
  return result_mas
}
func main() {
  fm := []float64{19, 21, 22}
  sm := []float64{3, 4, 5, 6}
  result_mas := combine_mas(fm, sm)
  fmt.Println(result_mas)

  unsorted_slice := []float64{12, 44, 3, 4, 5, 6}
  k := merge_sort(unsorted_slice)
  fmt.Println(k)
}
