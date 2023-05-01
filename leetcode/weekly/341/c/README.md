下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

下文将 $\textit{word}$ 简记为 $s$。

# 方法一：考虑相邻字母

对于两个相邻字符 $x$ 和 $y$（$x$ 在 $y$ 左侧），使 $s$ 有效的话需要插入

$$
y-x-1
$$

个字母。

考虑到这可能是个负数，可以通过如下技巧转换在 $[0,2]$ 内：

$$
(y-x-1+3)\bmod 3
$$

- 例如 $x=\text{`a'},y=\text{`c'}$，则有 $(\text{`c'}-\text{`a'}+2)\bmod 3 = 1$，意思是需要补一个字母 $\text{`b'}$。
- 例如 $x=\text{`c'},y=\text{`a'}$，则有 $(\text{`a'}-\text{`c'}+2)\bmod 3 = 0$，无需补字母。

最后补齐开头的 $s[0]-\text{`a'}$，和结尾的 $\text{`c'}-s[n-1]$。这俩可以合并为 $s[0]-s[n-1]+2$。

```py [sol1-Python3]
class Solution:
    def addMinimum(self, s: str) -> int:
        ans = ord(s[0]) - ord(s[-1]) + 2
        for x, y in pairwise(map(ord, s)):
            ans += (y - x + 2) % 3
        return ans
```

```java [sol1-Java]
class Solution {
    public int addMinimum(String word) {
        var s = word.toCharArray();
        int ans = s[0] + 2 - s[s.length - 1];
        for (int i = 1; i < s.length; ++i)
            ans += (s[i] + 2 - s[i - 1]) % 3;
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int addMinimum(string s) {
        int ans = s[0] + 2 - s.back();
        for (int i = 1; i < s.length(); ++i)
            ans += (s[i] + 2 - s[i - 1]) % 3;
        return ans;
    }
};
```

```go [sol1-Go]
func addMinimum(s string) int {
	ans := int(s[0]) - int(s[len(s)-1]) + 2
	for i := 1; i < len(s); i++ {
		ans += (int(s[i]) - int(s[i-1]) + 2) % 3
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。

# 方法二：考虑 abc 的周期数

计算出 $\text{`abc'}$ 的周期数 $t$，那么有效字符串的长度为 $3t$，需要插入的字符个数为 $3t-n$。

对于两个相邻字符 $x$ 和 $y$（$x$ 在 $y$ 左侧），如果 $x<y$，那么 $x$ 和 $y$ 可以在同一个 $\text{`abc'}$ 周期内，否则一定不在。

所以 $t$ 就是 $x\ge y$ 的次数加一。

```py [sol2-Python3]
class Solution:
    def addMinimum(self, s: str) -> int:
        t = 1 + sum(x >= y for x, y in pairwise(s))
        return t * 3 - len(s)
```

```java [sol2-Java]
class Solution {
    public int addMinimum(String word) {
        var s = word.toCharArray();
        int t = 1;
        for (int i = 1; i < s.length; ++i)
            if (s[i - 1] >= s[i]) // 必须生成一个新的 abc
                ++t;
        return t * 3 - s.length;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    int addMinimum(string s) {
        int t = 1;
        for (int i = 1; i < s.length(); ++i)
            t += s[i - 1] >= s[i]; // 必须生成一个新的 abc
        return t * 3 - s.length();
    }
};
```

```go [sol2-Go]
func addMinimum(s string) int {
	t := 1
	for i := 1; i < len(s); i++ {
		if s[i-1] >= s[i] { // 必须生成一个新的 abc
			t++
		}
	}
	return t*3 - len(s)
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。
