<template>

  <div v-if="show" class="modal">

    <div class="modal-content">

      <i class="modal-close fa fa-times" @click="$emit('close')"></i>

      <slot></slot>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'modal',

    props: ['show']

  }

</script>

<style scoped>

  .modal {
    display: flex;
    align-items: center;
    justify-content: center;
    position: fixed;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background: rgba(0, 0, 0, .75);
    z-index: 200;
  }

  .modal-content {
    position: relative;
    background: #fff;
    border-radius: 3px;
    border: #ddd;
    padding: 50px;
    box-shadow: 0 3px 12px rgba(27,31,35,0.15);
  }

  .modal-close {
    cursor: pointer;
    position: absolute;
    font-size: 20px;
    color: #555;
    top: 10px;
    right: 10px;
  }

  .modal-close:hover {
    color: #000;
  }

</style>
