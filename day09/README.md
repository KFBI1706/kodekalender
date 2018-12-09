# Day 9 - Artificial limits

So first I thought this task was going to be way harder. I was pulling up multi anagram code and thinking about ways to solve it that way.

Turns out this was much easier, regardless I still imposed another artificial limitation on myself, I didn't use the fact that the first letter hash is md5sum("julekalender") + char

This makes the code highly specific and not really reusable, so I'll just say that I didn't use it so my code is more portable. That said no one would ever give this much information given the way the hash chain is constructed, but I digress.

The chain is really easy to unlock either way. Loop through all the chain-hash pairs, combining the char and hash and seeing what hash it matches. When there's a match we store the relationship between the matching hash-char pair and the ones used to construct the hash.

After that I just iterated over the chain again, now with the relationships populated. Which means that any relationship that's not populated must be the start/end.
