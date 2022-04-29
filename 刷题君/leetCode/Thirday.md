# Thirday

## 三数之和

```go
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var ret [][]int
    n := len(nums)
    for i := 0; i < n-2; i++ {
        //去重，防止两个前后元素一致
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        j, k := i+1, n-1
        for j < k {
           if sum := nums[i] + nums[j] + nums[k]; sum > 0 {
               k--
           }else if sum < 0 {
               j++
           }else {
               ret = append(ret, []int{nums[i], nums[j], nums[k]})
               j++
               for j < k && nums[j] == nums[j-1] {
                   j++
               }
               k--
               for j < k && nums[k] == nums[k+1] {
                  k-- 
               }
           }
        }
    }
    return ret
}
```

我是真的绷住了，为什么还可以内置排序的啊，不过整挺好，go语言内置了数字和字符串的排序sort.Ints()和sort.String()，先排序加双指针