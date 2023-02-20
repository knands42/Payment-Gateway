import {
  BelongsTo,
  Column,
  CreatedAt,
  DataType,
  ForeignKey,
  Model,
  PrimaryKey,
  Table,
  UpdatedAt,
} from 'sequelize-typescript';
import { Account } from 'src/accounts/entities/account.entity';

export enum OrderStatus {
  Pending = 'pending',
  Approved = 'approved',
  Rejected = 'rejected',
}

@Table({
  tableName: 'orders',
})
export class Order extends Model {
  @PrimaryKey
  @Column({
    type: DataType.UUIDV4,
    defaultValue: DataType.UUIDV4,
  })
  id: string;

  @Column({
    type: DataType.DECIMAL(10, 2),
    allowNull: false,
  })
  amount: number;

  @Column({
    type: DataType.STRING,
    allowNull: false,
  })
  creditCardNumber: string;

  @Column({
    type: DataType.STRING,
    allowNull: false,
  })
  creditCardName: string;

  @Column({
    type: DataType.ENUM(
      OrderStatus.Pending,
      OrderStatus.Approved,
      OrderStatus.Rejected,
    ),
    allowNull: false,
    defaultValue: OrderStatus.Pending,
  })
  status: OrderStatus;

  @ForeignKey(() => Account)
  @Column({
    type: DataType.UUIDV4,
    allowNull: false,
  })
  accountId: string;

  @BelongsTo(() => Account)
  account: Account;

  @CreatedAt
  created_at: Date;

  @UpdatedAt
  updated_at: Date;
}
