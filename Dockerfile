FROM scratch
ENTRYPOINT ["/download-geofabrik"]
COPY download-geofabrik /
COPY geofabrik.yml /
COPY bbbike.yml /
COPY openstreetmap.fr.yml /
COPY osmtoday.yml /