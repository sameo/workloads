FROM debian
COPY tmpfile/tmpfile tmpfile
COPY ctrfile/ctrfile ctrfile
COPY httpclient/httpclient httpclient
COPY fork/fork fork
CMD ["/bin/sh"]