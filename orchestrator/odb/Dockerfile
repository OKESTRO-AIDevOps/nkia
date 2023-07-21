FROM mysql

ENV MYSQL_ROOT_PASSWORD youdonthavetoknow
ENV MYSQL_HOST '%'

EXPOSE 3306

ADD ./config.amld /var/lib/mysql-files
ADD ./init.sql /docker-entrypoint-initdb.d