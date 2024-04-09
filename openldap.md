
```yaml

[root@bk2 openldap]# cat docker-compose.yml 
# Copyright VMware, Inc.
# SPDX-License-Identifier: APACHE-2.0

version: '2'

services:
  openldap:
    image: docker.io/bitnami/openldap:2.6
    ports:
      - '389:389'
      - '636:636'
    environment:
      - LDAP_ROOT=dc=hx,dc=com
      - LDAP_ADMIN_DN=cn=admin,dc=hx,dc=com
      - LDAP_ADMIN_PASSWORD=adminpassword
      - LDAP_ADMIN_USERNAME=admin
      - LDAP_PORT_NUMBER=389
      - LDAP_LDAPS_PORT_NUMBER=636
    volumes:
      - 'openldap_data:/bitnami/openldap'
  phpldapadmin:
    image: osixia/phpldapadmin:latest
    ports:
      - 18081:80
    environment:
      - PHPLDAPADMIN_HTTPS=false
      - PHPLDAPADMIN_LDAP_HOSTS=openldap
volumes:
  openldap_data:
```