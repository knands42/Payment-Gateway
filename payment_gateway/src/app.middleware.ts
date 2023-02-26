import { Injectable, NestMiddleware } from '@nestjs/common';
import { convertPayloadCase, customMemoize } from './utils';
import { CaseType } from './utils/caseConverter';

@Injectable()
export class SnakeCaseToCamelCaseMiddleware implements NestMiddleware {
  static readonly toCamelConverterMemo = customMemoize(convertPayloadCase);

  use(req: any, _res: any, next: () => void) {
    if (req.body) {
      req.body = SnakeCaseToCamelCaseMiddleware.toCamelConverterMemo(
        req.body,
        CaseType.SNAKE_CASE,
        CaseType.CAMEL_CASE,
      ) as any;
    }
    next();
  }
}

@Injectable()
export class CamelCasetoSnakeCaseMiddleware implements NestMiddleware {
  static readonly toSnakeConverterMemo = customMemoize(convertPayloadCase);

  use(_req: any, res: any, next: (error?: any) => void) {
    next();

    if (res.locals.response) {
      res.locals.response = CamelCasetoSnakeCaseMiddleware.toSnakeConverterMemo(
        res.locals.response,
        CaseType.CAMEL_CASE,
        CaseType.SNAKE_CASE,
      );
    }
  }
}
