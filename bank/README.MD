## Bank API

The goal of this exercise is to implement a CORS enabled RESTFul API.

By completing this exercise, you will learn:
* How to enable full-fledged CORS for your RESTFul API
* How to share a state between multiple request handler functions.

In doing so, you will implement the following RESTFul API spec:

`GET /balance`

- Get the current deposit balance 

`PUT /balance`

- Update balance by some integer amount. Request body should be of `application/json` type, and have an amount field, which is a number.
- Ignore any query params, as well as excess parameters in the request body. 
- If amount field is missing in the request body, or if the request body is not provided, throw a BadRequest error. 
- Respond with 200 and a JSON containing `“amt”` field that contains the new balance. 


Implementation Hints:

1. Start a redis server as a docker container

`docker run -d -p 6379:6379 —name redisServer redis`

2. Create a shared resource object containing the  redis client.
- Set a Key named “Balance” to “0”    

3. Implement the BalanceHandler function.  
* For GET and PUT methods, make sure to add Access-Control-Allow-Origin: *  in the response header
* For PUT requests, make sure to handle the CORS pre-flight. 


Testing from the Browser: 

1. Open the provided UI (client.html) 
2. If everything is working, then debiting or crediting the balance should update the overall balance correctly without errors. 
