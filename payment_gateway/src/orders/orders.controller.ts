import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
  HttpCode,
  UseGuards,
  Logger,
} from '@nestjs/common';
import { OrdersService } from './orders.service';
import { CreateOrderDto } from './dto/create-order.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { TokenGuard } from 'src/accounts/token/token.guard';
import { MessagePattern, Payload } from '@nestjs/microservices';
import { KafkaMessage } from 'kafkajs';
import { OrderStatus } from './entities/order.entity';

@UseGuards(TokenGuard)
@Controller('orders')
export class OrdersController {
  constructor(private readonly ordersService: OrdersService) {}

  @Post()
  create(@Body() createOrderDto: CreateOrderDto) {
    return this.ordersService.create(createOrderDto);
  }

  @Get()
  findAll() {
    return this.ordersService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.ordersService.findOne(id);
  }

  @Patch(':id')
  update(@Param('id') id: string, @Body() updateOrderDto: UpdateOrderDto) {
    return this.ordersService.update(id, updateOrderDto);
  }

  @Delete(':id')
  @HttpCode(204)
  remove(@Param('id') id: string) {
    return this.ordersService.remove(id);
  }

  @MessagePattern('transactions_result')
  async consumerUpdateStatus(@Payload() message: KafkaMessage) {
    Logger.log(
      `OrdersController.consumerUpdateStatus() message received ${message.value}`,
    );
    const { id, status } = message.value as unknown as kafkaResponse;
    await this.ordersService.update(id, { status });
  }
}

type kafkaResponse = {
  id: string;
  status: OrderStatus;
  errorMessage: string;
};
