# [**407. 接雨水II Trapping Rain Water II**](https://leetcode.com/problems/trapping-rain-water-ii)

## 题目原文

### Problem

Given an `m x n` integer matrix `heightMap` representing the height of each unit cell in a 2D elevation map, return the volume of water it can trap after raining.

### Example 1

```shell
Input: heightMap = [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]
Output: 4
Explanation: After the rain, water is trapped between the blocks.
We have two small pounds 1 and 3 units trapped.
The total volume of water trapped is 4.
```

### Example 2

```shell
Input: heightMap = [[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]
Output: 10
```

**Constrains**:

- `m == heightMap.length`

- `n == heightMap[i].length`

- `1 <= m, n <= 200`

- `0 <= heightMap[i][j] <= 2 * 10^4`

## 解题思路

- 个人思路：

  - 遍历 `heightMap` 中存在的所有高度值，对于每个高度值：

    - 遍历所有格子，如果该格子的高度大于当前高度值，则将该格子标记；

    - 从四条边开始，向上下左右四个方向搜索，能搜索到的格子都是当前高度值下无法接到雨水的，将该格子标记；

    - 遍历所有格子，在结果中累加上两轮处理后没有被标记的格子数量。

- 最优思路：

  - 维护一个优先级队列 `h` ，高度值较低的格子优先级更高。

  - 维护一个 `visited` 二维数组，与 `heightMap` 大小相同，用于标记格子是否被遍历过了。

  - 维护一个高度值 `height` ，表示当前遍历范围的最低边缘高度。

  - 从四条边开始，将四条边上的格子加入到 `h` 中，并且在 `visited` 中标记。

  - 当 `h` 不为空时，取出其中优先级最高（格子的高度最低）的格子：

    - 如果当前格子的高度大于 `height` ，则将 `height` 更新为当前格子的高度，此时高度值小于 `height` 且与当前边缘能够连通的所有格子都被遍历过了，只剩下高度大于 `height` 的格子、完全被高度大于 `height` 的格子包围并且高度小于 `height` 的格子；

    - 向上下左右四个方向搜索，如果搜索到的格子是合法的且未在 `visited` 中标记过的，则：

      - 将该格子在 `visited` 中标记；

      - 将该格子加入到 `h` 中；

      - 如果该格子被搜索到了且该格子的高度值小于 `height` ，则说明当前高度是当前格子能够接到的雨水的最大高度，将当前格子的高度值与 `height` 的差值累加到结果中（如果当前格子被其他高度更高的格子们完全包围，即当前格子能接到的雨水的最大高度应该更大，那么当前格子是不会被搜索到的，这里巧妙地利用了**能被搜索到**这个条件，来决定每个能接到雨水的格子的最大雨水高度）。

- 对比分析：

  - 个人思路和最优思路都是从四条边开始，向上下左右四个方向搜索，最优思路强在利用了高度值作为搜索顺序的优先级，从而能够直接计算出每个格子能接到雨水的最大高度，不用按照高度值分层搜索，避免了重复计算，跟"接雨水I"中单调栈的使用思想如出一辙。
