#!/bin/bash
# Set IP and port variables
IP="127.0.0.1"
PORT="8080"
USER_DATA='{"name": "John", "age": 30}'

# Health check
printf "\nHealth check...\n"
curl -X GET $IP:$PORT

# Test GET request
printf "\n\n\nTesting GET request...\n"
curl -X GET $IP:$PORT/user/123

# Test POST request
printf "\n\n\nTesting POST request...\n"
curl -X POST -H "Content-Type: application/json" -d "$USER_DATA" $IP:$PORT/user

# Test PUT request
printf "\n\n\nTesting PUT request...\n"
curl -X PUT -H "Content-Type: application/json" -d "$USER_DATA" $IP:$PORT/user

# # Test DELETE request
printf "\n\n\nTesting DELETE request...\n"
curl -X DELETE $IP:$PORT/user/123