#include <iostream>
#include <algorithm>
#include <unordered_map>

using namespace std;

class Solution {
    public :
        int lengthOfLongestSubstring(string s) {
            int res = 0, left = -1, n = s.size();

            unordered_map<int, int> m;
            for (int i = 0; i < n; i++) {
                if (m.count(s[i]) && m[s[i]] > left) {
                    left = m[s[i]];
                }

                m[s[i]] = i;
                res = max(res, i - left);
            }

            return res;
        }
};

int main()
{
    string t = "abcdefghacegf";
    
    Solution s;
    int rc = s.lengthOfLongestSubstring(t);
    if (rc) {
        printf("the max length is: %d\n", rc);
    }

    return 0;
}
