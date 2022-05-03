# build environment
FROM envoyproxy/envoy:v1.21-latest
COPY envoy-grpc.yaml /etc/envoy/envoy.yaml
RUN chmod go+r /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml