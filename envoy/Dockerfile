FROM envoyproxy/envoy:v1.15.0
COPY envoy.yaml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l off --component-log-level upstream:info,connection:trace