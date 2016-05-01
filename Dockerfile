FROM scratch
ADD ldap-srv /
ENTRYPOINT [ "/ldap-srv" ]
