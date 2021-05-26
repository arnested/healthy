FROM scratch

ENV PATH=/

COPY healthy /healthy

ENTRYPOINT ["/healthy"]
