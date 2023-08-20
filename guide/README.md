## The LeetCode Beginner's Guide
### 1480. Running Sum of 1D Array
<b>Constraints:</b><br>

    1 <= nums.length <= 1000 <br>
    -10^6 <= nums[i] <= 10^6

<b>Description:</b><br>
Given an array nums. We define a running sum of an array as ```runningSum[i] = sum(nums[0]…nums[i])```.
Return the running sum of nums. 
### 1672. Richest Customer Wealth
<b>Constraints:</b><br>
``` m == accounts.length
    n == accounts[i].length
    1 <= m, n <= 50
    1 <= accounts[i][j] <= 100
```
<b>Description</b><br>
You are given an m x n integer grid accounts where ```accounts[i][j]``` is the amount of money the ```i^th```customer has in the ```j^th``` bank. Return the wealth that the richest customer has.
A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.
### 412. Fizz Buzz 
<b>Constraints:</b>
```1 <= n <= 104``` <br>
<b>Description</b><br> 
Given an integer n, return a string array answer (1-indexed) where:
```
answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true. 
```
### 1342. Number of Steps to Reduce a Number to Zero
<b>Constraints:</b> 0 <= num <= 106 <br>
<b>Description</b><br>
Given an integer num, return the number of steps to reduce it to zero.
In one step, if the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.

### 876. Middle of the Linked List
<b>Constraints:</b><br>
The number of nodes in the list is in the range [1, 100].<br>
```1 <= Node.val <= 100``` <br>

<b>Description</b><br>
Given the head of a singly linked list, return the middle node of the linked list.<br>
If there are two middle nodes, return the second middle node.<br>
<b>Example:</b><br>
```
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.
```

### 383. Ransom Note
<b>Constraints:</b><br>
```
1 <= ransomNote.length, magazine.length <= 105
ransomNote and magazine consist of lowercase English letters.
```
<b>Description</b><br>
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
Each letter in magazine can only be used once in ransomNote.
