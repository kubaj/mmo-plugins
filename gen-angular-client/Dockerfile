FROM node

RUN yarn global add @flowup/ngx-swagger-client-generator
ADD gen.sh /ngx-swagger/gen.sh

ENTRYPOINT ["/ngx-swagger/gen.sh"]

