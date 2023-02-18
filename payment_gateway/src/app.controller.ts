import { Controller, Get } from '@nestjs/common';

@Controller('/health')
export class AppController {
  @Get()
  health(): string {
    return 'OK';
  }
}
