# [**416. 分割等和子集 Partition Equal Subset Sum**](https://leetcode.com/problems/partition-equal-subset-sum)

## 题目原文

### Problem

Given a **non-empty** array `nums` containing **only positive integers**, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

### Example 1

```shell
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].
```

### Example 2

```shell
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.
```

**Constrains**:

- `1 <= nums.length <= 200`

- `1 <= nums[i] <= 100`

## 解题思路

- 最优思路：

  - 计算 `nums` 中所有数字的和 `sum`，如果 `sum` 为奇数，则直接返回 `false` 。

  - 维护一个 `(len(nums)+1)*(sum+1)` 的二维数组 `dp`，其中 `dp[i+1][j]` 表示 `nums` 中的前 `i` 个数字能否组合成 `j`，最终结果为 `dp[len(nums)][sum/2]` ，即 `nums` 中的全部数字能否组合成 `sum/2`。

  - 初始条件：`dp[0][0] = true`，即不使用 `nums` 中的任何一个数字，只能组合成 `0` 。

  - 递推条件：`dp[i+1][j] = dp[i][j-nums[i]] || dp[i][j]`，即

    - 如果 `nums` 中的前 `i-1` 个数字能组合成 `j`，那么前 `i` 个数字也能组合成 `j` （相当于 `nums[j]` 无贡献）；

    - 如果 `nums` 中的前 `i-1` 个数字能组合成 `j-nums[i]`，那么前 `i` 个数字能组合成 `j` （相当于 `nums[i]` 有贡献）。
