<template>

  <div class="filter-editor">

    <div class="editor-header">

      <span class="editor-result">{{result}}</span>

    </div>

    <div class="editor-section">

      <slot></slot>

      <button class="editor-button editor-done" @click="$emit('done')">Save</button>

      <button class="editor-button editor-quit" @click="$emit('quit')">Cancel</button>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'filter-editor',

    props: ['result']

  }

</script>

<style scoped>

  .filter-editor {

    background: #fff;
    border: 1px solid #ddd;
    padding: 20px;

  }

  .editor-header {

    margin-bottom: 10px

  }

  .editor-button {

    padding: 10px 50px;
    border: 1px solid #ddd;
    color: #fff;
    outline: none;
    cursor: pointer;

  }

  .editor-done {

    background: #4CAF50;

  }

  .editor-quit {

    background: #F44336;

  }

</style>
