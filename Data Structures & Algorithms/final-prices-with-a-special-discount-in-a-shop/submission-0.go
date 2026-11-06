func finalPrices(prices []int) []int {
    answer:=append([]int(nil),prices...); stack:=[]int{}
    for i,price:=range prices {
        for len(stack)>0 && prices[stack[len(stack)-1]]>=price { answer[stack[len(stack)-1]]-=price; stack=stack[:len(stack)-1] }
        stack=append(stack,i)
    }
    return answer
}
