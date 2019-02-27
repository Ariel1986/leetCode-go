<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/trapping-rain-water/

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.

The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

Example:

Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
```

- [Brute Force](#brutrforce)
- [Dynamic Programming](#dynamicprogramming)
- [Using Stack](#usingstack)
- [Using 2 Pointers](#using2pointers)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Brute Force
```
Intuition

Do as directed in question. For each element in the array, we find the maximum level of water it can trap after the rain, which is equal to the minimum of maximum height of bars on both the sides minus its own height.

Algorithm

Initialize ans=0
Iterate the array from left to right:
	Initialize max_left=height[i] and max_right=height[i]
	Iterate from the current element to the beginning of array updating:
	max_left=max(max_left,height[j])
	Iterate from the current element to the end of array updating:
	max_right=max(max_right,height[j])
	Add min(max_left,max_right) - height[i] to ans
```
```cpp
class Solution {
public:
    int trap(std::vector<int>& height) {
        int ans = 0, size = height.size();
        for (int i = 1; i < size-1; i++) {
            int maxLeft = height[i], maxRight = height[i];
            for (int j = i-1; j >=0; j--){ //search the left part for max bar size
                maxLeft = std::max(maxLeft, height[j]);
            }
            for (int j = i+1; j < size; j++){ //search the right part for max bar size
                maxRight = std::max(maxRight, height[j]);
            }
            ans += std::min(maxLeft, maxRight)-height[i];
        }
        
        return ans;
    }
};
```

# Dynamic Programming
```
Intuition

In brute force, we iterate over the left and right parts again and again just to find the highest bar size upto that index. But, this could be stored. Voila, dynamic programming.

Algorithm

Find maximum height of bar from the left end upto an index i in the array left_max.
Find maximum height of bar from the right end upto an index i in the array right_max.
Iterate over the height array and update ans:
Add min(max_left[i],max_right[i]) - height[i] to ans
```

```cpp
class Solution {
public:
    int trap(std::vector<int>& height) {
        if(height.size() == 0){
            return 0;
        }
        int ans = 0, size = height.size();
        std::vector<int> left(size), right(size);
        left[0]=height[0];
        for (int i = 1; i < size; i++){
            left[i]= std::max(left[i-1], height[i]);
        }
        
        right[size-1]=height[size-1];
        for(int i= size-2; i>=0; i--){
            right[i] = std::max(right[i+1], height[i]);
        }
        
        for(int i = 1; i < size-1; i++){
            ans += std::min(left[i], right[i]) - height[i];
        }
        return ans;
    }
};
```

# Using Stack
```
Intuition

Instead of storing the largest bar upto an index as in Approach 2, we can use stack to keep track of the bars that are bounded by longer bars and hence, may store water. Using the stack, we can do the calculations in only one iteration.

We keep a stack and iterate over the array. We add the index of the bar to the stack if bar is smaller than or equal to the bar at top of stack, which means that the current bar is bounded by the previous bar in the stack. If we found a bar longer than that at the top, we are sure that the bar at the top of the stack is bounded by the current bar and a previous bar in the stack, hence, we can pop it and add resulting trapped water to \text{ans}ans.

Algorithm

Use stack to store the indices of the bars.
Iterate the array:
While stack is not empty and height[current]>height[st.top()]height[current]>height[st.top()]
	It means that the stack element can be popped. Pop the top element as top.
	Find the distance between the current element and the element at top of stack, which is to be filled. distance = current - st.top() - 1
	Find the bounded height bounded_height = min(height[current], height[st.top()]) - height[top]
	Add resulting trapped water to answer ans += distance * bounded_height
Push current index to top of the stack
Move \text{current}current to the next position

```
```cpp
class Solution {
public:
    int trap(std::vector<int>& height) {
        if(height.size() == 0){
            return 0;
        }
        int ans = 0,idx = 0;
        std::stack<int> s;
        while(idx < height.size()){
            while(!s.empty() && height[idx] > height[s.top()]){
                int topIdx = s.top();
                s.pop();
                if(s.empty()){
                    break;
                }
                int distance = idx -s.top()-1;
                int bounded_height = std::min(height[idx], height[s.top()]) - height[topIdx];
                ans += distance * bounded_height;
            }
            s.emplace(idx++);
        }
        
        return ans;
    }
};
```

# Using 2 Pointers
```
Intuition As in Approach 2, instead of computing the left and right parts seperately, we may think of some way to do it in one iteration. From the figure in dynamic programming approach, notice that as long as \text{right_max}[i]>\text{left_max}[i] (from element 0 to 6), the water trapped depends upon the left_max, and similar is the case when \text{left_max}[i]>\text{right_max}[i] (from element 8 to 11). So, we can say that if there is a larger bar at one end (say right), we are assured that the water trapped would be dependant on height of bar in current direction (from left to right). As soon as we find the bar at other end (right) is smaller, we start iterating in opposite direction (from right to left). We must maintain \text{left_max} and \text{right_max} during the iteration, but now we can do it in one iteration using 2 pointers, switching between the two.

Algorithm

Initialize left pointer to 0 and right pointer to size-1
While left<right, do:
	If height[left] is smaller than height[right]
		If height[left]>=left_max, update left_max
		Else add left_max-height[left] to ans
		Add 1 to left.
	Else
		If height[right]>=right_max, update right_max
		Else add right_max-height[right] ans
		Subtract 1 from right.
```
```cpp
class Solution {
public:
    int trap(std::vector<int>& height) {
        if(height.size() == 0){
            return 0;
        }
        int ans = 0;
        int left =0, right = height.size()-1;
        int l_max = 0, r_max=0;
        while(left < right){
            if(height[left] < height[right]){
                height[left] > l_max? (l_max = height[left]):ans +=(l_max - height[left]);
                ++left;
            }else{
                height[right] >r_max?(r_max = height[right]) : ans +=(r_max-height[right]);
                --right;
            }
        }
 
        return ans;
    }
};

```