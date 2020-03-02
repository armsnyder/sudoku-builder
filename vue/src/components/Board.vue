<template>
  <table>
    <tbody>
      <tr v-for="col in 9" :key="col">
        <td v-for="row in 9" :key="row">&nbsp;</td>
      </tr>
    </tbody>
  </table>
</template>

<style scoped>
td {
  padding: 24px;
  border: black 1px solid;
}
td:nth-child(4) {
  border-left: black 3px solid;
}
td:nth-child(7) {
  border-left: black 3px solid;
}
tr:nth-child(4) {
  border-top: black 3px solid;
}
tr:nth-child(7) {
  border-top: black 3px solid;
}
table {
  border: black 3px solid;
  border-collapse: collapse;
}
</style>
