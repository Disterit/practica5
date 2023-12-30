from .views import HomePageView, requestURL, AboutPageView, MySelfPageView, DetailReportView, createReport
from django.urls import path


urlpatterns = [
    path('', HomePageView.as_view(), name='home'),
    path('shorten/', requestURL, name="requestURL"),
    path('about/', AboutPageView.as_view(), name='about'),
    path('myself/', MySelfPageView.as_view(), name='myself'),
    path('report/', DetailReportView.as_view(), name='report'),
    path('create_report/', createReport, name='create_report'),
]