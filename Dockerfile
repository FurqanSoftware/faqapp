FROM ubuntu:xenial

MAINTAINER Mahmud Ridwan <ridwan@furqansoftware.com>

WORKDIR "/opt/faqapp"

ADD faqappd /opt/faqapp/bin/faqappd
ADD ui/gohtml /opt/faqapp/ui/gohtml

CMD ["/opt/faqapp/bin/faqappd"]
