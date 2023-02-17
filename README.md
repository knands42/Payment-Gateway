# Payment-Gateway

 Application to simulate an entire payment transaction w/metrics and orchestration

## Overview

This application is a simple payment gateway that simulates a payment transaction. It is composed of 3 microservices:

- **Payment Processor**: A microservice that simulates the payment processing. It receives a payment request and returns a payment response.

- **Payment Gateway**: A microservice that receives a payment request and forwards it to the payment processor. It also receives a payment response and forwards it to the payment processor.