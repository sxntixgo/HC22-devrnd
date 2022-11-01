# Intended solution

There might be a way to solve this challenge using reverse engineering, but I did not try that.

Here's a python script that brute forces the next character based on the time delta:

```python
from subprocess import Popen, PIPE, STDOUT
import datetime


stored_delta = 0
stored_character = '0'
flag = 'flag{'
dictionary = "{}abcdefghijklmnopqrstuvwxyz"

for c in dictionary:
    start = datetime.datetime.today()
    p = Popen(['./nodots'], stdout=PIPE, stdin=PIPE, stderr=PIPE)
    cmdinput = flag + c
    stdout_data = p.communicate(input=cmdinput.encode())[0]
    end = datetime.datetime.today()
    delta = end-start
    if delta.total_seconds() > stored_delta:
        stored_delta = delta.total_seconds()
        stored_character = c

print(flag + stored_character, stored_delta)
```