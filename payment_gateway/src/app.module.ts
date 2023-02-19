import { Module } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';
import { AppController } from './app.controller';
import { OrdersModule } from './orders/orders.module';
import { AccountsModule } from './accounts/accounts.module';
import { ConfigModule } from '@nestjs/config';
import sequelizeModuleOptions from './config/database/sequelize.config';

@Module({
  imports: [
    OrdersModule,
    SequelizeModule.forRoot(sequelizeModuleOptions),
    ConfigModule.forRoot({
      isGlobal: true,
      ignoreEnvFile: true,
    }),
    AccountsModule,
  ],
  controllers: [AppController],
})
export class AppModule {}
