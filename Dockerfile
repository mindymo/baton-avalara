FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/baton-avalara"]
COPY baton-avalara /