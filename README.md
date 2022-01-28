# Capturing JSON data sent by a CircleCI webhook

CircleCI can send to a specified URL a JSON payload describing the state of a build.
For example, it can POST to the URL every time a workflow starts, fails, or finishes.

The code here is a simple example of how to accept a POST request, validate its HMAC signature, capture the JSON payload, and print the payload to the console.
