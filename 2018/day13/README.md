# Day 13 - Sleep on it

So today I woke up late, I saw the task and I had no idea how to do it, the instruction were not computing in my head and I was really unsure what to do. So I just stopped right there, without beating myself too much about it I just continued with the other work I had in front of me and let the task sit.

Then at around 13:00 I had a peek at the task again, and the description had been updated, and now it's much more understandable. First I wrote a function to check if there were multiple sums for a number, then I just wrote a for-loop going upwards, adding the numbers with a distinct sum to an array and summing the primes in that list.

This took around 14s, but it worked so, hey. Then I started writing a more optimized solution generating all the primes beforehand with a sieve. Only checking if there's 100 primes when a new prime is added etc etc. And now my solution is running at 250 ms.

Also as a side note, look at those closing brackets, not sure if I like the look of it. And if there's a better way to do it, ping me!
