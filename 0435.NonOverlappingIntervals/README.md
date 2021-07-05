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

