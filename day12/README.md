# Day 12 - Emojis here emojis there emojis everywhere!

So this task was interresting, I fumbled with a lot of different solution until I found the correct one. Some of the things I tried includes:
* base100 https://github.com/AdamNiederer/base100 which encodes data to and from emoji
* codemoji - crypto example by Mozilla, which encrypts and decrypts with emojis
* emojicode - programming with emojis

Then I went on to more time consuming solutions, I initially thought this was some sort of substitution cipher, as there were only 22 unique chars and I felt I could see a patterns lots of places in the ciphertext. Also some of the monograms / letter frequencies matched up pretty well.

All this turned out to be a read herring as I had way to little ciphertext to go off of. Then while sitting on the buss, with the most spotty internet and nothing else to do I started working on the task again. This time trying some sort of shifting of the values, mainly because of the name emoji-**rot**. I initially dismissed this idea, because I was set on the substitution cipher path, but now looking back at it this is way smarter.

The thing that got me working on that solution was the fact that there wasn't much spread in the emojis. I noticed that there were only faces. And in codemoji boats and other categories of the wide emoji spectrum would often appear. This made me check the difference between the highest and the lowest values emoji and it turned out to be only 55, BINGO! Regardless of the shift the values are all close together and they all fit within the ASCII values with good margin.

So I wrote a for-loop looping all values between 0 and 127 and shifting by that value, I initially just output it to file, and while scrolling through the lines hidden in lots of garbage output the answer was hidden on line 32. So the shift was 32, awesome!
