As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.

Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:
z -= (z*z - x) / (2*z)

Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.
