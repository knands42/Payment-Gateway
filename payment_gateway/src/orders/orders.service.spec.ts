import { ConfigService } from '@nestjs/config';
import { Producer, ProducerRecord } from 'kafkajs';
import { Account } from 'src/accounts/entities/account.entity';
import { AccountStorageService } from '../accounts/account-storage/account-storage.service';
import { CreateOrderDto } from './dto/create-order.dto';
import { Order, OrderStatus } from './entities/order.entity';
import { OrdersService } from './orders.service';

const mockOrderModel = {
  create: jest.fn(),
  findAll: jest.fn(),
  findOne: jest.fn(),
  update: jest.fn(),
  destroy: jest.fn(),
};

const mockAccountStorageService = {
  account: {
    id: '1',
    name: 'test',
    token: '',
    createdAt: new Date(),
    updatedAt: new Date(),
  } as Account,
  setBy: jest.fn(),
};

const mockKafkaProducer = {
  send: jest.fn(),
};

const mockConfigService = {
  get: jest.fn(),
};

describe('OrdersService', () => {
  let sut: OrdersService;

  beforeEach(async () => {
    sut = new OrdersService(
      mockOrderModel as unknown as typeof Order,
      mockAccountStorageService as unknown as AccountStorageService,
      mockKafkaProducer as unknown as Producer,
      mockConfigService as unknown as ConfigService,
    );
  });

  it('validate order returned', async () => {
    // Arrange
    const createOrderDto = new CreateOrderDto();
    createOrderDto.amount = 100;
    createOrderDto.creditCardName = 'Test';
    createOrderDto.creditCardCvv = '123';
    createOrderDto.creditCardExpirationMonth = 12;
    createOrderDto.creditCardExpirationYear = 2022;
    createOrderDto.creditCardNumber = '1234567890123456';

    const expectedResult: Partial<Order> = {
      id: '11',
      amount: 100,
      creditCardName: 'Test',
      creditCardNumber: '123',
      createdAt: new Date(),
      updatedAt: new Date(),
      accountId: '22',
      status: OrderStatus.Pending,
    };

    mockOrderModel.create.mockImplementation(() => expectedResult);

    // Act
    const result = await sut.create(createOrderDto);

    // Assert
    expect(result).toBe(expectedResult);
  });

  it('validate kafka publish message', async () => {
    // Arrange
    const createOrderDto = new CreateOrderDto();
    createOrderDto.amount = 100;
    createOrderDto.creditCardName = 'Test';
    createOrderDto.creditCardNumber = '1234567890123456';
    createOrderDto.creditCardCvv = '123';
    createOrderDto.creditCardExpirationMonth = 12;
    createOrderDto.creditCardExpirationYear = 2022;

    const expectedResult: Partial<Order> = {
      id: '11',
      amount: 100,
      creditCardName: 'Test',
      creditCardNumber: '1234567890123456',
      createdAt: new Date(),
      updatedAt: new Date(),
      accountId: '22',
      status: OrderStatus.Pending,
    };

    mockOrderModel.create.mockImplementation(() => expectedResult);
    const kafkaSendSpy = jest.spyOn(mockKafkaProducer, 'send');
    kafkaSendSpy.mockImplementation(() => Promise.resolve());

    const expectedInput = {
      amount: 100,
      credit_card_name: 'Test',
      credit_card_number: '1234567890123456',
      credit_card_cvv: '123',
      credit_card_expiration_month: 12,
      credit_card_expiration_year: 2022,
      id: '11',
    };

    // Act
    await sut.create(createOrderDto);

    // Assert
    await new Promise((resolve) => setTimeout(resolve, 1000));

    expect(kafkaSendSpy).toBeCalledWith({
      topic: 'transactions',
      messages: [
        {
          key: 'transactions',
          value: JSON.stringify(expectedInput),
        },
      ],
    } as ProducerRecord);
    expect(kafkaSendSpy).toBeCalledTimes(1);
  });
});
