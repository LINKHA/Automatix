docker build -t automatix .
docker network create automatix_net
docker-compose -f docker-compose-env.yml up -d