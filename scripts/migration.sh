#!/bin/sh

export MIGRATION_SHARDS_DIR=./migrations_shards

echo "Going to apply migrations in ${STAGE} env..."
echo migrations_shards_dir: $MIGRATION_SHARDS_DIR

export DB_CONNECTION="pgsql"
export DB_HOST="127.0.0.1"
export DB_PORT="5432"
export DB_DATABASE="product-api"
export DB_USERNAME="ocb-product-storage-api-user"

WARDEN_HOST=warden.platform.svc.cluster.local

if [ "${STAGE}" = "production" ] || [ "${STAGE}" = "staging" ]; then
  export SHARDS_COUNT=4
  for i in $(eval echo {1..$(($SHARDS_COUNT))}); do
    formatNum=$(printf "%d" $((i)))

    PG_SHARD_ADDRESS=$(wget -cq ${WARDEN_HOST}/endpoints?service=ocb-product-storage-api-shard-${formatNum}.pg:direct -O - | jq '.[] | select(.Role == "master").Address')
    echo PG_SHARD_${formatNum}_ADDRESS ${PG_SHARD_ADDRESS}
    export DB_SHARD_${i}_HOST=$(echo "${PG_SHARD_ADDRESS}" | jq '. | split(":")[0]' | tr -d '"')
    export DB_SHARD_${i}_PORT=$(echo "${PG_SHARD_ADDRESS}" | jq '. | split(":")[1]' | tr -d '"')

    export DB_SHARD_${i}_PASSWORD=$(cat /var/run/secrets/pg__ocb-product-storage-api-shard-${formatNum}__password)
    export DB_SHARD_${i}_USER="ocb-product-storage-api-shard-${formatNum}-user"
    export DB_SHARD_${i}_DB="ocb-product-storage-api-shard-${formatNum}"
  done

elif [ "${STAGE}" = "development" ]; then
  export DB_HOST="${O3_RELEASE_NAME}-infra-postgresql-db.ocb.svc.dev.kube"
  export DB_DATABASE="product-api"
  export DB_USERNAME="postgres"

  export SHARDS_COUNT=2
  for i in $(eval echo {1..$(($SHARDS_COUNT))}); do
    formatNum=$(printf "%d" $((i)))
    export DB_SHARD_${i}_HOST="${O3_RELEASE_NAME}-infra-postgresql-db-shard-${formatNum}.ocb.svc.dev.kube"
    export DB_SHARD_${i}_PORT="5432"
    export DB_SHARD_${i}_PASSWORD=""
    export DB_SHARD_${i}_USER="postgres"
    export DB_SHARD_${i}_DB="product-api-shard-${formatNum}"
  done
fi

postgres="host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_DATABASE} sslmode=disable"

if [ "${STAGE}" = "development" ]; then
  postgres="host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} dbname=${DB_DATABASE} sslmode=disable"
fi

COUNTDONE=0
COUNTALL=$((SHARDS_COUNT))
if [ "$1" = "--dryrun" ]; then
  echo "Migration status"

  for i in $(eval echo {1..$(($SHARDS_COUNT))}); do
    echo "Migration shard ${i} status"
    db_host=DB_SHARD_${i}_HOST
    db_port=DB_SHARD_${i}_PORT
    db_password=DB_SHARD_${i}_PASSWORD
    db_user=DB_SHARD_${i}_USER
    db_name=DB_SHARD_${i}_DB
    goose -dir=${MIGRATION_SHARDS_DIR} postgres "user=${!db_user} dbname=${!db_name} host=${!db_host} port=${!db_port} sslmode=disable password=${!db_password}" status
    if [ $? -eq 0 ]; then COUNTDONE=$((COUNTDONE + 1)); fi
  done

elif [ "$1" = "down" ]; then
  echo "Migration down"

  for i in $(eval echo {1..$(($SHARDS_COUNT))}); do
    echo "Migration shard ${i} down"
    db_host=DB_SHARD_${i}_HOST
    db_port=DB_SHARD_${i}_PORT
    db_password=DB_SHARD_${i}_PASSWORD
    db_user=DB_SHARD_${i}_USER
    db_name=DB_SHARD_${i}_DB
    goose -dir=${MIGRATION_SHARDS_DIR} postgres "user=${!db_user} dbname=${!db_name} host=${!db_host} port=${!db_port} sslmode=disable password=${!db_password}" down
    if [ $? -eq 0 ]; then COUNTDONE=$((COUNTDONE + 1)); fi
  done
else
  echo "Migration up"

  for i in $(eval echo {1..$(($SHARDS_COUNT))}); do
    echo "Migration shard ${i} up"
    db_host=DB_SHARD_${i}_HOST
    db_port=DB_SHARD_${i}_PORT
    db_password=DB_SHARD_${i}_PASSWORD
    db_user=DB_SHARD_${i}_USER
    db_name=DB_SHARD_${i}_DB
    goose -dir=${MIGRATION_SHARDS_DIR} postgres "user=${!db_user} dbname=${!db_name} host=${!db_host} port=${!db_port} sslmode=disable password=${!db_password}" up
    if [ $? -eq 0 ]; then COUNTDONE=$((COUNTDONE + 1)); fi
  done
fi

echo "Result success ${COUNTDONE} from ${COUNTALL}"
if [ "${COUNTDONE}" = "${COUNTALL}" ]; then
  echo OK
else
  echo FAIL
  exit 1
fi
