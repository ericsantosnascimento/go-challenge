travel audience Go challenge
============================

Solution
----

Basically my idea was make a simple as possible solution to call all the endpoints in the url and extract the url from it
the task sounded simple, at least if didn't miss something really important.

So first thing I tried to make sure that I would fail fast if the url wasn't present, so in case of missing u param i will
return empty array of numbers

In case attribute is present I will try to call the which api in a for loop, this method is real candidate to parallel execution
but with my given time frame I wont be able to so.

After call api I extract parse the json to a type and I extract the number of it, the array of numbers is fully added
to my result list.

After getting out of the loop I remove duplicated and sort this smaller array right after.

As I final step I check how long did it take and in case was longer than 500ms i return empty array, otherwise the numbers

I wasn't quite sure the idea of the 500ms check, I understood that would be some sort of circuit breaker emulation?

Concerns
-----

I didn't know how to structure this files on Go, on Java I create have classes/files doing specific responsibilities
but wasn't sure if Go works like that.

I didn't quite get the difference between methods and functions.

As I mentioned I wish i could have make parallel calls to call apis.

I started to read about unit testing on go but I didn't have time left for it.



