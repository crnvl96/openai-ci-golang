using System;

namespace PalindromeChecker {
    class Program {
        static void Main(string[] args) {
            Console.WriteLine(IsPalindrome("racecar"));
            Console.WriteLine(IsPalindrome("hello"));
        }
        
        static bool IsPalindrome(string word) {
            int left = 0;
            int right = word.Length - 1;
            
            while (left < right) {
                if (word[left] != word[right]) {
                    return false;
                }
                left++;
                right--;
            }
            
            return true;
        }
    }
}
