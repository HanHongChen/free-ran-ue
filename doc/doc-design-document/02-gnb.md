# gNB(RAN)

> [!Note]
> gNB is a base station for handling message transfer between core network and UE.

## Architecture Overview

![free-ran](../image/free-ran.png)

## Basic Connections

### Connections

At gNB(RAN), there are four basic connections:

- To core network:

    - Control Plane(NGAP with AMF)
    - Data Plane(GTP with UPF)

- To UE

    - Control Plane(NAS with UE)
    - Data Plane(Raw packet with UE)

### Procedures

1. NG setup with AMF(port:`38412`)

    - gNB will connect to AMF with setup the basic parameters.

2. GTP tunnel with UPF(port: `2152`)

    - gNB will connect to UPF with setup a GTP tunnel.

3. UE connections

    After receive a new UE control plane connection, gNB will start these two steps:

    - UE registration
    - PDU Session Establishment

## Xn Interface

In current implementation, the Xn interface is designed specially for exchange the TEID information for NR-DC feature.

## GTP Forwarding

For more details about GTP-U, please refer to: [Userspace GTP-U](01-userspace-gtp-u.md)
