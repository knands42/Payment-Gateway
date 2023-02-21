import { Inject, Injectable, Logger } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { Producer } from '@nestjs/microservices/external/kafka.interface';
import { InjectModel } from '@nestjs/sequelize/dist';
import { EmptyResultError } from 'sequelize';
import { AccountStorageService } from 'src/accounts/account-storage/account-storage.service';
import { CreateOrderDto } from './dto/create-order.dto';
import { PublishKafkaDto } from './dto/publish-kafka.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { Order } from './entities/order.entity';

@Injectable()
export class OrdersService {
  private readonly consumerTopic =
    this.configService.get<string>('KAFKA_TOPIC') ?? 'transactions';

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
    return this.orderModel.findOne({
      where: {
        id,
        account_id: this.accountStorageService.account.id,
      },
      rejectOnEmpty: new EmptyResultError(`Order with ID ${id} not found`),
    });
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
    const objToPublish = new PublishKafkaDto();
    objToPublish.amount = createOrderDto.amount;
    objToPublish.creditCardName = createOrderDto.creditCardName;
    objToPublish.creditCardNumber = createOrderDto.creditCardNumber;
    objToPublish.creditCardCvv = createOrderDto.creditCardCvv;
    objToPublish.creditCardExpirationMonth =
      createOrderDto.creditCardExpirationMonth;
    objToPublish.creditCardExpirationYear =
      createOrderDto.creditCardExpirationYear;
    objToPublish.id = order.id;

    const objToPublishStringify = JSON.stringify(objToPublish);

    Logger.log(`Publishing order to kafka ${objToPublishStringify}`);

    await this.kafkaProducer.send({
      topic: this.consumerTopic,
      messages: [
        {
          key: 'transactions',
          value: JSON.stringify(objToPublishStringify),
        },
      ],
    });
  }
}
