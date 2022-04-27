# FirstDay

每天两道题（不得不说，我是真的废！！！！）

## 两数之和

​	首先会发现对go语言了解的不够透彻，不知道怎样去实现一些基础的内容。不过还是走了过来。

### Way1：暴力求解

​	其实就是一个一个试，用循环来得出最终的解，不过这样做时间复杂度好像有点高,嵌套了两层循环，不高才怪！！！

```go
func twoSum(nums []int, target int) []int {
    var result [] int
    for i:=0;i<len(nums);i++{
        for j:=i+1;j<len(nums);j++{
            if nums[i]+nums[j]==target{
                result = append(result,i,j)
            }   
        }
    }
    fmt.Println(result)
    return result
}
```

空间复杂度倒是挺小的，只有O(1)

### Way2:哈希表

​	不得不说，这个判断语句乍一看还是有点懵懵的，不过后来就懂了，go在返回值时增加了一个err值，或者就是说判断数组里面数是否存在的时候的布尔值。

```go
func twoSum(nums []int, target int) []int {
   table := map[int]int{}
   for i , x := range  nums{
       if p ,ok := table[target-x];ok{
           return []int {p,i}
       }
       table[x]=i
   }
   return nil
}
```

​	题解将具体数值的index作为value，首先要把全部数值存进hashmap，然后直接找目标-当前数值是否存在，存在之后取出value（即index），直接返回，如果实在没有就返回空。减少了时间复杂度变成O(n);增加了空间复杂度,变成O(n)

#### 优化

​	等后续.......

后续来了

```go
func twoSum(nums []int, target int) []int {
   m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if item, ok := m[target-nums[i]]; ok {
			return []int{item, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
```

​	怎么个优化法儿呢，其实就是直接利用索引，将数值放进需要用到的地方，省去了返回索引的时间。

## 两数相加

​	看到第一眼，我的想法是，先从链表中取出数字，然后在外面完成相加，最后再将值存进链表里，但是问题来了，如果说只有几位数还好，如果有很多位怎么办，难道一层一层取出来吗？换个角度想，好像有点太麻烦了，不过是可以实现的，但是真的实现出来之后，我就发现了，时间好久，意义不大，所以这里就给放弃了吧！

### 直接在每位的基础上进行相加

​	因为是逆序排列，那还不如直接进行相加，然后如果有进位就考虑进位，一起向下一位进行相加。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    //定义输出链表
    var result *ListNode = nil
    var parent *ListNode = nil
    up:=0
    //确保可以循环到最后一个元素
    for l1!=nil || l2!=nil ||up>0 {
        a,b:=0,0
        //取值进行运算
        if l1!=nil{
            a=l1.Val
            l1=l1.Next
        }
        if l2 != nil{
            b = l2.Val
            l2=l2.Next
        }
        t:=a+b+up    
        val:=t%10
        up=t/10
        //得到两个数值之后，进行链表的构造
        if result == nil{
            result=new(ListNode)
            result.Val = val
            parent=result
        }else{
            curr := new (ListNode)
            curr.Val = val
            parent.Next=curr
            parent=curr
        }
    }
    return result
}
```

## 无重复字符的最长子串

​	其实就是上学期学的串的匹配算法，美其名曰，滑动窗体。看到这一幕，只能说是欲哭无泪，上学期没学好的数据结构继续来折磨我了。

​	我记得的就只剩下移动匹配了。

​	基于之前用到过的map来确定数组指数，真的是越用越习惯---->>>>大概就是先判断（判断有无重复），后记录（记录字母的下一个位置）

```go
func lengthOfLongestSubstring(s string) int {
    n:=len(s)
    ans:=0
    m:=map[byte]int{}
    end,start:=0,0
    for end<n{
        // 判断下一个字母是否出现在hash表中
        if v,ok:=m[s[end]];ok{
            // 如果出现了重复，那就将start更新到上一个重复字母的下一个位置,或者是
            start=max(v,start)
        }
        // 如果没有重复，就记录下该字符的下一个字母位置
        m[s[end]]=end+1
        ans=max(end-start+1,ans)
        // 进行下一轮迭代
        end++
    }
    return ans
}
func max(x,y int) int {
    if x>y{
        return x
    }
    return y
}
```

然后就是一种抄来的方法，不过是真的高级，因为是字符本符嘛，所以直接申请128个字符，可以描述所有的字符（大概就是说ascii码对应的字符），也就是优化了判断过程（如果字母的value非0）

```go
func lengthOfLongestSubstring(s string) int {
    n:=len(s)
    if n<=1{
        return n
    }
    maxlen:=1
    end,start,window:=0,0,[128]int{}
    for end<n{
        endA:=s[end]
        //刚刚申请的新数组里面的所有内容都是0
        endI:=window[endA]
      	//如果end所在位置的字母重复，则将start放在重复字母的首次出现位置的下一位
        if endI > start{
            start=endI
        }
        //更新一下maxlen
        if maxlen<(end-start+1){
            maxlen=end-start+1
        }
        //记录每个字符出现位置的后一位
        window[endA] = end+1
        //end后移
        end++
    }
    return maxlen 
}
```



