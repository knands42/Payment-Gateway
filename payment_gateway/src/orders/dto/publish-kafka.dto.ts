import { CreateOrderDto } from './create-order.dto';

export class PublishKafkaDto extends CreateOrderDto {
  id: string;
}
