<template>
  <div class="side-panel">
    <Picker
      :selection="pickerSelection"
      style="margin-bottom: 32px"
      @select="$emit('select', $event)"
    />
    <ClearButton @click="$emit('clearBoard')" />
  </div>
</template>

<script>
import Picker from './Picker';
import ClearButton from './ClearButton';

export default {
  components: {
    Picker,
    ClearButton,
  },
  props: {
    pickerSelection: {
      type: Number,
      default: 1,
    },
  },
};
</script>

<style scoped>
.side-panel {
  display: flex;
  flex-direction: column;
  align-items: center;
  border: 1px solid lightgray;
  border-radius: 16px;
  padding: 32px;
}
</style>
