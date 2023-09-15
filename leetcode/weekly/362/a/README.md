[视频讲解](https://www.bilibili.com/video/BV1U34y1N7Pe/)。

假设我们有一个数组，对于示例 1，我们可以把下标在 $[3,6]$ 的元素都加一，把下标在 $[1,5]$ 的元素都加一，下标在 $[4,7]$ 的元素都加一。

然后，我们来看看有多少个下标对应的元素值是大于 $0$ 的，这些下标就是题目要计算的，被任意区间覆盖的整数点。

那么，如何快速地「把区间内的数都加一」呢？

我之前在力扣上写过一篇文章：[【算法小课堂】差分数组](https://leetcode.cn/circle/discuss/FfMCgb/)。

根据这篇文章，用差分数组 $\textit{diff}$ 快速完成区间操作，然后求它的前缀和（恢复原数组），统计前缀和中有多少个数大于 $0$，即为答案。

```py [sol-Python3]
class Solution:
    def numberOfPoints(self, nums: List[List[int]]) -> int:
        max_end = max(end for _, end in nums)
        diff = [0] * (max_end + 2)
        for start, end in nums:
            diff[start] += 1
            diff[end + 1] -= 1
        return sum(s > 0 for s in accumulate(diff))
```

```java [sol-Java]
class Solution {
    public int numberOfPoints(List<List<Integer>> nums) {
        var diff = new int[102];
        for (var p : nums) {
            diff[p.get(0)]++;
            diff[p.get(1) + 1]--;
        }
        int ans = 0, s = 0;
        for (int d : diff) {
            s += d;
            if (s > 0) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfPoints(vector<vector<int>> &nums) {
        int diff[102]{};
        for (auto &p: nums) {
            diff[p[0]]++;
            diff[p[1] + 1]--;
        }
        int ans = 0, s = 0;
        for (int d: diff) {
            s += d;
            ans += s > 0;
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfPoints(nums [][]int) (ans int) {
	diff := [102]int{}
	for _, p := range nums {
		diff[p[0]]++
		diff[p[1]+1]--
	}
	s := 0
	for _, d := range diff {
		s += d
		if s > 0 {
			ans++
		}
	}
	return
}
```

```js [sol-JavaScript]
var numberOfPoints = function(nums) {
    const diff = new Array(102).fill(0);
    for (const p of nums) {
        diff[p[0]]++;
        diff[p[1] + 1]--;
    }
    let ans = 0, s = 0;
    for (const d of diff) {
        s += d;
        ans += s > 0;
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+\max\{\textit{end}_i\})$。其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(\max\{\textit{end}_i\})$。

## 相似题目

请看[【算法小课堂】差分数组](https://leetcode.cn/circle/discuss/FfMCgb/) 中的题单。
