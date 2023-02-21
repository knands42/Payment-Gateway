import {
  CanActivate,
  ExecutionContext,
  Injectable,
  Logger,
} from '@nestjs/common';
import { AccountStorageService } from '../account-storage/account-storage.service';

@Injectable()
export class TokenGuard implements CanActivate {
  constructor(private readonly accountStorage: AccountStorageService) {}

  async canActivate(context: ExecutionContext): Promise<boolean> {
    if (context.getType() !== 'http') {
      return true;
    }

    const request = context.switchToHttp().getRequest();
    const token = request.headers?.['x-token'] as string;

    Logger.log('TokenGuard.canActivate()');

    if (token) {
      try {
        await this.accountStorage.seyBy(token);
        return true;
      } catch (e) {
        console.error(e);
        return false;
      }
    }

    return false;
  }
}
