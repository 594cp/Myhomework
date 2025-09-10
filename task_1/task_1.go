package main
import (
    "fmt" 
    "strconv"
)
// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func singleNumber(nums []int) int {
    const MAX = 2*3*10000+10
    var x  [MAX]int 
    for i:=range nums{
        if x[nums[i]+3*10000]==0 {
            x[nums[i]+3*10000]++
        } else  {
            x[nums[i]+3*10000]--
        }
        
    }
    var ans int
    for i:=0; i<2*30000; i++{
        if x[i]==1{
            fmt.Println("i:",i)
            ans =int(i)-3*10000
        }
    }
     
    return ans;
}

// 9. 回文数 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
func isPalindrome(x int) bool {
    var find bool=true
    var str string = strconv.Itoa(x)
    fmt.Println("ff:",str)
    for i:=range str{
        fmt.Println("iff:",i)
        fmt.Println("ifr:",len(str)-i-1)
        if i>=len(str)-i-1{
            break
        } else if str[i]!=str[len(str)-i-1]{
            find  = false
            fmt.Println("i:",i)
            break
        } 
    }
    return find
}

// 20. 有效的括号 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
    var zuo [10010]byte
    top:=-1
    yes:=true
    fmt.Println("s:",s)
    for i:=range s{
        
        if (s[i]==')') || (s[i]=='}') || (s[i]==']') {
            if top<0  { 
                yes=false 
                break
            }
            if (zuo[top]=='(' && s[i]==')') || (zuo[top]=='{' && s[i]=='}') || (zuo[top]=='[' && s[i]==']') {
                top--
            } else {
                yes=false
                break
            }
        } else {
            top++
            fmt.Println("top:",top," ",byte(s[i]))
            zuo[top]=byte(s[i])
        }
    }
    if top>=0 {
       yes=false
    } 
    return yes
}


// 14. 最长公共前缀 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
    var ans string
    var x int =300
    for i:=range strs{
       if len(strs[i])<x{
          x=len(strs[i])
       }
    }
    find:=true
    fmt.Println("x:",x)
    for i:=0; i<x; i++{
        if len(strs)==0{
            break
        }
        if x==0 {
            break
        }
        fmt.Println("i:",i)
        //fmt.Println("str:", strs[0][i])
      
        same:=strs[0][i]
        fmt.Println("tr:",len(strs))
        for j:=0; j<len(strs); j++{
            if strs[j][i]!=same{
                find = false
                break
            }
        }
        if find==false {
            break
        }
        ans += string(same)
    }
    return ans; 
}


// 66. 加一 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
// 将大整数加 1，并返回结果的数字数组。
func plusOne(digits []int) []int {
    var ans []int
    can := false
    for i:=len(digits)-1; i>=0; i--{
        fmt.Println("i:",i)
        digits[i]++
        if digits[i]==10 {
            digits[i]=0
            can=true
        } else {
            can=false
        }
        if can==false {
               break
        }
    }
    fmt.Println("can:",can)
     fmt.Println("ssd:",digits)
    ans=digits
    if can {
        ans=append([]int{1},digits...)
    }
    return ans
}


// 26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，
// 使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在
// 原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指
// 针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j]
//  不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
func removeDuplicates(nums []int) int {
    M:=make(map[int]int)
    ans :=0  
    p := 0
    for i:=range nums{
        val,exist :=M[nums[i]]
        if exist ==false {
            M[nums[i]]=10
            nums[p]=nums[i]
            p++
            ans++
            fmt.Println(M)
            fmt.Println("p:",p," val=",val) 
        } else {
            M[nums[i]]=0
            continue
        }
    }
    return ans
}

// 56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数
// 组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与
// 切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
		return [][]int{}
	}
    var ret [][]int 
    var ans [40400]int
    for i := range intervals{
          l:=intervals[i][0]
          r:=intervals[i][1]
          for i:=l*2; i<=r*2; i++{
             ans[i]=1
          }
    }
    var zuo,you int
    for i:=0; i<20400; i++{
        if ans[i]==0 && ans[i+1]==1{
            zuo = (i + 1) / 2
        } 
        if ans[i]==1 && ans[i+1]==0{
            you=i/2
            //ret = append(ret,[][]int{{zuo,you}})
            ret = append(ret, []int{zuo, you})
        }
    }
    return ret
}

// 两数之和 
// 考察：数组遍历、map使用
// 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func twoSum(nums []int, target int) []int {
    var find bool
    var i,j int
    find = false
    for i=range nums {
        
        for j=i+1; j<len(nums); j++{
            if (nums[i]+nums[j]==target) {
                find = true
                break
            }
        }
        if find==true {
            
            break
        }
      
    }
    var re=[]int{}
    re = append(re,i,j)
    return re
}

// 主函数
func main() {
	fmt.Println("welcome to golang learing!")
}