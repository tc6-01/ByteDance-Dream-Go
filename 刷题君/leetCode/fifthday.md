# 第五天

## 移除元素

大致的思想就是不需要用额外的空间，在原来的数组上进行操作，将目的元素进行移除，返回移除后数组的长度或者移除后的数组。

因为数组的元素在内存空间中是连续的，不能直接删除，因此只能采取覆盖操作。

这里使用比较牛的双指针法（快慢指针）。

![27.移除元素-双指针法](https://tva1.sinaimg.cn/large/008eGmZEly1gntrds6r59g30du09mnpd.gif)

### 移除数组中的重复元素

​	移除重复元素的时候使用快慢指针，有一个很重要的前提就是需要确保数组是有序的，如果不是，可以先进行排序之后再进行移除。

​	其实数组中的重复元素移除，按照双指针来说，就是将有效的记录下来，然后直接返回有效的长度即可。但是实际问题总是要考虑边界问题。

```go
func removeDuplicates(nums []int) int {
    slow,n:=1,len(nums)
    if n==0{
        return 0
    }
    for i:=1;i<=n-1;i++{
        if nums[i-1]!=nums[i]{
            nums[slow]=nums[i]
            slow++
        }
    }
    return slow
}
```

### 移动0

​	与上一道题相比，无非就是多了后面无效元素的填充，将所有无效数字更新为0，这样就到达了目的

```go
func moveZeroes(nums []int)  {
    n:=len(nums)
    left:=0
    //记录所有非零数字
    for i:=0;i<=n-1;i++{
        if nums[i]!=0{
            nums[left]=nums[i]
            left++    
        }
    }
    for left<=n-1{
        nums[left]=0
        left++
    }
}
```

### 比较含退格的字符串

​	这道题目相对来讲会稍稍难一丢丢，但是我觉得确实学到东西了，比如人家用的条件判断，如果if了之后，后面使用else if就自动包括这个条件，然后新加的条件就与默认条件之间形成&关系。同样回归正题，只需要记录有效字符，当遇到退格符的时候，就删除前一个有效字符记录，最后对比有效字符即可

```go
func build(str string) string {
    s:=[]byte{}
    for i:=range(str){
        if i!='#'{
            s=append(s,i)
        }else if len(s)>0{
            s=s[:len(s)-1]
        }
    }
    return string(s)
}
func backspaceCompare(s, t string) bool {
    return build(s) == build(t)
}
```

### 总结

​	移除重复元素带来了数组解题的一种新思想，即元素覆盖，但是前提就是不用保留无效元素，这样才能使用该方法，否则另外开辟空间记录无效元素，其空间开销将与新开辟数组空间一致，费力不讨好的效果了属于是！