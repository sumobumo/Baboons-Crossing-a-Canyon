# Baboons-Crossing-a-Canyon
Cloud and distributed computing 

This is an assignment I turned in for the Cloud and distributed computing class I took. This assignment shows proper handling of deadlocks and go routines. 

## Problem statement:

A student majoring in anthropology and minoring in computer science has
embarked on a research project to see if African baboons can be taught about
deadlocks. She locates a deep canyon and fastens a rope across it, so the
baboons can cross hand-over-hand.

Passage along the rope follows these rules:

  1. Several baboons can cross at the same time, provided that they are all
going in the same direction.

  2. If eastward moving and westward moving baboons ever get onto the rope
at the same time, a deadlock will result (the baboons will get stuck in the
middle) because it is impossible for one baboon to climb over another one
while suspended over the canyon.

  3. If a baboon wants to cross the canyon, he must check to see that no other
baboon is currently crossing in the opposite direction.

  4. Your solution should avoid starvation. When a baboon that wants to cross
to the east arrives at the rope and finds baboons crossing to the west, the
baboon waits until the rope in empty, but no more westward moving
baboons are allowed to start until at least one baboon has crossed the other
way.


Write a go program to simulate activity for this canyon crossing problem:

  a. Simulate each baboon as a separate goroutine. 

  b. Altogether, 30 baboons will cross the canyon, with a random number generator
specifying whether they are eastward moving or westward moving (with equal probability).

  c. Use a random number generator, so the time between baboon arrivals is between 1 and 8 seconds.

  d. Each baboon takes 1 second to get on the rope. (That is, the minimum inter-baboon spacing is 1 second.)

  e. All baboons travel at the same speed. Each traversal takes exactly 4 seconds, after the baboon. 

  f. Use shared memory for synchronization. Additional communication via channels is allowed,
but do not use channels unless such communication is clearly needed.

  g. Each baboon should report when it arrives at the canyon and which direction it
is trying to go. It should report when it gets on the rope and the id’s of the other
baboons already on the rope, and when it gets off the rope. Give each baboon a unique
identifier and use that in the reporting.
