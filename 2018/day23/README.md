# Day 23 - Too personal

This task had us make verification code for the Norwegian identification number. This was not only fun but also really useful. It also illustrates the blatant flaw with the system, namely that it's laughably easy to brute force a number and even more simple to just generate a valid one. Some validators are really bad as well and allow higher bounds for the day and month.

The only bad thing is as always the dataset, lots' of edge-cases where not included, and led to me submitting code which checked for males instead of females and it still worked. That's not fun. I want to have the definite right solution and not wonder if I've done anything wrong when my answer is accepted. This maybe intentional but it would've been fun if they made the challenges harder and the datasets individual or at least much larger.

In the end I fixed all the bugs and ended up writing a pretty short and sweet solution for the task. 0.004 runtime.
