FROM gcr.io/distroless/static-debian13:nonroot
ARG TARGETPLATFORM
WORKDIR /
COPY $TARGETPLATFORM/download-geofabrik /download-geofabrik
COPY bbbike.yml /etc/download-geofabrik/bbbike.yml
COPY geo2day.yml /etc/download-geofabrik/geo2day.yml
COPY geofabrik.yml /etc/download-geofabrik/geofabrik.yml
COPY movisda.yml /etc/download-geofabrik/movisda.yml
COPY openstreetmap.fr.yml /etc/download-geofabrik/openstreetmap.fr.yml
COPY osm.fit.vutbr.cz.yml /etc/download-geofabrik/osm.fit.vutbr.cz.yml
COPY osm.kcwu.csie.org.yml /etc/download-geofabrik/osm.kcwu.csie.org.yml
COPY osm.kewl.lu.yml /etc/download-geofabrik/osm.kewl.lu.yml
COPY osmit-estratti.yml /etc/download-geofabrik/osmit-estratti.yml
COPY planet.osm.ch.yml /etc/download-geofabrik/planet.osm.ch.yml
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