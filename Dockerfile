FROM gcr.io/distroless/base
COPY download-geofabrik /
COPY bbbike.yml /
COPY geofabrik.yml /
COPY openstreetmap.fr.yml /
ENV OUTPUT_DIR /data
ENTRYPOINT ["/download-geofabrik"]

