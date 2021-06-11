# [**406. 根据身高重建队列 Queue Reconstruction by Height**](https://leetcode.com/problems/queue-reconstruction-by-height)

## 题目原文

### Problem

You are given an array of people, `people`, which are the attributes of some people in a queue (not necessarily in order). Each `people[i] = [hi, ki]` represents the `ith` person of height `hi` with **exactly** `ki` other people in front who have a height greater than or equal to `hi`.

Reconstruct and return the queue that is represented by the input array `people`. The returned queue should be formatted as an array `queue`, where `queue[j] = [hj, kj]` is the attributes of the `jth` person in the queue (`queue[0]` is the person at the front of the queue).

### Example 1

```shell
Input: people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
Output: [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
Explanation:
Person 0 has height 5 with no other people taller or the same height in front.
Person 1 has height 7 with no other people taller or the same height in front.
Person 2 has height 5 with two persons taller or the same height in front, which is person 0 and 1.
Person 3 has height 6 with one person taller or the same height in front, which is person 1.
Person 4 has height 4 with four people taller or the same height in front, which are people 0, 1, 2, and 3.
Person 5 has height 7 with one person taller or the same height in front, which is person 1.
Hence [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] is the reconstructed queue.
```

### Example 2

```shell
Input: people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
Output: [[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
```

**Constrains**:

- `1 <= people.length <= 2000`

- `0 <= hi <= 10^6`

- `0 <= ki < people.length`

- It is guaranteed that the queue can be reconstructed.

## 解题思路

- 个人思路：

  - 对输入的 `people` 进行排序，`h` 较小的排在前面，`h` 相同时 `k` 较小的排在前面。

  - 初始化一个返回结果 `ret`，长度与 `people` 相同。

  - 对于排序后 `people` 中的每个元素，将其放置在 `ret` 中从前到后数的第 `k+1` 个位置中，计数方式为：

    - 如果当前位置为空位，则计数加一；

    - 如果当前位置不为空位，当前位置的元素的 `h` 更大，则计数加一；

    - 如果当前位置不为空位，当前位置的元素的 `h` 更小或者相等，则计数不变。

- 最优思路：

  - 对输入的 `people` 进行排序，`h` 较大的排在前面，`h` 相同时 `k` 较小的排在前面。

  - 初始化一个链表结构的 `ret`。

  - 对于排序后 `people` 中的每个元素，将其放置在 `ret` 中从前到后数的第 `k` 个位置中。

  - 将链表结构的 `ret` 转换为数组结构。

- 对比分析：

  - 最优思路更加清晰直观，个人思路太绕了，而且放置每个元素时要遍历的元素数更多，更加耗时。

