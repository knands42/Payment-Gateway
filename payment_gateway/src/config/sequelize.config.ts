import { SequelizeModuleOptions } from '@nestjs/sequelize';
import { join } from 'path';
import { Account } from 'src/accounts/entities/account.entity';
import { Order } from 'src/orders/entities/order.entity';

const sequelizeModuleOptionsDev: SequelizeModuleOptions = {
  dialect: 'sqlite',
  host: join(__dirname, '..', 'orders.db'),
  autoLoadModels: true,
  models: [Order, Account],
  sync: { alter: true },
};

const sequelizeModuleOptionsProd: SequelizeModuleOptions = {
  dialect: 'postgres',
  host: process.env.DB_HOST,
  port: parseInt(process.env.DB_PORT),
  username: process.env.DB_USERNAME,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
  autoLoadModels: false,
  models: [Order, Account],
  sync: { alter: false, force: false },
};

export const sequelizeModuleOptions =
  process.env.NODE_ENV === 'production'
    ? sequelizeModuleOptionsProd
    : sequelizeModuleOptionsDev;
