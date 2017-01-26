<template>

  <div v-if="show" class="instruction">

    <button type="button" class="close" @click="show = false">

      <i class="fa fa-times"></i>

    </button>

    <slot></slot>

  </div>

</template>

<script>

  export default {

    name: 'instruction',

    data() {

      return {

        show: true

      }

    }

  }

</script>

<style>

  .instruction {

    position: relative;
    padding: 10px;
    background: #fcf8e3;
    color: #8a6d3b;
    border: 1px solid #ddd;

  }

  .close {

    position: absolute;
    top: 0;
    right: 0;
    color: inherit;
    padding: 10px;
    background: transparent;
    border: none;
    outline: none;
    cursor: pointer;

  }

</style>
