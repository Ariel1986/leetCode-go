package leetCode

/*
https://leetcode.com/problems/implement-queue-using-stacks/
Implement the following operations of a queue using stacks.

push(x) -- Push element x to the back of queue.
pop() -- Removes the element from in front of queue.
peek() -- Get the front element.
empty() -- Return whether the queue is empty.
Example:

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek();  // returns 1
queue.pop();   // returns 1
queue.empty(); // returns false
Notes:

You must use only standard operations of a stack -- which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.
You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).
*/

// M1: use 2 stacks
class MyQueue {
	private:
		std::stack<int> pushS, popS;
		
		void topHelper(std::stack<int> *pushS, std::stack<int> *popS){
			if(!popS->empty()){
				return;
			}
			while(!pushS->empty()){
				popS->emplace(pushS->top());
				pushS->pop();
			}
		}
		
	public:
		/** Initialize your data structure here. */
		MyQueue() {
			
		}
		
		/** Push element x to the back of queue. */
		void push(int x) {
			pushS.push(x);
		}
		
		/** Removes the element from in front of queue and returns that element. */
		int pop() {
			topHelper(&pushS, &popS);
			int result = popS.top();
			popS.pop();
			
			return result;
		}
		
		/** Get the front element. */
		int peek() {
			topHelper(&pushS, &popS);
			return popS.top();
		}
		
		/** Returns whether the queue is empty. */
		bool empty() {
			return pushS.empty() && popS.empty();
		}
	};
	
	//M2
	/*
 ou can implement queue using just one stack by either making push process costlier or pop costlier. Since we have two functions (top() and pop()) which require the top element of the stack, well make the push() function costlier - that is, for pushing a new element, we recursively pop the stack till it is empty and push it at the bottom of the stack, and take advantage of the recursive call to push back in the popped elements once the recursive call hits the base condition and returns. This implementation makes pop() and peek() functions easier. pop() is just going to pop off the top element in stack and peek() will return the top most element.
*/
class MyQueue {
	private:
		std::stack<int> s;
		void pushHelper(int x){
			if(s.empty()){
				s.emplace(x);
				return ;
			}
			int data = s.top();
			s.pop();
			pushHelper(x);
			s.emplace(data);
			return ;
		}
		
	public:
		/** Initialize your data structure here. */
		MyQueue() {
			
		}
		
		/** Push element x to the back of queue. */
		void push(int x) {
			pushHelper(x);
		}
		
		/** Removes the element from in front of queue and returns that element. */
		int pop() {
			int result = s.top();
			s.pop();
			
			return result;
		}
		
		/** Get the front element. */
		int peek() {
			return s.top();
		}
		
		/** Returns whether the queue is empty. */
		bool empty() {
			return s.empty();
		}
	};
	
	/**
	 * Your MyQueue object will be instantiated and called as such:
	 * MyQueue obj = new MyQueue();
	 * obj.push(x);
	 * int param_2 = obj.pop();
	 * int param_3 = obj.peek();
	 * bool param_4 = obj.empty();
	 */