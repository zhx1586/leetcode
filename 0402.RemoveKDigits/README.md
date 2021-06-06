# [**402. 移掉K位数字 Remove K Digits**](https://leetcode.com/problems/remove-k-digits)

## 题目原文

### Problem

Given string num representing a non-negative integer `num`, and an integer `k`, return the smallest possible integer after removing `k` digits from `num`.

### Example 1

```shell
Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
```

### Example 2

```shell
Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
```

### Example 3

```shell
Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.
```

**Constrains**:

- `1 <= k <= num.length <= 10^5`

- `num` consists of only digits.

- `num` does not have any leading zeros except for the zero itself.

## 解题思路

- 个人思路：

  - 采用贪心的思想，每次从前 `k+1` 个数字中寻找第一个最小数字，将其作为结果中的第一位数字。

  - 第一位数字确定后，剩余数字又是一个原问题的子问题，递归处理子问题直到移除完所有 `k` 个数字。

- 最优思路：

  - 采用单调栈的思想。

  - 遍历所有数字：

    - 如果栈顶元素大于新遍历到的元素，且当前已移除的数字数量小于 `k` 个，则将栈顶元素出栈，新元素入栈，已移除的数字数量加一。

    - 如果栈顶元素不大于新遍历到的元素，或者当前已移除的数字数量等于 `k` 个，则将新元素入栈。

- 耗时原因：

  - 单调栈的思路不会重复遍历元素，贪心的思路每次从 `k+1` 个元素中寻找最小值会重复遍历最小值之后那一部分元素。

