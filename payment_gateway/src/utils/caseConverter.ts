export enum CaseType {
  SNAKE_CASE = 'snakecase',
  CAMEL_CASE = 'camelcase',
}

const fromCamelRegex = /([A-Z])/g;
const fromSnakeRegex = /([-_][a-z])/g;
const toCamelRegex = (letter) => letter.toUpperCase();
const toSnakeRegex = (letter) => `_${letter.toLowerCase()}`;

export function convertPayloadCase(
  obj: unknown,
  fromCase: CaseType = CaseType.CAMEL_CASE,
  toCase: CaseType = CaseType.SNAKE_CASE,
): unknown {
  if (typeof obj !== 'object' || obj === null) return obj;

  if (Array.isArray(obj))
    return obj.map((item) => convertPayloadCase(item, fromCase, toCase));

  return Object.keys(obj).reduce((acc, key) => {
    const newKey = key.replace(
      fromCase === CaseType.CAMEL_CASE ? fromCamelRegex : fromSnakeRegex,
      toCase === CaseType.CAMEL_CASE ? toCamelRegex : toSnakeRegex,
    );
    acc[newKey] = convertPayloadCase(obj[key], fromCase, toCase);
    return acc;
  }, {});
}
