FROM gcr.io/distroless/static-debian13:nonroot
ARG TARGETPLATFORM
WORKDIR /
COPY $TARGETPLATFORM/download-geofabrik /download-geofabrik
COPY bbbike.yml /bbbike.yml
COPY geo2day.yml /geo2day.yml
COPY geofabrik.yml /geofabrik.yml
COPY movisda.yml /movisda.yml
COPY openstreetmap.fr.yml /openstreetmap.fr.yml
COPY osm.fit.vutbr.cz.yml /osm.fit.vutbr.cz.yml
COPY osm.kcwu.csie.org.yml /osm.kcwu.csie.org.yml
COPY osm.kewl.lu.yml /osm.kewl.lu.yml
COPY osmit-estratti.yml /osmit-estratti.yml
COPY planet.osm.ch.yml /planet.osm.ch.yml
USER nonroot:nonroot
ENTRYPOINT ["/download-geofabrik"]