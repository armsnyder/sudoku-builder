<template>
  <div class="content">
    <Board style="margin-right: 128px" :picker-selection="pickerSelection" />
    <Picker :selection="pickerSelection" @select="pickerSelection = $event" />
  </div>
</template>

<script>
import Board from './components/Board';
import Picker from './components/Picker';

export default {
  components: {
    Picker,
    Board,
  },
  data() {
    return {
      pickerSelection: 1,
    };
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
