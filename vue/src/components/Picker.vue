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
  created() {
    window.addEventListener('keypress', this.onKeyPress);
  },
  beforeDestroy() {
    window.removeEventListener('keypress', this.onKeyDown);
  },
  methods: {
    digit(row, col) { return (row - 1) * 3 + col; },
    onKeyPress({ key }) {
      const numericKeyPress = parseInt(key, 10);
      if (numericKeyPress) {
        this.$emit('select', numericKeyPress);
      }
    },
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
