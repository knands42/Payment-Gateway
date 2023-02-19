// eslint-disable-next-line @typescript-eslint/no-var-requires
const config = require('../jest.config');
config.testRegex = '.e2e-spec.ts$';
config.rootDir = '.';

module.exports = config;
