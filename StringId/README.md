# String ID Problem

## Requirement
Please create the function that can input any character (up to 255 chars) and return unique 10 digits in java or golang.
Same character, same return, and case-insensitive. Any printable character will support it.

For example, if we input
```text
abc --> 1234567890
Abc --> 123456789
ABC --> 1234567890
กขค --> 9874561230
ก_8 --> 3214560789
😀 --> 5555555555
```

### Assumptions
1. digital ID for a string value must be <u>truly unique</u> as required and not *probably unique*, eg. a hashcode 
1. no requirements for performance or concurrency - will design for but not test
1. no warm-up or latency limitations
1. unique values will be <u>only per runtime environment</u>, and not across runtimes, machines, clusters etc.
1. 1 emoji counts as 2 characters
1. any character up to 255 characters means up to and including 255 characters
1. KISS/YAGNI no frills 

## Solution
Maven + Java (11) Project.

No pure function such as a hash function can reduce any string up to length 255 chars to a 10 digit integer without an
approximation and therefore will eventually produce collisions. To ensure uniqueness we must therefore utilise a 
sequence persisted with an index.