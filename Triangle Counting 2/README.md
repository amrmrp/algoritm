

### 📄 Triangle Counting 2 — Problem Description (English)

# Triangle Counting 2

This problem is a variant of the original **"Triangle Counting 1"**, with the same logic but updated constraints.

## Time Limit
- 1 second

## Memory Limit
- 256 MB

## Problem Statement

You are given a positive integer `n`.

Your task is to count the number of **unordered triplets** `(a, b, c)` of positive integers such that:

1. They can form a valid triangle.
2. The sum of the sides is exactly `n`:
   - `a + b + c = n`

### Triangle Validity Condition

To form a valid triangle, the triplet `(a, b, c)` must satisfy the **triangle inequality**, i.e.:

- `a + b > c`
- `a + c > b`
- `b + c > a`

However, since the sum is fixed as `n`, and we are only counting unordered triplets where `a ≤ b ≤ c`, we only need to check the condition `a + b > c`.

### Unordered Triplets

Two triplets are considered the same if they contain the same numbers, regardless of order.

For example:
- `(2, 3, 2)` and `(2, 2, 3)` are considered the **same triplet** and should be counted only once.

## Input

A single line containing a positive integer `n`:

```
3 ≤ n ≤ 3000
```

## Output

Print a single integer — the number of valid unordered triplets `(a, b, c)` such that:
- `a + b + c = n`
- `(a, b, c)` can form a triangle.

## Examples

### Input 1
```
5
```

### Output 1
```
1
```

**Explanation:**  
The only valid triangle is `(1, 2, 2)`.

---

### Input 2
```
12
```

### Output 2
```
3
```

**Explanation:**  
Valid unordered triangles:
- `(2, 5, 5)`
- `(3, 4, 5)`
- `(4, 4, 4)`

