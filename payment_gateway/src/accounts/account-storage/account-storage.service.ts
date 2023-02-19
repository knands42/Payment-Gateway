import { Injectable, Scope, Logger } from '@nestjs/common';
import { AccountsService } from '../accounts.service';
import { Account } from '../entities/account.entity';

@Injectable({ scope: Scope.REQUEST })
export class AccountStorageService {
  private _account: Account | null;

  constructor(private readonly accountsService: AccountsService) {}

  get account(): Account | null {
    return this._account;
  }

  async seyBy(token: string): Promise<void> {
    Logger.log('AccountStorageService.seyBy()');
    this._account = await this.accountsService.findOne(token);
  }
}
