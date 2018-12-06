# Day 6 - The saga of the brute force continues

I wrote a short python snippet to prototype and decided to rewrite it in Go when I had more time.

While writing it in go I found some other optimizations and the code in the end turned out pretty nice.

What made the go code that long in the first place is the `eval` function, go is intended as a compiled language and it works against you if you try to make it interpreted.

So if you want to interpret something you have to write your own interpreter. This was pretty simple, but frankly also kind of tedious, but now that I have it it will probably be of use in the future.

The function is easily extendible and if I add some validation and error handling this could turn out to be a really useful and safe function.

And of course it's always great to be rewarded in the end for what might seems like wasted effort to most :)

```console
kb@day05: time go run main.go && time python test.py
5857

real    0m4.915s
user    0m5.101s
sys     0m0.091s
5857

real    0m44.099s
user    0m43.550s
sys     0m0.042s
```

