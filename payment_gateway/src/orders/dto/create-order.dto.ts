import { IsNumber, IsString } from 'class-validator';

export class CreateOrderDto {
  @IsNumber()
  amount: number;

  @IsString()
  creditCardNumber: string;

  @IsString()
  creditCardName: string;

  @IsNumber()
  creditCardExpirationMonth: number;

  @IsNumber()
  creditCardExpirationYear: number;

  @IsString()
  creditCardCvv: string;
}
