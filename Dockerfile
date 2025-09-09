# BUILD ENV ###################################################################

FROM golang:1.24-alpine AS build

ARG BAO_USER=alketas

ENV OPENBAO_RELEASE='https://github.com/openbao/openbao/releases/download/v2.4.0/bao_2.4.0_Linux_x86_64.tar.gz'
ENV BAO_VERSION=2.4.0

COPY ./src/  /${BAO_USER}-src/

RUN adduser  --system --no-create-home  ${BAO_USER}

RUN wget ${OPENBAO_RELEASE}
RUN tar -xvzf bao_${BAO_VERSION}_Linux_x86_64.tar.gz
RUN mv bao /bin/bao
RUN bao --version

WORKDIR /${BAO_USER}-src
RUN go get gopkg.in/yaml.v2
RUN go mod download
RUN go build  -v -o /bin/${BAO_USER} main.go
RUN chmod o=x /bin/${BAO_USER}

# RUN ENV #####################################################################
FROM golang:1.24-alpine
ARG BAO_USER=alketas

ENV VAULT_RENEW_TOKENS='/alketas/confs/tokens.yaml'
ENV VAULT_ADDR='https://openbao.XXXX.YYY.local:443'
ENV VAULT_CACERT="root-ca.crt"
ENV VAULT_SKIP_VERIFY=false

COPY --from=build /bin/alketas  /bin/alketas
COPY --from=build /bin/bao  /bin/bao


RUN adduser  --system --no-create-home  ${BAO_USER}
USER ${BAO_USER}
WORKDIR /
RUN ls -lah /bin/alketas
CMD ["/bin/alketas"]