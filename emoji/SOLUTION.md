# Intended solution

To solve this challenge, the password has to have 42 different emojis. Here are the steps to get to that value:

* We know that entropy of a string is calculated as $$log_2(b^n)$$ 
where $b$ is the number of elements in the dictionary (the number of possible values that a character can have) and $n$ is the length of the password
* For this challenge, we want to know the value of $n$ given that the entropy needs to be $497.246559$
* $b$ is $3664$ and it is obtained from here: https://unicode.org/emoji/charts-15.0/emoji-counts.html
* $n$ is obtained as follows:
    $$n=log(2^{497.246559})/log(3664)=41.999999$$