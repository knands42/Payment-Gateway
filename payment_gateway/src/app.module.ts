import { Module } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';
import { AppController } from './app.controller';
import { OrdersModule } from './orders/orders.module';
import { AccountsModule } from './accounts/accounts.module';
import { ConfigModule } from '@nestjs/config';
import sequelizeModuleOptions from './config/database/sequelize.config';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      ignoreEnvFile: true,
    }),
    SequelizeModule.forRoot(
      sequelizeModuleOptions[process.env.NODE_ENV ?? 'local'],
    ),
    AccountsModule,
    OrdersModule,
  ],
  controllers: [AppController],
})
export class AppModule {}
