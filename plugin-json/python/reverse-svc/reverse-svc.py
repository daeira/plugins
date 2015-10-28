#!/usr/bin/env python2
# coding: utf-8

import pyjsonrpc

class Reverse(pyjsonrpc.HttpRequestHandler):

    @pyjsonrpc.rpcmethod
    def Process(self, data):
        print("reverse greeting")
        reply = {}
        if "greeting" in data:
            reply["greeting"] = data["greeting"][::-1]
            reply["reverse"] = "done"
        else:
            reply["reverse"] = "greeting not found"
        return reply


def main():
    port = 8882
    print("listening on 127.0.0.1:{0}".format(port))
    http_server = pyjsonrpc.ThreadingHttpServer(
        server_address = ('localhost', port),
        RequestHandlerClass = Reverse
    )
    try:
        http_server.serve_forever()
    except KeyboardInterrupt:
        http_server.shutdown()


if __name__ == "__main__":
    main()
