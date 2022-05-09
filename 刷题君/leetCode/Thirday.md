# Thirday

## 三数之和

```go
func threeSum(nums []int) [][]int {
    n := len(nums)
    sort.Ints(nums)
    ans := make([][]int, 0)
 
    // 枚举 a
    for first := 0; first < n; first++ {
        // 需要和上一次枚举的数不相同
        if first > 0 && nums[first] == nums[first - 1] {
            continue
        }
        // c 对应的指针初始指向数组的最右端
        third := n - 1
        target := -1 * nums[first]
        // 枚举 b
        for second := first + 1; second < n; second++ {
            // 需要和上一次枚举的数不相同
            if second > first + 1 && nums[second] == nums[second - 1] {
                continue
            }
            // 需要保证 b 的指针在 c 的指针的左侧
            for second < third && nums[second] + nums[third] > target {
                third--
            }
            // 如果指针重合，随着 b 后续的增加
            // 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
            if second == third {
                break
            }
            if nums[second] + nums[third] == target {
                ans = append(ans, []int{nums[first], nums[second], nums[third]})
            }
        }
    }
    return ans
}
```

我是真的绷住了，为什么还可以内置排序的啊，不过整挺好，go语言内置了数字和字符串的排序sort.Ints()和sort.String()，先排序加双指针

为什么要先排序，为了使每次可以确定两个元素不重复，因为a，b，c三个元素均是以1为步长进行递增，然后获取该位置的数字，保证一个位置两次使用的数字不是同一个数字（保证不重复的三元组）

然后固定第一个数字，当然要筛选出符合条件的第一个数字的位置，如果不行，直接跳过，条件一般都是不重复

双指针是为了减少暴力推荐的时间复杂度，将O（n^3）减少至O(n^2)

## 最接近的三数之和

```go
func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    n:=len(nums)
    result:=999
    for first:=0;first<n;first++{
        if first>0&&nums[first]==nums[first-1]{
            continue
        }
        second,third:=first+1,n-1
        for second<third{
            sum:=nums[first]+nums[second]+nums[third]
            if sum==target{
                return sum
            }
            if abs(sum-target)<abs(result-target){
                result=sum
            }
            if sum>target{
                third--
                for second<third&&nums[third]==nums[third+1]{
                    third--
                }
                
            }else{
                second++
                for second<third&&nums[second]==nums[second-1]{
                    second++
                }
            }
        }
    }
return result
}
func abs( x int)int{
    if x<0{
        return -1*x
    }
    return x
}
```

​	重点是判断条件的使用，以及后续指针的更新需要什么条件。

## 总结---双指针

首先使用有序数组，然后可以进行三个数或者两个数的便利，其中主要用于运算，尤其需要注意的是输出和判断条件，比如想要得到的内容，怎样将其抽象的得到一个表达式，并根据改表达式进行下一个指针的更新。

## 电话号码字母组合

```go
var phonemap map[string]string=map[string]string{
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}
var combinations []string

func letterCombinations(digits string) []string {
    if len(digits)==0{
        return []string{}
    }
    combinations=[]string{}
    //初始化
    getCorrectString(digits,0,"")
    return combinations
}
// 返回输入的电话号码对应的字符串
func getCorrectString(digits string,index int,combination string){
    if index==len(digits){
        combinations=append(combinations,combination)
    }else{
        digit:=string(digits[index])
        letters:= phonemap[digit]
        lettersCount:=len(letters)
        for i:=0;i<lettersCount;i++{
            getCorrectString(digits,index+1,combination+string(letters[i]))
        }
    }
}
```

使用dfs进行回溯得到拼接字符串

## 四数之和

```go
func fourSum(nums []int, target int)  [][]int {
    sort.Ints(nums)
    n := len(nums)
    quadruplets:=make([][]int,0)
    for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
        if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
            continue
        }
        for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
            if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
                continue
            }
            for left, right := j+1, n-1; left < right; {
                if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
                    quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
                    for left++; left < right && nums[left] == nums[left-1]; left++ {
                    }
                    for right--; left < right && nums[right] == nums[right+1]; right-- {
                    }
                } else if sum < target {
                    left++
                } else {
                    right--
                }
            }
        }
    }
    return quadruplets
}
```

外侧增加了一个循环，然后最内部的两个指针采用的策略是将其和与target进行比较得到left或right的指针移动。

这道题综合了三数之和与接近三数之和，将双指针的移动策略做了进一步综合。