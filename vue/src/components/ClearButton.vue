<template>
  <button class="btn-clear" @click="$emit('click')">Clear Board</button>
</template>

<style scoped>
.btn-clear {
  border: 1px solid hsl(350, 100%, 85%);
  border-radius: 4px;
  padding: 8px 12px;
  background-color: hsl(350, 100%, 90%);
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: medium;
  transition-duration: 0.2s;
  cursor: pointer;
  outline: none;
}

.btn-clear:hover {
  background-color: hsl(350, 100%, 80%);
}

.btn-clear:active {
  background-color: hsl(350, 100%, 75%);
}
</style>
