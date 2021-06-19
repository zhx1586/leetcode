# [**401. 强密码检验器 Strong Password Checker**](https://leetcode.com/problems/strong-password-checker)

## 题目原文

### Problem

A password is considered strong if the below conditions are all met:

- It has at least `6` characters and at most `20` characters.

- It contains at least one lowercase letter, at least one uppercase letter, and at least one digit.

- It does not contain three repeating characters in a row (i.e., "`...aaa...`" is weak, but "`...aa...a...`" is strong, assuming other conditions are met).

Given a string `password`, return the minimum number of steps required to make `password` strong. if `password` is already strong, return `0`.

In one step, you can:

- Insert one character to `password`,

- Delete one character from `password`, or

- Replace one character of `password` with another character.

### Example 1

```shell
Input: password = "a"
Output: 5
```

### Example 2

```shell
Input: password = "aA1"
Output: 3
```

### Example 3

```shell
Input: password = "1337C0d3"
Output: 0
```

**Constrains**:

- `1 <= password.length <= 50`

- `password` consists of letters, digits, dot `'.'` or exclamation mark `'!'`.

