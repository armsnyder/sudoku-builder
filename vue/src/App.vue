<template>
  <div class="content">
    <Board
      style="margin-right: 128px"
      :picker-selection="pickerSelection"
      :grid="grid"
      @cellChange="onCellChange"
    />
    <Picker :selection="pickerSelection" @select="pickerSelection = $event" />
  </div>
</template>

<script>
import Board from './components/Board';
import Picker from './components/Picker';

const LOCAL_STORAGE_KEY = 'boardGrid';

export default {
  components: {
    Picker,
    Board,
  },
  data() {
    return {
      pickerSelection: 1,
      grid: [],
    };
  },
  watch: {
    grid() {
      localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(this.grid));
    },
  },
  created() {
    if (localStorage.getItem(LOCAL_STORAGE_KEY)) {
      this.grid = JSON.parse(localStorage.getItem(LOCAL_STORAGE_KEY));
    } else {
      this.clearBoard();
    }
  },
  methods: {
    onCellChange({ i, j, value }) {
      const newI = this.grid[i].slice(0);
      newI[j] = value;
      this.$set(this.grid, i, newI);
    },
    clearBoard() {
      const newGrid = [];
      for (let i = 0; i < 9; i++) {
        newGrid[i] = [];
        for (let j = 0; j < 9; j++) {
          newGrid[i][j] = 0;
        }
      }
      this.grid = newGrid;
    },
  },
};
</script>

<style>
body, html {
  height: 100%;
}

body {
  background: white;
  margin: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: sans-serif;
}

.content {
  display: flex;
  align-items: center;
}
</style>
