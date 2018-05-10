FROM debian
COPY tmpfile/tmpfile tmpfile
COPY httpclient/httpclient httpclient
CMD ["/bin/sh"]