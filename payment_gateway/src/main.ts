import {
  INestApplication,
  Logger,
  ValidationPipe,
  VersioningType,
} from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { NestFactory } from '@nestjs/core';
import 'colors';
import { AppModule } from './app.module';
import { kafkaMicroserviceConfig } from './config/kafka/kafka.config';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  await appSetup(app);

  const configService = app.get(ConfigService);
  const port = configService.get('PORT') ?? 3001;
  await app.listen(port);
  Logger.log(
    `Server running on port ${port} in ${process.env.NODE_ENV ?? 'Debug'} mode`
      .blue.bold,
  );
}

async function appSetup(app: INestApplication) {
  app.enableCors();
  app.setGlobalPrefix('api');
  app.enableVersioning({
    type: VersioningType.URI,
    defaultVersion: '1',
  });
  app.useGlobalPipes(
    new ValidationPipe({
      whitelist: true,
    }),
  );
  app.connectMicroservice(kafkaMicroserviceConfig);
  await app.startAllMicroservices();
}

bootstrap();
