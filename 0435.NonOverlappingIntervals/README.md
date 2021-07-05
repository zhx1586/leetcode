# [**435. 无重叠区间 Non-overlapping Intervals**](https://leetcode.com/problems/non-overlapping-intervals)

## 题目原文

### Problem

Given an array of intervals `intervals` where `intervals[i] = [start_i, end_i]`, return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

### Example 1

```shell
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.
```

### Example 2

```shell
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.
```

### Example 3

```shell
Input: intervals = [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
```

**Constrains**:

- `1 <= intervals.length <= 2 * 10^4`

- `intervals[i].length == 2`

- `-2 * 10^4 <= start_i < end_i <= 2 * 10^4`

## 解题思路

- 最优思路：（排序+贪心）

  - 对 `intervals` 进行排序，`start_i` 较小的在前，`start_i` 相等时 `end_i` 较小的在前。

  - 遍历排序后的 `intervals`：

    - 如果 `intervals[curr][0] < intervals[last][1]`，则当前区间与前一个区间有重叠，必须删除一个，计数值 `count` 加一；

      - 如果 `intervals[curr][1] < intervals[last][1]`，则删除 `last` 对应的区间，否则删除 `curr` 对应的区间。

    - 如果没有重叠，则不需要删除，将 `last` 更新为 `curr`，继续遍历。

  - 核心思想就是在重叠时尽可能保留右端点较小的区间。
