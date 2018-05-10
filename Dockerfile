FROM debian
COPY tmpfile/tmpfile tmpfile
COPY httpclient/httpclient httpclient
COPY fork/fork fork
CMD ["/bin/sh"]