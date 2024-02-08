FROM gitlab-registry.dg.ru/docker/golang/go-runner:1.17

COPY scripts/migration.sh /migration
COPY migrations_shards /migrations_shards
COPY bin/ .

CMD ["/product-storage-api"]
