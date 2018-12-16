# Day 16 - Obscurity in the name of fewer lines

My sleep schedule is pretty messed up at this point, but I digress. This task is pretty easy and because of that I think I've found a pattern in what I do with these kinds of tasks. If the task is easy to code, it's often combined with a long runtime. So with these kinds of tasks I like to cripple the readability and make a complicated and blazingly fast solution.

This is a ***prime*** example of that kind of task :)

So the solution is pretty straightforward, read in the text as a string then iterate over each second byte as those are the bytes with numbers. Then there's the ~complicated~ part, This is just a loop with two iteration, where j is set to 0 and 1, and then I've used those values to generalize the palindrome checking. Because as I wrote the checks, I realized that the two differences were the starting sum, which is doubled at the start when it's an even palindrome and the start index for palindrome checking, as in an odd-palindrome we start from the same char.

So with some arithmetic I got the inner palindrome checking working for both even and odd palindrome, thus saving about 12 lines of repeated code :). I continually add onto the sum as I check if the palindrome is bigger and if it's a prime, overwriting the max variable when both are true.

All in all I'm pretty happy with the solution, it runs in about 2s. Looking past the really long for loop and the shoddy arithmetic it's pretty concise and more important, fast code!

If you have any improvements, ping me!

