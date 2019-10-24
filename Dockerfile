FROM scratch
COPY download-geofabrik /
COPY bbbike.yml /
COPY geofabrik.yml /
COPY openstreetmap.fr.yml /
ENTRYPOINT ["/download-geofabrik"]

