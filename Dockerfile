# syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM golang:alpine AS build

WORKDIR /build

COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /usr/bin/eirclode ./cmd/eirclode 


################################################################################

FROM scratch AS final

COPY --from=build /usr/bin/eirclode /usr/bin/eirclode

ENV PORT=8080

EXPOSE ${PORT}

USER nonroot:nonroot

ENTRYPOINT [ "/usr/bin/eirclode" ]
