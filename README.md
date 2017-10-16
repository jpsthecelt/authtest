# Authtest
I originally wrote this commandline program in Go after having some difficulty producing 
my swift program to do the same (in a timely manner).  

The Need
--------
I needed a way to query/update a Casper-instance from the commandline (suitable for 
automation).

After being enamored with, and playing with Swift, I needed a working demo but was having
trouble getting it done in time to show a proof-of-concept.

Thinking about The Solution
---------------------------
I figured I'd 'tackle' the problem by trying to rewrite it in go, which I'd used before
and found pretty accessable. I knew I'd need to read some 'private' details (like username,
password, and server-url) from a file (preferably JSON), so reading/parsing needed to be 
easy.  After perusing some solutions on Stack-Overflow, I found it extremely easy to 
quickly produce a prototype.

Go:
---
I liked the fact that I could produce free-standing binaries for a variety of
platforms, and although I'd previously maligned the language/environmenet for the lack of 
an IDE/source-line debugger, I'd recently found the GOGLAND (jetbrains) implementation, which,
though in beta, worked very well.
voila!

My first implementation was 70 lines (including whitespace),

Modification History
---------------------------
Initially happy with my results, I'll continue modifying this and get to a more usable 
'product'.
jpsthecelt-101517

