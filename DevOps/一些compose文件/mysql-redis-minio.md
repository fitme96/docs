version: '2.1'

services:
  mysql:
    image: mysql:8.0.32
    container_name: mysql
    environment:
      MYSQL_ROOT_PASSWORD: HanXin@123
    volumes:
      - mysql_data:/var/lib/mysql
    ports:
      - 3306:3306
  redis:
    image: redis:5.0.5
    container_name: redis
    command: --requirepass HanXin@123
    ports:
      - 6379:6379
  minio:
    image: bitnami/minio:2023.12.23
    container_name: minio
    environment:
      MINIO_ROOT_USER: root
      MINIO_ROOT_PASSWORD: HanXin@123
    volumes:
      - minio_data:/bitnami/minio/data
    ports:
      - 9000:9000
      - 9001:9001
volumes:
  mysql_data:
  minio_data: