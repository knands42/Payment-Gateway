'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    await queryInterface.addColumn('orders', 'account_id', {
      type: Sequelize.UUID,
      allowNull: false,
      references: {
        model: 'accounts',
        key: 'id',
      },
    });
  },

  async down(queryInterface, Sequelize) {
    await queryInterface.removeColumn('orders', 'account_id');
  },
};
