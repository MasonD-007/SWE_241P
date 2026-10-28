Task Description 

 An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once. For example, “mango” and “ogman” are anagrams. 

Task-1: Given an array of strings Strings, group the words that are anagrams to each other. You can return the answer in any order. 

For example, 

Input: strings = ["bucket","rat","mango","tango","ogtan","tar"]
Output: [["bucket"],["rat","tar"],["mango"],["tango","ogtan"]]

For this assignment, complete the following steps -

Step-1: Implement a function List<List<String>> groupAnagram(List<String> strings) that takes a list of strings and returns a list of lists of strings

Step-2: To group the anagrams, you would need to sort all the given strings by the characters, i.e., the ASCII code of the characters. After sorting the strings that converge to the same string, they are grouped together in the same anagram bucket. For example, after sorting, both mango and ogman become agmno; therefore, they are in the same bucket. Implement a function String sortString(String str) that sorts a string. You can’t use sorting API for this. Please use two of the following sorting algorithms and compare their performance. Write sample test cases to validate your implementation.

Mergesort
Quicksort
Heapsort
Radix sort
