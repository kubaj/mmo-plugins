FROM node

RUN git clone https://github.com/flowup/ngx-swagger-client-generator /ngx-swagger && cd /ngx-swagger && yarn
ADD gen.sh /ngx-swagger/gen.sh

WORKDIR /ngx-swagger

ENTRYPOINT ["/ngx-swagger/gen.sh"]

