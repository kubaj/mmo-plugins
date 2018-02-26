FROM node

RUN yarn global add npm i api-client-generator@3.0.0-alpha.0
ADD gen.sh /ngx-swagger/gen.sh

ENTRYPOINT ["/ngx-swagger/gen.sh"]

