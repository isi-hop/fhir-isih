version: '3.5'
services:
    hapi-fhir-server:
      build: .
      container_name: hapi-fhir-server
      hostname: hapi-fhir-server
      restart: on-failure
      ports:
        - "8181:8080"
    hapi-fhir-pg:
      image: postgres:12
      container_name: hapi-fhir-pg
      hostname: hapi-db
      restart: always
      environment:
        - POSTGRES_DB=hapi_r4
        - POSTGRES_USER=postgres
        - POSTGRES_PASSWORD=admin
        #- POSTGRES_HOST_AUTH_METHOD=trust
      command: postgres -c 'max_connections=200'
      ports:
        - "5432:5432"
      volumes:
       - "./pgdatafhir:/var/lib/postgresql/data"

networks:
    hapi-bridge:
      name: hapi-network