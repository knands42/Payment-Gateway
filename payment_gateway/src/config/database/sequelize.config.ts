import { SequelizeModuleOptions } from '@nestjs/sequelize';
import { join } from 'path';
import { Account } from 'src/accounts/entities/account.entity';
import { Order } from 'src/orders/entities/order.entity';

const sequelizeModuleOptionsAllEnvironments = {
  test: {
    dialect: 'sqlite',
    host: join(__dirname, '..', 'orders.db'),
    autoLoadModels: true,
    models: [Order, Account],
    sync: { alter: true },
  } as SequelizeModuleOptions,
  default: {
    dialect: 'postgres',
    host: process.env.DB_HOST ?? 'localhost',
    port: parseInt(process.env.DB_PORT) ?? 5432,
    username: process.env.DB_USERNAME ?? 'postgres',
    password: process.env.DB_PASSWORD ?? 'postgres',
    database: process.env.DB_NAME ?? 'orders',
    autoLoadModels: false,
    models: [Order, Account],
    sync: { alter: false, force: false },
  } as SequelizeModuleOptions,
};

export default sequelizeModuleOptionsAllEnvironments[
  process.env.NODE_ENV ?? 'default'
];
