# Intended solution

There might be a way to solve this challenge using reverse engineering, but I did not try that.

There are two key insights to solve this challenge:

* The flag will start with "flag{"
* The program compares character by character and outputs a dot if the comparison is correct; if the comparison is wrong, the program exists