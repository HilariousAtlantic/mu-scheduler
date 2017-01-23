<template>

  <div class="double-header">

    <slot name="left"></slot>

    <slot name="right"></slot>

  </div>

</template>

<script>

  export default {

    name: 'double-header'

  }

</script>

<style scoped>

  .double-header {

    display: flex;
    justify-content: space-between;

  }

</style>
