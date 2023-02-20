'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    await queryInterface.bulkInsert('accounts', [
      {
        id: 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11',
        name: 'João da Silva',
        token: 'ceo0pwvvg0n',
        created_at: new Date(),
        updated_at: new Date(),
      },
      {
        id: 'b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11',
        name: 'João da Silva',
        token: 'wft84s9a9sl',
        created_at: new Date(),
        updated_at: new Date(),
      },
    ]);
  },

  async down(queryInterface, Sequelize) {
    await queryInterface.bulkDelete('accounts', null, {});
  },
};
