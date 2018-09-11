FROM node:10

RUN yarn global add npm i api-client-generator@4.0.1
ADD gen.sh /ngx-swagger/gen.sh

ENTRYPOINT ["/ngx-swagger/gen.sh"]

