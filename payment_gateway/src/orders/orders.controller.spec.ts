import { Test, TestingModule } from '@nestjs/testing';
import { AccountStorageService } from 'src/accounts/account-storage/account-storage.service';
import { AccountsModule } from 'src/accounts/accounts.module';
import { TokenGuard } from 'src/accounts/token/token.guard';
import { OrdersController } from './orders.controller';
import { OrdersService } from './orders.service';

describe('OrdersController', () => {
  let controller: OrdersController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [OrdersController],
      providers: [
        {
          provide: OrdersService,
          useValue: {},
        },
        {
          provide: TokenGuard,
          useValue: {},
        },
        {
          provide: AccountStorageService,
          useValue: {},
        },
      ],
    }).compile();

    controller = module.get<OrdersController>(OrdersController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
