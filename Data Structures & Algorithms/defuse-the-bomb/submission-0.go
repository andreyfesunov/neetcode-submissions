func decrypt(code []int, k int) []int {
    n:=len(code); answer:=make([]int,n); count,direction:=k,1; if k<0 { count,direction=-k,-1 }; for i:=0;i<n;i++ { for step:=1;step<=count;step++ { answer[i]+=code[(i+direction*step+n)%n] } }; return answer
}
