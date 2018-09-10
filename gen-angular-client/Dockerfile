FROM node

RUN yarn global add npm i api-client-generator@3.0.0-alpha.1
ADD gen.sh /ngx-swagger/gen.sh

ENTRYPOINT ["/ngx-swagger/gen.sh"]

