import { MiddlewareConsumer, Module, NestModule } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';
import { AppController } from './app.controller';
import { OrdersModule } from './orders/orders.module';
import { AccountsModule } from './accounts/accounts.module';
import { ConfigModule } from '@nestjs/config';
import sequelizeModuleOptions from './config/database/sequelize.config';
import {
  CamelCasetoSnakeCaseMiddleware,
  SnakeCaseToCamelCaseMiddleware,
} from './app.middleware';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      envFilePath: '.env',
    }),
    SequelizeModule.forRoot(
      sequelizeModuleOptions[process.env.NODE_ENV ?? 'local'],
    ),
    AccountsModule,
    OrdersModule,
  ],
  controllers: [AppController],
})
export class AppModule implements NestModule {
  configure(consumer: MiddlewareConsumer) {
    consumer.apply(SnakeCaseToCamelCaseMiddleware).forRoutes('*');
    consumer.apply(CamelCasetoSnakeCaseMiddleware).forRoutes('*');
  }
}
