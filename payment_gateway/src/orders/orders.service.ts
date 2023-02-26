import { Inject, Injectable, Logger } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { Producer } from '@nestjs/microservices/external/kafka.interface';
import { InjectModel } from '@nestjs/sequelize/dist';
import { Attributes, EmptyResultError, NonNullFindOptions } from 'sequelize';
import { AccountStorageService } from 'src/accounts/account-storage/account-storage.service';
import { convertPayloadCase, customMemoize } from 'src/utils';
import { CreateOrderDto } from './dto/create-order.dto';
import { PublishKafkaDto } from './dto/publish-kafka.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { Order } from './entities/order.entity';

@Injectable()
export class OrdersService {
  private readonly consumerTopic =
    this.configService.get<string>('KAFKA_TOPIC') ?? 'transactions';

  private static readonly toSnakeConverterMemo =
    customMemoize(convertPayloadCase);

  constructor(
    @InjectModel(Order)
    private readonly orderModel: typeof Order,
    private readonly accountStorageService: AccountStorageService,
    @Inject('KAFKA_PRODUCER')
    private readonly kafkaProducer: Producer,
    private readonly configService: ConfigService,
  ) {}

  async create(createOrderDto: CreateOrderDto): Promise<Order> {
    Logger.log(`Creating order ${JSON.stringify(createOrderDto)}`);
    const order = await this.orderModel.create({
      ...createOrderDto,
      accountId: this.accountStorageService.account.id,
    });

    void this.publishToKafka(createOrderDto, order);

    return order;
  }

  findAll() {
    return this.orderModel.findAll({
      where: {
        account_id: this.accountStorageService.account.id,
      },
    });
  }

  findOne(id: string) {
    const account_id = this.accountStorageService.account?.id;
    const query: NonNullFindOptions<Attributes<Order>> = {
      where: {
        id,
      },
      rejectOnEmpty: new EmptyResultError(`Order with ID ${id} not found`),
    };
    if (account_id) query.where['account_id'] = account_id;

    return this.orderModel.findOne(query);
  }

  async update(id: string, updateOrderDto: UpdateOrderDto) {
    const order = await this.findOne(id);
    return order.update(updateOrderDto);
  }

  async remove(id: string) {
    const order = await this.findOne(id);
    return order.destroy();
  }

  private async publishToKafka(createOrderDto: CreateOrderDto, order: Order) {
    const publishKafkaDto = new PublishKafkaDto();
    Object.assign(publishKafkaDto, createOrderDto, { id: order.id });

    const publishKafkaDtoSnakeCase =
      OrdersService.toSnakeConverterMemo(publishKafkaDto);

    const publishKafkaDtoStringify = JSON.stringify(publishKafkaDtoSnakeCase);

    Logger.log(`Publishing order to kafka ${publishKafkaDtoStringify}`);
    await this.kafkaProducer.send({
      topic: this.consumerTopic,
      messages: [
        {
          key: 'transactions',
          value: publishKafkaDtoStringify,
        },
      ],
    });
  }
}
