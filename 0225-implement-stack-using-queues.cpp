package leetCode

/*
https://leetcode.com/problems/implement-stack-using-queues/

Implement the following operations of a stack using queues.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
empty() -- Return whether the stack is empty.
Example:

MyStack stack = new MyStack();

stack.push(1);
stack.push(2);
stack.top();   // returns 2
stack.pop();   // returns 2
stack.empty(); // returns false
Notes:

You must use only standard operations of a queue -- which means only push to back, peek/pop from front, size, and is empty operations are valid.
Depending on your language, queue may not be supported natively. You may simulate a queue by using a list or deque (double-ended queue), as long as you use only standard operations of a queue.
You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).

*/

//M1: use two queues
class MyStack {
	private:
		std::queue<int> q1, q2;
	public:
		/** Initialize your data structure here. */
		MyStack() {
		}
		
		/** Push element x onto stack. */
		void push(int x) {
			if(!q1.empty()){
				q1.emplace(x);
			}else{
				q2.emplace(x);
			}
		}
		
		/** Removes the element on top of the stack and returns that element. */
		int pop() {
			int result = 0;
			if(!q1.empty()){
				while(q1.size() > 1){
					q2.emplace(q1.front());
					q1.pop();
				}
				result = q1.front();
				q1.pop();
			}else{
				while(q2.size() > 1){
					q1.emplace(q2.front());
					q2.pop();
				}
				result = q2.front();
				q2.pop();
			}
			return result;
		}
		
		/** Get the top element. */
		int top() {
			if(q1.empty() && q2.empty()){
				return 0;
			}
			
			int result = 0;
			if(!q1.empty()){
				while(q1.size() > 1){
					q2.push(q1.front());
					q1.pop();
				}
				result = q1.front();
				q2.push(q1.front());
				q1.pop();
			}else{
				while(q2.size() > 1){
					q1.push(q2.front());
					q2.pop();
				}
				result = q2.front();
				q1.push(q2.front());
				q2.pop();
			}
			
			return result;
		}
		
		/** Returns whether the stack is empty. */
		bool empty() {
			return q1.empty() && q2.empty();
		}
	};

// M2: use one queue
class MyStack {
private:
    std::queue<int> q;
    int size;
    
public:
    /** Initialize your data structure here. */
    MyStack() {
        size = 0;
    }
    
    /** Push element x onto stack. */
    void push(int x) {
        q.emplace(x);
        if (size > 0){
            for (int i = 0; i < size; i++){
                q.emplace(q.front());
                q.pop();
            }
        }
        size++;
    }
    
    /** Removes the element on top of the stack and returns that element. */
    int pop() {
        int result = q.front();
        q.pop();
        size--;
        return result;
    }
    
    /** Get the top element. */
    int top() {
        return q.front();
    }
    
    /** Returns whether the stack is empty. */
    bool empty() {
        return q.empty();
    }
};