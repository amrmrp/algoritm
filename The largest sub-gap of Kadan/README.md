
---

# Maximum Subarray 1

## Problem Statement
You are given a sequence of `n` integers:
```
a1, a2, ..., an
```
You need to select a **contiguous subarray** (a subarray consists of consecutive elements) such that the sum of its elements is maximized.

Write a program to find and print the maximum possible subarray sum.

---

## Input
- The first line contains a single integer `n` — the number of elements in the sequence.  
  `1 ≤ n ≤ 100`
- The second line contains `n` integers `a1, a2, ..., an`, separated by spaces.  
  `-10^9 ≤ ai ≤ 10^9`

---

## Output
- Print a single integer — the maximum sum of any subarray.

---

## Sample Input and Output

### Sample 1
Input:
```
6
-7 3 -1 2 -4 3
```
Output:
```
4
```
(The best subarray is `[3, -1, 2]` with a sum of `4`.)

---

### Sample 2
Input:
```
3
-1 -2 -3
```
Output:
```
-1
```
(Only the largest single number, `-1`, can be chosen.)

---

## Notes
The optimal solution uses **Kadane's Algorithm**, which finds the maximum subarray sum in linear time `O(n)`.

---

