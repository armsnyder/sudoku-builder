<template>
  <Board class="board" />
</template>

<script>
import Board from './components/Board.vue'

export default {
  name: 'app',
  components: {
    Board
  }
}
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
}
</style>
