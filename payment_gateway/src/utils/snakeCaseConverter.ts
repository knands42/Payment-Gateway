// eslint-disable-next-line @typescript-eslint/no-unused-vars
export function convertPublisherToSnakeCase(obj: unknown): unknown {
  if (typeof obj !== 'object' || obj === null) return obj;

  if (Array.isArray(obj)) return obj.map(convertPublisherToSnakeCase);

  return Object.keys(obj).reduce((acc, key) => {
    const newKey = key.replace(
      /[A-Z]/g,
      (letter) => `_${letter.toLowerCase()}`,
    );
    acc[newKey] = convertPublisherToSnakeCase(obj[key]);
    return acc;
  }, {});
}
