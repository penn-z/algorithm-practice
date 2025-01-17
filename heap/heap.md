[TOC]

# 堆化及堆的操作

## 堆化

### 从前往后处理数组数据，从下往上堆化

TODO: ...


### 从后往前处理数组数据，从上往下堆化
1.  从最后一个非叶子节点开始，依次堆化

![堆化1](https://static001.geekbang.org/resource/image/50/1e/50c1e6bc6fe68378d0a66bdccfff441e.jpg)

![堆化2](https://static001.geekbang.org/resource/image/aa/9d/aabb8d15b1b92d5e040895589c60419d.jpg)

```java
private static void buildHeap(int[] a, int n) {
  for (int i = n/2; i >= 1; --i) {
    heapify(a, n, i);
  }
}

// 大顶堆
private static void heapify(int[] a, int n, int i) {
  while (true) {
    int maxPos = i;
    if (i*2 <= n && a[i] < a[i*2]) maxPos = i*2;
    if (i*2+1 <= n && a[maxPos] < a[i*2+1]) maxPos = i*2+1;
    if (maxPos == i) break;
    swap(a, i, maxPos);
    i = maxPos;
  }
}
```

## 排序
1. 建堆结束之后，数组中的数据已经是按照大顶堆的特性来组织的。数组中的第一个元素就是堆顶，也就是最大的元素。我们把它跟最后一个元素交换，那最大元素就放到了下标为 n 的位置。
2. 当堆顶元素移除之后，我们把下标为 n 的元素放到堆顶，然后再通过堆化的方法，将剩下的 n−1 个元素重新构建成堆。堆化完成之后，我们再取堆顶的元素，放到下标是 n−1 的位置，一直重复这个过程，直到最后堆中只剩下标为 1 的一个元素，排序工作就完成了

![排序](https://static001.geekbang.org/resource/image/23/d1/23958f889ca48dbb8373f521708408d1.jpg)

```java
// n表示数据的个数，数组a中的数据从下标1到n的位置。
public static void sort(int[] a, int n) {
  buildHeap(a, n);
  int k = n;
  while (k > 1) {
    swap(a, 1, k);
    --k;
    heapify(a, k, 1);
  }
}
```

***

# 堆的应用


## 堆

| 堆的用法     | 时间复杂度 | 空间复杂度 |
| ------------ | ---------- | ---------- |
| 创建堆       | O(N)       | O(N)       |
| 插入元素     | O(logN)    | O(1)       |
| 获取堆顶元素 | O(1)       | 0(1)       |
| 删除堆顶元素 | O(logN)    | O(1)       |
| 获取堆长度   | O(1)       | O(1)       |



## TopK问题

eg: 求最大的K个元素;

### 解法一

1. 创建一个==最大堆==
2. 将所有元素都加到==最大堆==中;
3. 通过『边删除边遍历』方法，讲堆顶元素删除，并将它保存到结果集T中;
4. 重复3步骤==K==次，直到取出前K个最大元素;



### 解法二

1. 创建一个大小为==K==的==最小堆==;

2. 依次将元素添加到==最小堆==中;

3. 当最小堆的元素个数达到==K==时，将当前元素与堆顶元素进行对比:

   >1. 如果当前元素小于堆顶元素，则放弃当前元素，继续进行下一个元素;
   >2. 如果当前元素大于堆顶元素，则删除堆顶元素，将当前元素加入到最小堆中

4. 重复步骤2和步骤3，知道所有元素遍历完毕;

5. 此时最小堆中的k个元素就是前==K==个最大元素。





## The Kth问题

eg: 求第K大元素



### 解法一

1. 创建一个==最大堆==;
2. 将所有元素加入到==最大堆==中;
3. 通过『边删除边遍历』方法，将堆顶元素删除;
4. 重复3步骤==K==次，直到获取到第==K==个最大的元素;





### 解法二

1. 创建一个大小为==K==的==最小堆==;

2. 依次将元素添加到==最小堆==中：

3. 当==最小堆==的元素达到==K==个时，将当前元素与堆顶元素进行对比:

   >1. 如果当前元素小于堆顶元素，则放弃当前元素，继续进行下一个元素;
   >
   >2. 如果当前元素大于堆顶元素，则删除堆顶元素，将当前元素加入到==最小堆==中

4. 重复步骤2和步骤3，直到所有元素遍历完毕

5. 此时==最小堆==中的堆顶元素就是第==K==个最大的元素。



