import {
  Column,
  Model,
  PrimaryKey,
  DataType,
  Table,
  CreatedAt,
  UpdatedAt,
} from 'sequelize-typescript';

@Table({
  tableName: 'accounts',
})
export class Account extends Model {
  @PrimaryKey
  @Column({ type: DataType.UUIDV4, defaultValue: DataType.UUIDV4 })
  id: string;

  @Column({ allowNull: false })
  name: string;

  @Column({
    allowNull: false,
    defaultValue: () => Math.random().toString(36).slice(2),
  })
  token: string;

  @CreatedAt
  created_at: Date;

  @UpdatedAt
  updated_at: Date;
}
