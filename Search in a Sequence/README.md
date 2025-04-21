
---

## 🔍 **Search in a Sequence 1**

**Time Limit:** 1 second  
**Memory Limit:** 64 MB

### 📝 Problem Statement

You are given a sequence of `n` integers:  
`a₁, a₂, ..., aₙ`.

You need to answer `q` queries. For the `i`-th query, an integer `xᵢ` is given, and you must print the **number of elements in the sequence** that are **strictly less than `xᵢ`**.

---

### 📥 Input

- The first line contains two integers `n` and `q`  
  (`1 ≤ n, q ≤ 1000`) — the number of elements in the sequence and the number of queries.
- The second line contains `n` integers `a₁, a₂, ..., aₙ`  
  (`1 ≤ aᵢ ≤ 1,000,000`) — the elements of the sequence.
- Each of the next `q` lines contains a single integer `xᵢ`  
  (`1 ≤ xᵢ ≤ 1,000,000`) — the number from the `i`-th query.

---

### 📤 Output

For each query, print a single line containing the answer — the number of elements in the sequence that are less than `xᵢ`.

---

### 📌 Sample Input 1
```
2 3
1 2
1
2
3
```

### ✅ Sample Output 1
```
0
1
2
```

---

### 📌 Sample Input 2
```
3 3
4 1 4
2
5
4
```

### ✅ Sample Output 2
```
1
3
1
```

---

### Help
```
initial n,q:
5 5

Sequence n:
1 2 3 4 6

q i query:
1
2
3
4
5
6

```

### 💡 Explanation

- In the first query of Sample 1, no number is less than 1.
- In the second query, only one number (1) is less than 2.
- In the third query, both 1 and 2 are less than 3.

--- 
