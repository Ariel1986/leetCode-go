<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/merge-k-sorted-lists/

Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
```

- [Bruce Force](#bruceforce)
```
Intuition & Algorithm

1. Traverse all the linked lists and collect the values of the nodes into an array.
2. Sort and iterate over this array to get the proper value of nodes.
3. Create a new sorted linked list and extend it with the new nodes.
*/
```

- [Compare one by one](#compareonebyone)
```
1. Compare every \text{k}k nodes (head of every linked list) and get the node with the smallest value.
2. Extend the final sorted linked list with the selected nodes.
```

- [Optimize Approach 2 by Priority Queue](#bypriorityqueue)
```
Almost the same as the one above but optimize the comparison process by 
```

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Optimize Approach 2 by Priority Queue
```cpp
/* priority queue*/
ListNode* mergeKLists(std::vector<ListNode*>& lists) {
    if (lists.size() == 0) return nullptr;
    
    auto cmp = [](ListNode* l, ListNode* r){
        if (l == nullptr || r == nullptr){
            return false;
        }
        return l->val > r->val;
    };
    
    std::priority_queue<ListNode*, std::vector<ListNode*>, decltype(cmp)> pq(cmp);

    //另一种声明pq的方式
    //std::priority_queue<ListNode*, std::vector<ListNode*>, cmp>pq;
    for (auto list : lists){
        if (list != nullptr){ //Note:注意判断链表是否为空
            pq.emplace(list);
        }
    }
    
    ListNode dummyHead(0);
    ListNode* tail = &dummyHead;
    
    while(!pq.empty()){
        tail->next = pq.top();
        tail = tail->next;
        pq.pop();
        
        if(tail->next != nullptr){
            pq.emplace(tail->next);
        }
    }
    
    return dummyHead.next;
}
```

```cpp
/* by heap */
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
struct cmp {
    bool operator() (const ListNode* l, const ListNode* r){
        if (l == NULL or r == NULL) return 0;
        return l->val > r->val;
    }
};


class Solution {
public:
    ListNode* mergeKLists(std::vector<ListNode*>& lists) {
        if(lists.size() == 0){return nullptr;}
        
        std::vector<ListNode*> v;
        for(auto node : lists){
            if(node){
                v.push_back(node);
            }
        }
        std::make_heap(v.begin(), v.end(), cmp());
      
        ListNode dummyHead(0);
        ListNode *tail = &dummyHead;
        while(!v.empty()){
            //pop element
            std::pop_heap(v.begin(), v.end(),cmp());
            tail->next = v.back();
            tail = tail->next;
            v.pop_back();

            //push element
            if(tail->next != nullptr){
               v.emplace_back(tail->next);
                std::push_heap(v.begin(), v.end(),cmp());
            }
        }
        
        return dummyHead.next;
    }
};
```





func mergeKLists(lists []*ListNode) *ListNode {
	dummyHead := ListNode{}

	return dummyHead.Next
}

/*cpp

 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };

 struct cmp {
    bool operator() (const ListNode* l, const ListNode* r){
        if (l == NULL or r == NULL) return 0;
        return l->val > r->val;
    }
};


class Solution {
public:
    ListNode* mergeKLists(std::vector<ListNode*>& lists) {
        if(lists.size() == 0){return nullptr;}

        std::vector<ListNode*> v;
        for(auto node : lists){
            if(node){
                v.push_back(node);
            }
        }
        std::make_heap(v.begin(), v.end(), cmp());

        ListNode dummyHead(0);
        ListNode *tail = &dummyHead;
        while(!v.empty()){
            //pop element
            std::pop_heap(v.begin(), v.end(),cmp());
            tail->next = v.back();
            tail = tail->next;
            v.pop_back();

            //push element
            if(tail->next != nullptr){
               v.emplace_back(tail->next);
                std::push_heap(v.begin(), v.end(),cmp());
            }
        }

        return dummyHead.next;
    }
};
*/
