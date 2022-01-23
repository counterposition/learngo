# Capturing JSON data sent by a CircleCI webhook

CircleCI can send to a specified URL a JSON payload describing the state of a build.
For example, it can POST to the URL every time a workflow starts, fails, or finishes.

The code here is a simple example of how to capture the JSON payload and print it to the console.

Next up: validating that the HTTP request did indeed come from CircleCI (and not an impostor) by checking the body's HMAC signature
