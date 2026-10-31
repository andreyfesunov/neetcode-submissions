func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    totals:=make([]int,1001);for _,row:=range nums1{totals[row[0]]+=row[1]};for _,row:=range nums2{totals[row[0]]+=row[1]};answer:=[][]int{};for id:=1;id<=1000;id++{if totals[id]>0{answer=append(answer,[]int{id,totals[id]})}};return answer
}
