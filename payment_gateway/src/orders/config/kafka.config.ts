import { Transport } from '@nestjs/microservices/enums/transport.enum';
import { ClientProvider } from '@nestjs/microservices/module/interfaces';

export const kafkaMicroserviceConfig = {
  transport: Transport.KAFKA,
  options: {
    client: {
      clientId: process.env.KAFKA_CLIENT_ID ?? 'payment_gateway',
      brokers: [process.env.KAFKA_HOST ?? 'host.docker.internal:9094'],
      ssl: process.env.KAFKA_USE_SSL === 'true',
    },
    consumer: {
      groupId: process.env.KAFKA_CONSUMER_GROUP_ID ?? 'payment_gateway',
    },
  },
} as ClientProvider;
