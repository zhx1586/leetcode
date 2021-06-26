# [**424. 替换后的最长重复字符 Longest Repeating Character Replacement**](https://leetcode.com/problems/longest-repeating-character-replacement)

## 题目原文

### Problem

You are given a string `s` and an integer `k`. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most `k` times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

### Example 1

```shell
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
```

### Example 2

```shell
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
```

**Constrains**:

- `1 <= s.length <= 10^5`

- `s` consists of only uppercase English letters.

- `0 <= k <= s.length`

## 解题思路

- 最优思路（滑动窗口）：

  - 维护一个 `s[lo:hi]` 范围内的窗口，并维护一个 `counts` 数组记录窗口内每个字符出现的次数；

  - 判断当前窗口内，是否有一个字符出现的次数加上 `k` 大于等于 窗口长度（即替换 `k` 个字符后，窗口内的字符全部为该字符）:

    - 如果满足条件，则用当前窗口长度与返回值 `ret` 比较并更新，并且将窗口的右端点向右滑动一个单位；

    - 如果不满足条件且当前窗口长度为 `1` ，则将窗口的右端点向右滑动一个单位；

    - 如果不满足条件且当前窗口长度大于 `1` ，则将窗口的左端点向左滑动一个单位。

  - 重复窗口的滑动过程直到窗口的右端点到达 `s` 的尾部。

- 个人理解：

  - 滑动窗口可以看作是一种特殊的动态规划，或者看作是动态规划的一种特殊实现方式，与普通的动态规划相比：

    - 滑动窗口更省时间，利用了子问题之间的关系，跳过了一部分没有必要计算处理的子问题；

    - 滑动窗口更省空间，只需要存储当前窗口的状态，没有必要存储所有子问题的状态。
