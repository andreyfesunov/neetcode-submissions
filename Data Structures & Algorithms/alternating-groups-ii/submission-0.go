func numberOfAlternatingGroups(colors []int, k int) int {
    n,run,answer:=len(colors),0,0; for i:=0;i<n+k-1;i++ { if i>0 && colors[i%n]!=colors[(i-1)%n] { run++ } else { run=1 }; if i>=k-1 && run>=k { answer++ } }; return answer
}