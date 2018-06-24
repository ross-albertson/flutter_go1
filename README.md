# flutter_go1

An example of Flutter with a MongoDB-driven Go back end

## Getting Started

For help getting started with Flutter, view our online
[documentation](https://flutter.io/).

You can find the instructions for installing Go at [golang.org](https://golang.org)
After installing Go, you will need to run on a command line:

<code>go get gopkg.in/mgo.v2</code>

This demo uses a database called "rpc1", with a collection called "names".
The fields are "firstName" and "lastName".

## Running flutter_test2
First, start the MongoDB server with the default settings; next, connect your device and

<code>go run name_server1/main.go</code>

<code>flutter run</code> (in a separate console window)

After those steps, the app should work. 

