<template>
  <table>
    <tbody>
      <tr v-for="row in 3" :key="row">
        <td
          v-for="col in 3"
          :key="col"
          :class="selection === digit(row, col) ? 'selected' : ''"
          @click="$emit('select', digit(row, col))"
        >
          {{ digit(row, col) }}
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  props: {
    selection: {
      type: Number,
      default: 1,
    },
  },
  mounted() {
    window.addEventListener('keypress', (e) => {
      const numericKeyPress = parseInt(e.key, 10);
      if (numericKeyPress) {
        this.$emit('select', numericKeyPress);
      }
    });
  },
  methods: {
    digit(row, col) { return (row - 1) * 3 + col; },
  },
};
</script>

<style scoped>
td {
  background-color: lightgray;
  height: 32px;
  width: 32px;
  border-radius: 8px;
  text-align: center;
  cursor: pointer;
  user-select: none;
}

td.selected {
  background-color: lightblue;
  cursor: default;
}
</style>
