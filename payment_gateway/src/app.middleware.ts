import { Injectable, NestMiddleware } from '@nestjs/common';
import * as _ from 'lodash';

@Injectable()
export class SnakeCaseToCamelCaseMiddleware implements NestMiddleware {
  use(req: any, res: any, next: () => void) {
    if (req.body) {
      req.body = _.mapKeys(req.body, (value, key) => _.camelCase(key));
    }
    next();
  }
}
