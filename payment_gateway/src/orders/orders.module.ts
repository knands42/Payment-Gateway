import { Module } from '@nestjs/common';
import { ClientsModule } from '@nestjs/microservices';
import { ClientKafka } from '@nestjs/microservices/client/client-kafka';
import { SequelizeModule } from '@nestjs/sequelize/dist';
import { AccountsModule } from 'src/accounts/accounts.module';
import { kafkaMicroserviceConfig } from './config/kafka.config';
import { Order } from './entities/order.entity';
import { OrdersController } from './orders.controller';
import { OrdersService } from './orders.service';

@Module({
  imports: [
    SequelizeModule.forFeature([Order]),
    AccountsModule,
    ClientsModule.registerAsync([
      {
        name: 'KAFKA_SERVICE',
        useFactory: () => kafkaMicroserviceConfig,
      },
    ]),
  ],
  controllers: [OrdersController],
  providers: [
    OrdersService,
    {
      provide: 'KAFKA_PRODUCER',
      inject: ['KAFKA_SERVICE'],
      useFactory: async (kafkaService: ClientKafka) => kafkaService.connect(),
    },
  ],
})
export class OrdersModule {}
