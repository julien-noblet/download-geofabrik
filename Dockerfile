FROM gcr.io/distroless/static-debian13:nonroot
WORKDIR /
COPY download-geofabrik /download-geofabrik
COPY geofabrik.yml /geofabrik.yml
COPY bbbike.yml /bbbike.yml
COPY openstreetmap.fr.yml /openstreetmap.fr.yml
COPY geo2day.yml /geo2day.yml
USER nonroot:nonroot
ENTRYPOINT ["/download-geofabrik"]