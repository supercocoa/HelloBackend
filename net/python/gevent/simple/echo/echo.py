import socket
import gevent

HOST = 'localhost'
PORT = 50009


def handleReq(conn, addr):
    print 'handleReq'
    while 1:
        data = conn.recv(1024)
        if not data:
            break
        conn.sendall(data)
    conn.close()


def createSvr():
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.bind((HOST, PORT))
    sock.listen(1)
    while True:
        conn, addr = sock.accept()
        print 'conn by ', addr
        gevent.spawn(handleReq, conn, addr)


if __name__ == '__main__':
    createSvr()
