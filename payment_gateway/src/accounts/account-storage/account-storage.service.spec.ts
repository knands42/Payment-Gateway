import { Test, TestingModule } from '@nestjs/testing';
import { AccountsService } from '../accounts.service';
import { AccountStorageService } from './account-storage.service';

describe('AccountStorageService', () => {
  let service: AccountStorageService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        AccountStorageService,
        {
          provide: AccountsService,
          useValue: {},
        },
      ],
    }).compile();

    service = await module.resolve(AccountStorageService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
