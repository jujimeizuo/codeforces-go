下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

### 前置知识：二叉树的 BFS（层序遍历）

见[【基础算法精讲 13】](https://www.bilibili.com/video/BV1hG4y1277i/)。

### 提示 1

下文将具有相同父节点的节点互称为**兄弟节点**。

对于一个节点 $x$ 来说，它的所有**堂兄弟节点**值的和，等价于 $x$ 这一层的所有节点值之和，减去 $x$ 及其兄弟节点的值之和。

例如样例 1：

- $4$ 的左右儿子的节点值，都被更新成了 $7$，也就是左右儿子这一层的节点值之和 $1+10+7=18$，减去 $4$ 的左右儿子的节点值之和 $1+10=11$，得到 $7$。
- $9$ 的右儿子的节点值，被更新成了 $11$，也就是右儿子这一层的节点值之和 $1+10+7=18$，减去 $9$ 的右儿子的节点值 $7$，得到 $11$。

### 提示 2

怎么实现呢？

用 BFS 遍历二叉树，对于每一层：

- 首先，遍历当前层的每个节点，通过节点的左右儿子，计算下一层的节点值之和 $\textit{nextLevelSum}$；
- 然后，再次遍历当前层的每个节点 $x$，计算 $x$ 的左右儿子的节点值之和 $\textit{childrenSum}$，更新 $x$ 的左右儿子的节点值为 $\textit{nextLevelSum}-\textit{childrenSum}$。

```py [sol1-Python3]
class Solution:
    def replaceValueInTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        root.val = 0
        q = [root]
        while q:
            tmp = q
            q = []
            next_level_sum = 0  # 下一层的节点值之和
            for node in tmp:
                if node.left:
                    q.append(node.left)
                    next_level_sum += node.left.val
                if node.right:
                    q.append(node.right)
                    next_level_sum += node.right.val

            # 再次遍历，更新下一层的节点值
            for node in tmp:
                children_sum = (node.left.val if node.left else 0) + \
                               (node.right.val if node.right else 0)
                if node.left: node.left.val = next_level_sum - children_sum
                if node.right: node.right.val = next_level_sum - children_sum
        return root
```

```java [sol1-Java]
class Solution {
    public TreeNode replaceValueInTree(TreeNode root) {
        root.val = 0;
        var q = new ArrayList<TreeNode>();
        q.add(root);
        while (!q.isEmpty()) {
            var tmp = q;
            q = new ArrayList<>();
            int nextLevelSum = 0; // 下一层的节点值之和
            for (var node : tmp) {
                if (node.left != null) {
                    q.add(node.left);
                    nextLevelSum += node.left.val;
                }
                if (node.right != null) {
                    q.add(node.right);
                    nextLevelSum += node.right.val;
                }
            }

            // 再次遍历，更新下一层的节点值
            for (var node : tmp) {
                int childrenSum = (node.left != null ? node.left.val : 0) +
                                  (node.right != null ? node.right.val : 0);
                if (node.left != null) node.left.val = nextLevelSum - childrenSum;
                if (node.right != null) node.right.val = nextLevelSum - childrenSum;
            }
        }
        return root;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    TreeNode *replaceValueInTree(TreeNode *root) {
        root->val = 0;
        vector<TreeNode*> q = {root};
        while (!q.empty()) {
            vector<TreeNode*> nxt;
            int next_level_sum = 0; // 下一层的节点值之和
            for (auto node: q) {
                if (node->left) {
                    nxt.push_back(node->left);
                    next_level_sum += node->left->val;
                }
                if (node->right) {
                    nxt.push_back(node->right);
                    next_level_sum += node->right->val;
                }
            }

            // 再次遍历，更新下一层的节点值
            for (auto node: q) {
                int children_sum = (node->left ? node->left->val : 0) +
                                   (node->right ? node->right->val : 0);
                if (node->left) node->left->val = next_level_sum - children_sum;
                if (node->right) node->right->val = next_level_sum - children_sum;
            }
            q = move(nxt);
        }
        return root;
    }
};
```

```go [sol1-Go]
func replaceValueInTree(root *TreeNode) *TreeNode {
	root.Val = 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		nextLevelSum := 0 // 下一层的节点值之和
		for _, node := range tmp {
			if node.Left != nil {
				q = append(q, node.Left)
				nextLevelSum += node.Left.Val
			}
			if node.Right != nil {
				q = append(q, node.Right)
				nextLevelSum += node.Right.Val
			}
		}

		// 再次遍历，更新下一层的节点值
		for _, node := range tmp {
			childrenSum := 0 // node 左右儿子的节点值之和
			if node.Left != nil {
				childrenSum += node.Left.Val
			}
			if node.Right != nil {
				childrenSum += node.Right.Val
			}
			// 更新 node 左右儿子的节点值
			if node.Left != nil {
				node.Left.Val = nextLevelSum - childrenSum
			}
			if node.Right != nil {
				node.Right.Val = nextLevelSum - childrenSum
			}
		}
	}
	return root
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为二叉树的节点个数。
- 空间复杂度：$O(n)$。
