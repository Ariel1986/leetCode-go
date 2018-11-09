package leetCode

/*
Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.

Example 1:

Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".
Example 2:

Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".
Example 3:

Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".
Example 4:

Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".
Note:

1 <= S.length <= 200
1 <= T.length <= 200
S and T only contain lowercase letters and '#' characters.
Follow up:

Can you solve it in O(N) time and O(1) space?
*/

// M1: use stack
class Solution {
	public:
		bool backspaceCompare(std::string S, std::string T) {
			std::stack<char> s, t;
			auto pushStack = [](std::stack<char> *sta, std::string *str){
				for(auto&& c : *str){
					if (c == '#' ){
						if(!sta->empty()){
							sta->pop();
						}
					}else{
						sta->emplace(c);
					}
				}
			};
			
			auto stackToString = [](std::stack<char> *sta)->std::string{
				std::string str;
				while(!sta->empty()){
					str += sta->top();
					sta->pop();
				}
				
				return str;
			};
			
			pushStack(&s, &S);
			pushStack(&t, &T);
			
			if(s.size() != t.size()){
				return false;
			}
			
			std::string sResp = stackToString(&s);
			std::string tResp = stackToString(&t);
			
			return sResp.compare(tResp) == 0;
		}
	};

	// M2: use string
	class Solution {
		public:
			bool backspaceCompare(std::string S, std::string T) {
				std::string s, t;
				auto convStr = [](std::string *str)->std::string {
					std::string result;
					for(auto&& c : *str){
						if (c == '#' ){
							result.push_back(c);
						}else{
							result.pop_back();
						}
					}
					
					return result;
				};
			  
				std::string sResp = convStr(&S);
				std::string tResp = convStr(&T);
				
				return sResp.compare(tResp) == 0;
			}
		};