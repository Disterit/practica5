from django.shortcuts import redirect
import socket


class Custom404Middleware:
    def __init__(self, get_response):
        self.get_response = get_response

    def __call__(self, request):
        response = self.get_response(request)
        if response.status_code == 404:
            host = 'host.docker.internal'
            port = 6379

            requested_url = request.path

            if '/shorten' in requested_url:
                requested_url = requested_url[9:]
            else:
                requested_url = "127.0.0.1:8000/" + requested_url

            sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            sock.connect((host, port))


            message = f"--file url.data --query \"TOKENGET url {requested_url}\"\n"
            message_bytes = message.encode('utf-8')

            sock.send(message_bytes)

            received_data = sock.recv(1024)
            site = received_data.decode('utf-8').strip()
            
            sock.close()

            if site != "Элемент не найден":
                host_stat = 'host.docker.internal'
                port_stat = 1337

                sock_stat = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
                sock_stat.connect((host_stat, port_stat))

                ip = get_client_ip(request)

                message = f"1 {site} {requested_url} {ip}"
                message_bytes = message.encode('utf-8')

                sock_stat.send(message_bytes)

                sock_stat.close()

            if site != "Элемент не найден":
                return redirect(site)
            else:
                return redirect("http://localhost:8000/")
        return response


def get_client_ip(request):
    x_forwarded_for = request.META.get('HTTP_X_FORWARDED_FOR')
    if x_forwarded_for:
        ip = x_forwarded_for.split(',')[0]
    else:
        ip = request.META.get('REMOTE_ADDR')
    return ip
