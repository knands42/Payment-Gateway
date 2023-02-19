import { SequelizeModuleOptions } from '@nestjs/sequelize';
import { join } from 'path';
import { Account } from 'src/accounts/entities/account.entity';
import { Order } from 'src/orders/entities/order.entity';

type Environments = 'test' | 'local' | 'default';

const sequelizeModuleOptions = {
  test: {
    dialect: 'sqlite',
    host: join(__dirname, '..', 'orders.db'),
    autoLoadModels: true,
    models: [Order, Account],
    sync: { alter: true },
  },
  local: {
    dialect: 'postgres',
    host: process.env.DB_HOST ?? 'localhost',
    port: parseInt(process.env.DB_PORT) ?? 5432,
    username: process.env.DB_USERNAME ?? 'postgres',
    password: process.env.DB_PASSWORD ?? 'postgres',
    database: process.env.DB_NAME ?? 'orders',
    autoLoadModels: false,
    models: [Order, Account],
    sync: { alter: false, force: false },
  },
  default: {
    dialect: 'postgres',
    host: process.env.DB_HOST,
    port: parseInt(process.env.DB_PORT),
    username: process.env.DB_USERNAME,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_NAME,
    autoLoadModels: false,
    models: [Order, Account],
    sync: { alter: false, force: false },
  },
} as { [key in Environments]: SequelizeModuleOptions };

// Named exports used by the Sequelize CLI
export const { test, local, default: production } = sequelizeModuleOptions;
// Default export used by the NestJS app
export default sequelizeModuleOptions;
