# 第四天

## 二分法

​	有一说一，人家这代码随想录是真的香，今天刷了二分法的题。

从简单到复杂，场景依次变得复杂

### 有序无重复数组

#### 查找数值指定位置

​	作为最简单的数组，只需要进行查找，返回对应的索引即可，然后就是边界问题，因为要确定判断区间！--》循环不变量原则。

```go
n:=len(nums)
ans,left,right:=-1,0,n-1
for left<=right{
    //动态确定中间位置
    mid:=left+(right-left)>>1
    if nums[mid]==target{
        ans=mid
    }else if nums[mid]>target{
        right=mid-1
    }else{
        left=mid+1
    }
}
return ans
```

#### 搜索插入位置

相对于上一个而言，此处增加的是有些没有数值的数字在数组中的位置，当然也不会有重复的数组。

这个时候就要做一些简单的处理,比如在遇到相等的时候继续循环，直到没有相等出现，或者在出现小于等于的位置记录一下答案，

```go
func searchInsert(nums []int, target int) int {
    // 使用左闭右闭区间进行二分查找
    n:=len(nums)
    left,right:=0,n-1
    ans:=n
    for left<=right{
        m:=left+(right-left)/2
        if nums[m]>=target{
            ans=m
            right=m-1
        }else{
            left=m+1
        }
    }
    return ans
}
```

### 有序重复数组

​	在一个重复数组里，进行目标值开始位置的索引与结束位置的索引的查找，这个时候可以采用改变答案的记录来修改左右边界，所以这个时候编写两个用来确定数字的函数来进行左右边界的确定，然后在主函数中确认返回值，然后进行反馈值的处理。

```go
func searchRange(nums []int, target int) []int {
    leftmost := getLIndex(nums, target)
    //该种情况为不相等情况
    if leftmost == len(nums) || nums[leftmost] != target {
        return []int{-1, -1}
    }
    rightmost := getRIndex(nums, target)
    return []int{leftmost, rightmost}
}
func getLIndex(nums []int,target int)int{
    ans,left,right:=0,0,len(nums)-1
    for left<=right{
        mid:=left+(right-left)>>1
        if nums[mid]<target{
            left=mid+1
        }else{
            ans=mid
            right=mid-1
        }
    }
    return ans
}

func getRIndex(nums []int,target int)int{
    ans,left,right:=0,0,len(nums)-1
    for left<=right{
        mid:=left+(right-left)>>1
        if nums[mid]<=target{
            left=mid+1
            ans=mid
        }else{
            right=mid-1
        }
    }
    return ans
}
```

### 具体应用问题

#### 求一个平方根的整数部分，不用求根函数？？

换种思路其实就是确定出来一个区间，将这个区间里的所有数字都试一遍，来确认出平方根的整数部分。

```go
func mySqrt(x int) int {
    ans,left,right:=0,0,x
    for left<=right{
        mid:=left+(right-left)>>1
        if mid*mid<=x{
            ans=mid
            left=mid+1
        }else{
            right=mid-1
        }
    }
    return ans
}
```

#### 有效的完全平方数

如果该数值可以开平方就返回true，否则是false，那么这里就引用了之前提到的二分法，采用对半的方法来试，

```go
func isPerfectSquare(num int) bool {
    ok:=false
    l,r:=0,(1<<31)-1
    for l<=r{
        m:=l+(r-l)>>1
        if m*m<num{
            l=m+1
        }else if m*m>num{
            r=m-1
        }else{
            ok=true
            return ok
        }
    }
    return ok
}
```

