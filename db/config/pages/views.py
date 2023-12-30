import json
import os
from django.shortcuts import render
from django.views.generic import TemplateView
import hashlib
import socket


class HomePageView(TemplateView):
    template_name = "home.html"


class AboutPageView(TemplateView):
    template_name = "about.html"


class MySelfPageView(TemplateView):
    template_name = "myself.html"


class DetailReportView(TemplateView):
    template_name = "report.html"


def requestURL(request):
    if request.method == 'GET':
        url = request.GET.get('url', '')

        def shorten_url(url):
            return hashlib.md5(url.encode('utf-8')).hexdigest()[:6]

        shortURL = shorten_url(url)

        host = 'host.docker.internal'
        port = 6379

        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.connect((host, port))

        message = f"--file url.data --query \"TOKENSET url 127.0.0.1:8000/{shortURL} {url}\""

        message_bytes = message.encode('utf-8')

        sock.send(message_bytes)

        sock.close()

        return render(request, 'home.html', {'shortURL': 'localhost:8000/' + shortURL})


def createReport(request):
    if request.method == 'GET':
        report = request.GET.get('report', '')

        host = 'host.docker.internal'
        port = 1337

        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.connect((host, port))

        message = f"2 {report}"

        message_bytes = message.encode('utf-8')

        sock.send(message_bytes)

        sock.close()
        file_path = 'C:\\Users\\Levch\\GolandProjects\\Stat\\report.json'

        if os.path.exists(file_path):
            with open(file_path, 'r') as file:
                report_data = json.load(file)
                return render(request, 'report_create.html', {'report_data': report_data})

        else:
            error_message = "Файл JSON не найден"
            return render(request, 'error.html', {'error_message': error_message})
