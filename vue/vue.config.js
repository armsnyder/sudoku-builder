module.exports = {
  // ref: https://cli.vuejs.org/config/#devserver
  devServer: {
    port: '3000',
    proxy: 'http://localhost:8080',
  },
};
