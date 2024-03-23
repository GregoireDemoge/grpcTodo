### Instruction:
Create a simple todo front end to consume the gRPC API. 

It needs to be able to:
* Create a Todo item
* Mark a todo item as completed
* Delete a todo item
* Get the previously created todos in a past session

This should take between one to two hours.

For the front-end, you can use the framework of your choice (React, Vue, or just JS). 

We strongly advise you NOT to use Svelte, even if that's what we use with FullEnrich. There is a poorly documented bug between Vite and gRPC web. We fixed it with FullEnrich, but this can make your test needlessly hard.


The goal of this test is not to check if you have the JS level to do that. This shouldn't be an issue.

We expect you to never have worked with gRPC-web in the past, and the challenge lies in:
* Finding how gRPC-web works
* Making it work on your machine

With all that entails (finding documentation online, fixing unexpected problems, etc.)

### To start the gRPC server
On docker:
```console
sudo docker-compose up -d
```

### Ports
50052: Http gateway for grpc-web

### Notes 
The GO Server already has an Envoy proxy setup. No need to set it up yourself.

Restarting the server will flush the memory of all previously created items.
