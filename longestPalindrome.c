#include <stdio.h>
#include <string.h>

// Manacher算法解决最长回文子串

#define max(a,b) ((a) > (b)? (a):(b))
#define min(a,b) ((a) < (b)? (a):(b))

int p[100];

void init(const char *str, char *tmp)
{
    tmp[0] = '@';
    int len = strlen(str);
    int i = 0;
    for (i = 1; i < 2 *len; i += 2)
    {
        tmp[i] = '#';
        tmp[i+1] = str[i/2];
    }
    tmp[2*len + 1] = '#';
    tmp[2*len + 2] = '@';
    tmp[2*len + 3] = 0;
}

//Manacher算法计算过程
int longestPalindrome(const char *str)
{
    int mx = 0, id = 0;
    int len = strlen(str);
    int longest = 0;

    int i = 0;
    for (i = 1; i <= len; i++)
    {
        p[i] = mx > i ? min(p[2*id - i], mx - i):1;

        while (str[i + p[i]] == str[i - p[i]])
            p[i]++;

        if (p[i] + i > mx) {
            mx = p[i] + i;
            id = i;
        }

        longest = max(longest, p[id]);
    }

    return longest - 1;
}

int main() {
    char S[] = {'1', '2', '2', '1', '2' ,'3', '2', '1'};
    char temp[100];
    init(S, temp);
    
    int result = longestPalindrome(temp);
    printf("result: %d\n", result);

    return 0;
}
