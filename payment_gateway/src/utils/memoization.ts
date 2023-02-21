type MemoizeFunc<T extends (...args: any[]) => any> = T & {
  cache: Map<string, ReturnType<T>>;
};

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export function customMemoize<T extends (...args: any[]) => any>(
  func: T,
): MemoizeFunc<T> {
  const memoize = ((...args: Parameters<T>): ReturnType<T> => {
    const cacheKey = JSON.stringify(args);

    if (memoize.cache.has(cacheKey)) {
      return memoize.cache.get(cacheKey);
    }

    const result = func(...args);
    memoize.cache.set(cacheKey, result);
    return result;
  }) as MemoizeFunc<T>;

  memoize.cache = new Map();
  return memoize;
}

export function customMemoize2<T extends (...args: any[]) => any>(
  func: T,
): MemoizeFunc<T> {
  const memoize = {} as MemoizeFunc<T>;
  memoize.cache = new Map();

  return ((...args: Parameters<T>): ReturnType<T> => {
    const cacheKey = JSON.stringify(args);

    if (memoize.cache.has(cacheKey)) {
      return memoize.cache[cacheKey];
    }

    const result = func(...args);
    memoize.cache.set(cacheKey, result);
    return result;
  }) as MemoizeFunc<T>;
}
