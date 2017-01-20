<template>

  <div class="modal">

    <slot></slot>

  </div>

</template>

<script>

  export default {

    name: 'modal'

  }

</script>

<style scoped>

  .modal {

    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background: rgba(0, 0, 0, .5);
    z-index: 10;

  }

</style>
