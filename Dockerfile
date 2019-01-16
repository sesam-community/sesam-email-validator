FROM iron/base

COPY sesam-email-validator /opt/service/

WORKDIR /opt/service

RUN chmod +x /opt/service/sesam-email-validator

EXPOSE 8080:8080

CMD /opt/service/sesam-email-validator