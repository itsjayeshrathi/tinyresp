# RESP DOCS

## two exceptions in this model
  - redis supports pipelining -> client can send bulk of commands 
    at once, and wait for replies later. 
  - if client sub to pub/sub channel then it becomes resp to push protocol
    so client don't need to send req as it will get the data as it subs to channel 
  
## RESP description
protocol suppports following data types: Simple Strings, Errors, Integers, Bulk Strings and Arrays
 - clients send commands to redis server as RESP Arary of Bulk Strings
 - server replies with one of RESP Type

In RESP, the type of some data depends of first byte: 
  - Simple Strings the first byte is `+`
  - Errors the first byte is `-`
  - Integers the first byte is `:`
  - Bulk Strings the first by is `$`
  - Arrays the first byte is `*` 

In RESP different parts of protocol are terminated with "\r\n"
