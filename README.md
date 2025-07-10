# free-ran-ue

![free-ran-ue](/doc/image/free-ran-ue.jpg)

## Introduction

This project is going to simulate the behavior between core network and RAN/UE. The basic idea is to design a useful tool for testing the feature [NR-DC](https://free5gc.org/blog/20250219/20250219/).

For more details, please checkout: [free-ran-ue official website](https://alonza0314.github.io/free-ran-ue/)

## Log Description

There will be 5 types of log level can choose in both **gNB** and **UE**:

- ERROR: Significant error which makes the app stop.
- WARN: Someting strange but does not get impact on the app.
- INFO: Informations that user should know.
- DEBUG: Informations for developer getting
- TRACE: More detail information in each steps.

This can customized in the [configuration files](/config).
