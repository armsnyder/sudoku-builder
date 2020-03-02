<template>
  <table>
    <tbody>
      <tr v-for="(col, i) in grid" :key="i">
        <td
          v-for="(cell, j) in col"
          :key="j"
          :class="cellClass(i, j)"
          @click="onClickCell(i, j)"
          @contextmenu.prevent="onRightClickCell(i, j)"
        >
          {{ cell || '' }}
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  props: {
    pickerSelection: {
      type: Number,
      required: true,
    },
    grid: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      gridSelection: { i: 0, j: 0 },
      showGridSelection: false,
    };
  },
  computed: {
    invalidCells() {
      const result = new Set();

      // Validate columns
      for (let i = 0; i < 9; i++) {
        const seen = new Map();
        for (let j = 0; j < 9; j++) {
          const value = this.grid[i][j];
          if (value) {
            const cell = this.cellIndex(i, j);
            if (seen.has(value)) {
              result.add(cell);
              result.add(seen.get(value));
            }
            seen.set(value, cell);
          }
        }
      }

      // Validate rows
      for (let j = 0; j < 9; j++) {
        const seen = new Map();
        for (let i = 0; i < 9; i++) {
          const value = this.grid[i][j];
          if (value) {
            const cell = this.cellIndex(i, j);
            if (seen.has(value)) {
              result.add(cell);
              result.add(seen.get(value));
            }
            seen.set(value, cell);
          }
        }
      }

      // Validate sections
      for (let outerOffsetI = 0; outerOffsetI < 9; outerOffsetI += 3) {
        for (let outerOffsetJ = 0; outerOffsetJ < 9; outerOffsetJ += 3) {
          const seen = new Map();
          for (let innerOffsetI = 0; innerOffsetI < 3; innerOffsetI++) {
            for (let innerOffsetJ = 0; innerOffsetJ < 3; innerOffsetJ++) {
              const i = outerOffsetI + innerOffsetI;
              const j = outerOffsetJ + innerOffsetJ;
              const value = this.grid[i][j];
              if (value) {
                const cell = this.cellIndex(i, j);
                if (seen.has(value)) {
                  result.add(cell);
                  result.add(seen.get(value));
                }
                seen.set(value, cell);
              }
            }
          }
        }
      }

      return result;
    },
  },
  mounted() {
    window.addEventListener('keydown', (e) => {
      const numericKeyPress = parseInt(e.key, 10);
      if (this.showGridSelection && numericKeyPress) {
        this.setCell(this.gridSelection.i, this.gridSelection.j, numericKeyPress);
        return;
      }

      switch (e.code) {
        case 'ArrowUp':
        case 'KeyW':
          this.showGridSelection = true;
          this.gridSelection.i = Math.max(this.gridSelection.i - 1, 0);
          break;
        case 'ArrowDown':
        case 'KeyS':
          this.showGridSelection = true;
          this.gridSelection.i = Math.min(this.gridSelection.i + 1, 8);
          break;
        case 'ArrowLeft':
        case 'KeyA':
          this.showGridSelection = true;
          this.gridSelection.j = Math.max(this.gridSelection.j - 1, 0);
          break;
        case 'ArrowRight':
        case 'KeyD':
          this.showGridSelection = true;
          this.gridSelection.j = Math.min(this.gridSelection.j + 1, 8);
          break;
        case 'Backspace':
        case 'Delete':
        case 'Digit0':
          this.setCell(this.gridSelection.i, this.gridSelection.j, 0);
          break;
        case 'Escape':
          this.showGridSelection = false;
          break;
        case 'Enter':
          this.setCell(this.gridSelection.i, this.gridSelection.j, this.pickerSelection);
          break;
        default:
          break;
      }
    });
  },
  methods: {
    onClickCell(i, j) {
      this.setCell(i, j, this.pickerSelection);
      this.showGridSelection = false;
      this.gridSelection = { i, j };
    },
    onRightClickCell(i, j) {
      this.setCell(i, j, 0);
      this.showGridSelection = false;
    },
    setCell(i, j, value) {
      this.$emit('cellChange', { i, j, value });
    },
    cellIndex(i, j) {
      return j * 9 + i;
    },
    cellClass(i, j) {
      const result = [];
      if (this.invalidCells.has(this.cellIndex(i, j))) {
        result.push('invalid');
      }
      if (this.showGridSelection && this.gridSelection.i === i && this.gridSelection.j === j) {
        result.push('selected');
      }
      return result.join(' ');
    },
  },
};
</script>

<style scoped>
td {
  height: 48px;
  width: 48px;
  text-align: center;
  font-size: x-large;
  border: black 1px solid;
  cursor: pointer;
  user-select: none;
}

.invalid {
  background-color: lightpink;
}

.selected.selected {
  background-color: lightblue;
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
