# [**410. 分割数组的最大值 Split Array Largest Sum**](https://leetcode.com/problems/split-array-largest-sum)

## 题目原文

### Problem

Given an array `nums` which consists of non-negative integers and an integer `m`, you can split the array into `m` non-empty continuous subarrays.

Write an algorithm to minimize the largest sum among these `m` subarrays.

### Example 1

```shell
Input: nums = [7,2,5,10,8], m = 2
Output: 18
Explanation:
There are four ways to split nums into two subarrays.
The best way is to split it into [7,2,5] and [10,8],
where the largest sum among the two subarrays is only 18.
```

### Example 2

```shell
Input: nums = [1,2,3,4,5], m = 2
Output: 9
```

### Example 3

```shell
Input: nums = [1,4,4], m = 3
Output: 4
```

**Constrains**:

- `1 <= nums.length <= 1000`

- `0 <= nums[i] <= 10^6`

- `1 <= m <= min(50, nums.length)`

## 解题思路

- 个人思路（递归+缓存+剪枝，时间复杂度 `O(n^m)`）：

  - 初始化缓存，维护一个 `n * n` 的二维数组 `cache`，其中 `cache[x][y]` 表示 `nums[x:y]` 范围内的数字之和。

  - 递归搜索每种分割方式，借助 `cache` 快速计算每个分割出的子数组元素之和，之后计算得到每种分割方式下的最大和，进一步求出最大和的最小值。

  - 剪枝：如果当前搜索路径上的已知最大和，超过了已知的最大和的最小值，则停止对这条路径的搜索。

- 个人思路（动态规划，时间复杂度 `O(n*n*m)`）：

  - 维护一个 `n * m` 的二维数组 `dp`，其中 `dp[y][x-1]` 表示 `nums[0:y]` 范围内分割后的子数组个数为 `x` 时的最大和的最小值，`dp[n-1][m-1]` 即为原问题的结果。

  - 初始化：

    - `dp[i][0]` 为 `nums[0:i]` 范围内的所有数字之和，分割后的子数组个数为 `1` （不分割）。

  - 递推：

    - `dp[y][x-1] = min(dp[y][x-1], max(dp[z-1][x-2], nums[z:y]))` ；

    - `dp[y][x-1]` 与 `dp[y-1][x-2]` 相比，追加了新元素 `nums[y]`，`nums[y]` 可以向前包含 `0` 个或者多个元素来组成一个新的分割后的子数组，即将 `nums[0:y]` 看作 `nums[0:z-1]` （其中包含 `x-1` 个子数组）和 `nums[z:y]`（其中包含 `1` 个子数组）的组合，`z` 在 `[x-1:y]` 范围内从右向左滑动，每滑动一次会产生一种新的分割方式，每种分割方式的最大和为 `max(dp[z-1][x-2], nums[z:y])`。

    - `z` 从右向左滑动时，`nums[z:y]` 是单调增加的，所以如果 `nums[z:y]` 超过了已知的 `dp[y][x-1]`，可以提前终止滑动。

