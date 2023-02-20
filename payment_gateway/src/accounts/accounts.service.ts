import { Injectable } from '@nestjs/common';
import { CreateAccountDto } from './dto/create-account.dto';
import { UpdateAccountDto } from './dto/update-account.dto';
import { InjectModel } from '@nestjs/sequelize';
import { Account } from './entities/account.entity';
import { EmptyResultError, Op } from 'sequelize';

@Injectable()
export class AccountsService {
  constructor(
    @InjectModel(Account) private readonly accountModel: typeof Account,
  ) {}

  create(createAccountDto: CreateAccountDto) {
    return this.accountModel.create({ ...createAccountDto });
  }

  findAll() {
    return this.accountModel.findAll();
  }

  findOne(idOrToken: string) {
    const isUUID = idOrToken.match(
      /^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i,
    );
    return isUUID
      ? this.accountModel.findByPk(idOrToken)
      : this.accountModel.findOne({
          where: { token: idOrToken },
        });
  }

  async update(id: string, updateAccountDto: UpdateAccountDto) {
    const account = await this.findOne(id);
    return account.update({ ...updateAccountDto });
  }

  async remove(id: string) {
    const account = await this.findOne(id);
    return account.destroy();
  }
}
