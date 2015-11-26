# cmpe273-lab3
1) To run this project, first compile and run files from Part1 folder.
This will activate PUT and GET operations on following ports:
3000
3001
3002

2) Now compile and run files from ClientFolder.
It has Client that does the consistent hashing.
Client gets activated on port 8080.

a) PUT operation:

PUT /keys/{key_id}/value output from client. 

{"key":1,"value":"a"} 

Put similar key/value pairs for different combinations.
This will shard the data oto different servers (3000,3001,3002)

b) GET http://localhost:3000/keys

Output:

[{"key":2,"value":"b"},{"key":3,"value":"c"}]

c) GET http://localhost:3001/keys

[{"key":9,"value":"i"},{"key":10,"value":"j"},{"key":1,"value":"a"},
{"key":4,"value":"d"},{"key":6,"value":"f"},{"key":7,"value":"g"} 

d) GET http://localhost:3002/keys

[{"key":5,"value":"e"},{"key":8,"value":"h"}]

Packages Imported:
"github.com/julienschmidt/httprouter"
"stathat.com/c/consistent"
