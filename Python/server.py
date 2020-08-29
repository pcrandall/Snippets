#! /usr/bin/python3
import sys
import http.server
import socketserver

class MyHttpRequestHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        if self.path == '/':
            return http.server.SimpleHTTPRequestHandler.do_GET(self)
        else:
            self.path += '.html'
            return http.server.SimpleHTTPRequestHandler.do_GET(self)

args = len(sys.argv) - 1

if args == 0:
    print("Usage: ./server.py <port>")
    exit()

#print("here are your args ", sys.argv[1])

PORT = sys.argv[1]

# Create an object of the above class
handler_object = MyHttpRequestHandler

my_server = socketserver.TCPServer(("", PORT), handler_object)

# Start the server
print("Serving at port", PORT)
my_server.serve_forever()
