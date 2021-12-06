set PASSWORD=%1

docker login -u hopsiia -p %PASSWORD%
docker build -t hopsiia/fhir-isih:latest .
docker push hopsiia/fhir-isih:latest