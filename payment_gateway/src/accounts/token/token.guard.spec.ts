import { AccountStorageService } from '../account-storage/account-storage.service';
import { TokenGuard } from './token.guard';

describe('TokenGuard', () => {
  it('should be defined', () => {
    expect(new TokenGuard({} as AccountStorageService)).toBeDefined();
  });
});
