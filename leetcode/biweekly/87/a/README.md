### 本题视频讲解

见[【双周赛 87】](https://www.bilibili.com/video/BV1MT411u7fW)。

### 思路

把字符串日期转换成具体在这一年的第几天，也就是一个整数。

那么问题变成：给你两个闭区间，求这两个闭区间的交集区间中的整数个数，这等价于交集区间长度加一。

例如 $[1,3]$ 与 $[2,4]$ 的交集为 $[2,3]$，区间长度为 $3-2=1$，区间内整数个数为 $3-2+1=2$。

如何计算区间长度呢？知道了交集区间的左右端点，就知道了交集区间的长度：

- 左端点 $\textit{left}$ 等于两个区间左端点的最大值；
- 右端点 $\textit{right}$ 等于两个区间右端点的最小值。

那么区间长度为 $\textit{right}-\textit{left}$，区间内的整数个数为 $\textit{right}-\textit{left}+1$。

如果交集区间为空，即 $\textit{right}-\textit{left}+1<0$，此时答案为 $0$。

```py [sol1-Python3]
DAYS_SUM = list(accumulate((31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31), initial=0))

def calc_days(date: str) -> int:
    return DAYS_SUM[int(date[:2]) - 1] + int(date[3:])

class Solution:
    def countDaysTogether(self, arriveAlice: str, leaveAlice: str, arriveBob: str, leaveBob: str) -> int:
        left = calc_days(max(arriveAlice, arriveBob))  # 直接比较字符串再计算
        right = calc_days(min(leaveAlice, leaveBob))
        return max(right - left + 1, 0)  # 答案不能为负数
```

```py [sol1-Python3 API]
def calc_dt(date: str) -> datetime.datetime:
    return datetime.datetime.strptime(date, '%m-%d')  # 默认是 1900 年（平年）

class Solution:
    def countDaysTogether(self, arriveAlice: str, leaveAlice: str, arriveBob: str, leaveBob: str) -> int:
        left = calc_dt(max(arriveAlice, arriveBob))  # 直接比较字符串再计算
        right = calc_dt(min(leaveAlice, leaveBob))
        return max((right - left).days + 1, 0)  # 答案不能为负数
```

```java [sol1-Java]
class Solution {
    private static final int[] DAYS = new int[]{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};

    private int calcDays(String S) {
        var s = S.toCharArray();
        int day = (s[3] - '0') * 10 + s[4] - '0';
        int month = (s[0] - '0') * 10 + s[1] - '0';
        for (int i = 0; i < month - 1; ++i) // 除了最后一个月份，前面的月份直接累加
            day += DAYS[i];
        return day;
    }

    public int countDaysTogether(String arriveAlice, String leaveAlice, String arriveBob, String leaveBob) {
        int left = calcDays(arriveAlice.compareTo(arriveBob) > 0 ? arriveAlice : arriveBob); // 直接比较字符串再计算
        int right = calcDays(leaveAlice.compareTo(leaveBob) < 0 ? leaveAlice : leaveBob);
        return Math.max(right - left + 1, 0); // 答案不能为负数
    }
}
```

```cpp [sol1-C++]
int DAYS[] = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};

class Solution {
    int calc_days(string s) {
        int day = (s[3] - '0') * 10 + s[4] - '0';
        int month = (s[0] - '0') * 10 + s[1] - '0';
        for (int i = 0; i < month - 1; ++i) // 除了最后一个月份，前面的月份直接累加
            day += DAYS[i];
        return day;
    }

public:
    int countDaysTogether(string arriveAlice, string leaveAlice, string arriveBob, string leaveBob) {
        int left = calc_days(max(arriveAlice, arriveBob)); // 直接比较字符串再计算
        int right = calc_days(min(leaveAlice, leaveBob));
        return max(right - left + 1, 0); // 答案不能为负数
    }
};
```

```go [sol1-Go]
var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func calcDays(s string) (day int) {
	for _, d := range days[:s[0]&15*10+s[1]&15-1] { // 数字字符 &15 等同于 -'0'
		day += d
	}
	return day + int(s[3]&15*10+s[4]&15)
}

func countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob string) int {
	left := calcDays(max(arriveAlice, arriveBob)) // 直接比较字符串再计算
	right := calcDays(min(leaveAlice, leaveBob))
	return max(right - left + 1, 0) // 答案不能为负数
}

func max[T int | string](a, b T) T { if b > a { return b }; return a }
func min(a, b string) string { if b < a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$。仅用到若干额外变量。

---

欢迎关注[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)，高质量算法教学，持续输出中~

附：[每日一题·高质量题解精选](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
