'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    await queryInterface.bulkInsert('orders', [
      {
        id: '11eebc99-9c0b-4ef8-bb6d-6bb9bd380a11',
        amount: 100.0,
        credit_card_number: '1234567890123456',
        credit_card_name: 'João da Silva',
        status: 'pending',
        account_id: 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11',
        created_at: new Date(),
        updated_at: new Date(),
      },
      {
        id: '22eebc99-9c0b-4ef8-bb6d-6bb9bd380a11',
        amount: 200.0,
        credit_card_number: '1234567890123456',
        credit_card_name: 'João da Silva',
        status: 'pending',
        account_id: 'b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11',
        created_at: new Date(),
        updated_at: new Date(),
      },
    ]);
  },

  async down(queryInterface, Sequelize) {
    await queryInterface.bulkDelete('orders', null, {});
  },
};
